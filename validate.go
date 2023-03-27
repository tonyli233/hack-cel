package hack_cel

import (
	"fmt"
	validatev2 "github.com/bufbuild/hack-cel/gen/buf/validate/v2"
	"github.com/bufbuild/hack-cel/gen/validate"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"log"
	"reflect"
)

type result struct {
	err error
}

func Validate(m proto.Message) error {
	r := &result{}
	m.ProtoReflect().Range(reflectUserDefinedFields(r))
	if r.err != nil {
		return r.err
	}
	return nil
}

func reflectUserDefinedFields(r *result) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch kind := fd.Kind(); kind {
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			opts, ok := fd.Options().(*descriptorpb.FieldOptions)
			if ok {
				extension := proto.GetExtension(opts, validate.E_Rules)
				if fieldRules, ok := extension.(*validate.FieldRules); ok {
					floatRulesReflect := fieldRules.GetFloat().ProtoReflect()
					floatRulesReflect.Range(reflectInternalRules(r, v, fieldRules))
				}
			}
		default:
			log.Fatalf("kind: %v is not implemented", kind.String())
		}
		return true
	}
}

func reflectInternalRules(r *result, value protoreflect.Value, fieldRules *validate.FieldRules) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch kind := fd.Kind(); kind {
		case protoreflect.FloatKind:
			opts, ok := fd.Options().(*descriptorpb.FieldOptions)
			if ok {
				extension := proto.GetExtension(opts, validatev2.E_InternalRules)
				if internalRules, ok := extension.(*validatev2.InternalRules); ok {
					if constRule := internalRules.GetFloat().GetConst(); constRule != nil {
						if err := hackCel(constRule, value.Float(), fieldRules); err != nil {
							r.err = err
							return false
						}
					}
					// TODO:
					//  if ltRule := internalRules.GetFloat().GetLt(); ltRule != nil {
					//	 ...
					//  }
				}
			}
		}
		return true
	}
}

type rule interface {
	GetExpression() string
}

// hackCel into the validator, floatConstRules.expression is going to be evaluated against CEL
func hackCel(r rule, this float64, fieldRules *validate.FieldRules) error {
	env, err := cel.NewEnv(
		cel.Variable("this", cel.DoubleType),
		cel.Types(&validate.FieldRules{}),
		cel.Variable("rules", cel.ObjectType("validate.FieldRules")),
	)
	if err != nil {
		return err
	}
	ast, err := compile(env, r.GetExpression(), cel.BoolType)
	if err != nil {
		return err
	}
	program, err := env.Program(ast)
	if err != nil {
		return err
	}
	vars := map[string]interface{}{
		"this":  this,
		"rules": fieldRules,
	}
	out, _, err := program.Eval(vars)
	if err != nil {
		return err
	}
	celResult, ok := out.Value().(bool)
	if !ok {
		return fmt.Errorf("expect cel.BoolType output, but got %T", out.Value())
	}
	if !celResult {
		return fmt.Errorf("cel expression evaluated to false in %v", fieldRules.GetType())
	}
	return nil
}

// below are code copied from cel-go/codelabs

// compile will parse and check an expression `expr` against a given
// environment `env` and determine whether the resulting type of the expression
// matches the `exprType` provided as input.
func compile(env *cel.Env, expr string, celType *cel.Type) (*cel.Ast, error) {
	ast, iss := env.Compile(expr)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	if !reflect.DeepEqual(ast.OutputType(), celType) {
		return nil, fmt.Errorf("got %v, wanted %v result type", ast.OutputType(), celType)
	}
	return ast, nil
}

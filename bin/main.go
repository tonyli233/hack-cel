package main

import (
	"encoding/json"
	"fmt"
	validatev2 "github.com/bufbuild/buf-tour/gen/buf/validate/v2"
	petv1 "github.com/bufbuild/buf-tour/gen/pet/v1"
	"github.com/bufbuild/buf-tour/gen/validate"
	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"reflect"
	"sort"
	"strings"
)

func main() {
	pet := &petv1.Pet{
		Foo: 1,
	}
	myValidate(pet)
}

func myValidate(pb proto.Message) proto.Message {
	m := pb.ProtoReflect()

	c := validateTopLevel(m)
	m.Range(c)
	return pb
}

func validateTopLevel(m protoreflect.Message) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		kind := fd.Kind()
		switch kind {
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			opts, ok := fd.Options().(*descriptorpb.FieldOptions)
			if ok {
				extension := proto.GetExtension(opts, validate.E_Rules)
				if fieldRules, ok := extension.(*validate.FieldRules); ok {
					floatRulesReflect := fieldRules.GetFloat().ProtoReflect()
					floatRulesReflect.Range(reflectInternalRules(floatRulesReflect, v, fieldRules))
				}
			}
		}
		return true
	}
}

func reflectInternalRules(m protoreflect.Message, value protoreflect.Value, fieldRules *validate.FieldRules) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		switch kind := fd.Kind(); kind {
		case protoreflect.FloatKind:
			opts, ok := fd.Options().(*descriptorpb.FieldOptions)
			if ok {
				extension := proto.GetExtension(opts, validatev2.E_InternalRules)
				if internalRules, ok := extension.(*validatev2.InternalRules); ok {
					hackCel(internalRules.GetFloat().GetConst(), value.Float(), fieldRules)
				}
			}
		}
		return true
	}
}

// hackCel into the validator, floatConstRules.expression is going to be evaluated against CEL
func hackCel(floatConstRules *validatev2.InternalFloatsConstRules, this float64, fieldRules *validate.FieldRules) {
	env, err := cel.NewEnv(
		cel.Variable("this", cel.DoubleType),
		cel.Types(&validate.FieldRules{}),
		cel.Variable("rules", cel.ObjectType("validate.FieldRules")),
	)
	if err != nil {
		panic(err)
	}
	ast := compile(env, *floatConstRules.Expression, cel.BoolType)
	program, err := env.Program(ast)
	if err != nil {
		panic(err)
	}
	vars := map[string]interface{}{
		"this":  this,
		"rules": fieldRules,
	}
	out, _, err := eval(program, vars)
	if err != nil {
		panic(err)
	}
	fmt.Println(out.Type(), out.Value())
}

// below are code copied from cel-go/codelabs

// compile will parse and check an expression `expr` against a given
// environment `env` and determine whether the resulting type of the expression
// matches the `exprType` provided as input.
func compile(env *cel.Env, expr string, celType *cel.Type) *cel.Ast {
	ast, iss := env.Compile(expr)
	if iss.Err() != nil {
		panic(iss.Err())
	}
	if !reflect.DeepEqual(ast.OutputType(), celType) {
		panic(fmt.Sprintf(
			"Got %v, wanted %v result type", ast.OutputType(), celType))
	}
	fmt.Printf("%s\n\n", strings.ReplaceAll(expr, "\t", " "))
	return ast
}

func eval(prg cel.Program,
	vars any) (out ref.Val, det *cel.EvalDetails, err error) {
	varMap, isMap := vars.(map[string]any)
	fmt.Println("------ input ------")
	if !isMap {
		fmt.Printf("(%T)\n", vars)
	} else {
		for k, v := range varMap {
			switch val := v.(type) {
			case proto.Message:
				bytes, err := prototext.Marshal(val)
				if err != nil {
					glog.Exitf("failed to marshal proto to text: %v", val)
				}
				fmt.Printf("%s = %s", k, string(bytes))
			case map[string]any:
				b, _ := json.MarshalIndent(v, "", "  ")
				fmt.Printf("%s = %v\n", k, string(b))
			case uint64:
				fmt.Printf("%s = %vu\n", k, v)
			default:
				fmt.Printf("%s = %v\n", k, v)
			}
		}
	}
	fmt.Println()
	out, det, err = prg.Eval(vars)
	report(out, det, err)
	fmt.Println()
	return
}

// report prints out the result of evaluation in human-friendly terms.
func report(result ref.Val, details *cel.EvalDetails, err error) {
	fmt.Println("------ result ------")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("value: %v (%T)\n", result, result)
	}
	if details != nil {
		fmt.Printf("\n------ eval states ------\n")
		state := details.State()
		stateIDs := state.IDs()
		ids := make([]int, len(stateIDs), len(stateIDs))
		for i, id := range stateIDs {
			ids[i] = int(id)
		}
		sort.Ints(ids)
		for _, id := range ids {
			v, found := state.Value(int64(id))
			if !found {
				continue
			}
			fmt.Printf("%d: %v (%T)\n", id, v, v)
		}
	}
}

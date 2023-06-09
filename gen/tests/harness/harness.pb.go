// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tests/harness/harness.proto

package harness

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *anypb.Any `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestCase) Reset() {
	*x = TestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_harness_harness_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCase) ProtoMessage() {}

func (x *TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_tests_harness_harness_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCase.ProtoReflect.Descriptor instead.
func (*TestCase) Descriptor() ([]byte, []int) {
	return file_tests_harness_harness_proto_rawDescGZIP(), []int{0}
}

func (x *TestCase) GetMessage() *anypb.Any {
	if x != nil {
		return x.Message
	}
	return nil
}

type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid               bool     `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`
	Error               bool     `protobuf:"varint,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Reasons             []string `protobuf:"bytes,3,rep,name=Reasons,proto3" json:"Reasons,omitempty"`
	AllowFailure        bool     `protobuf:"varint,4,opt,name=AllowFailure,proto3" json:"AllowFailure,omitempty"`
	CheckMultipleErrors bool     `protobuf:"varint,5,opt,name=CheckMultipleErrors,proto3" json:"CheckMultipleErrors,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_harness_harness_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_tests_harness_harness_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_tests_harness_harness_proto_rawDescGZIP(), []int{1}
}

func (x *TestResult) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *TestResult) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *TestResult) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

func (x *TestResult) GetAllowFailure() bool {
	if x != nil {
		return x.AllowFailure
	}
	return false
}

func (x *TestResult) GetCheckMultipleErrors() bool {
	if x != nil {
		return x.CheckMultipleErrors
	}
	return false
}

var File_tests_harness_harness_proto protoreflect.FileDescriptor

var file_tests_harness_harness_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x13,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0xa6,
	0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72,
	0x6e, 0x65, 0x73, 0x73, 0x42, 0x0c, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x68, 0x61, 0x63, 0x6b, 0x2d, 0x63,
	0x65, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x68, 0x61, 0x72,
	0x6e, 0x65, 0x73, 0x73, 0xa2, 0x02, 0x03, 0x54, 0x48, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0xca, 0x02, 0x0d, 0x54, 0x65, 0x73,
	0x74, 0x73, 0x5c, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x19, 0x54, 0x65, 0x73,
	0x74, 0x73, 0x5c, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x3a,
	0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_harness_harness_proto_rawDescOnce sync.Once
	file_tests_harness_harness_proto_rawDescData = file_tests_harness_harness_proto_rawDesc
)

func file_tests_harness_harness_proto_rawDescGZIP() []byte {
	file_tests_harness_harness_proto_rawDescOnce.Do(func() {
		file_tests_harness_harness_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_harness_harness_proto_rawDescData)
	})
	return file_tests_harness_harness_proto_rawDescData
}

var file_tests_harness_harness_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_harness_harness_proto_goTypes = []interface{}{
	(*TestCase)(nil),   // 0: tests.harness.TestCase
	(*TestResult)(nil), // 1: tests.harness.TestResult
	(*anypb.Any)(nil),  // 2: google.protobuf.Any
}
var file_tests_harness_harness_proto_depIdxs = []int32{
	2, // 0: tests.harness.TestCase.message:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tests_harness_harness_proto_init() }
func file_tests_harness_harness_proto_init() {
	if File_tests_harness_harness_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_harness_harness_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tests_harness_harness_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_harness_harness_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_harness_harness_proto_goTypes,
		DependencyIndexes: file_tests_harness_harness_proto_depIdxs,
		MessageInfos:      file_tests_harness_harness_proto_msgTypes,
	}.Build()
	File_tests_harness_harness_proto = out.File
	file_tests_harness_harness_proto_rawDesc = nil
	file_tests_harness_harness_proto_goTypes = nil
	file_tests_harness_harness_proto_depIdxs = nil
}

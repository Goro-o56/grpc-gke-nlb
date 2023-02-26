// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: gpt.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 型の定義
type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{0}
}

func (x *Prompt) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type InputAndInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input       string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Instruction string `protobuf:"bytes,2,opt,name=instruction,proto3" json:"instruction,omitempty"`
}

func (x *InputAndInstruction) Reset() {
	*x = InputAndInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputAndInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputAndInstruction) ProtoMessage() {}

func (x *InputAndInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputAndInstruction.ProtoReflect.Descriptor instead.
func (*InputAndInstruction) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{1}
}

func (x *InputAndInstruction) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *InputAndInstruction) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

type PersonaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PersonaResponse) Reset() {
	*x = PersonaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonaResponse) ProtoMessage() {}

func (x *PersonaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gpt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonaResponse.ProtoReflect.Descriptor instead.
func (*PersonaResponse) Descriptor() ([]byte, []int) {
	return file_gpt_proto_rawDescGZIP(), []int{2}
}

func (x *PersonaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gpt_proto protoreflect.FileDescriptor

var file_gpt_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x79, 0x61,
	0x70, 0x70, 0x22, 0x1e, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x4d, 0x0a, 0x13, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x6e, 0x64, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x2b, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x67,
	0x0a, 0x1a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x45, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x13,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x45,
	0x64, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x41, 0x6e, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x16, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x66, 0x0a, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x19, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_gpt_proto_rawDescOnce sync.Once
	file_gpt_proto_rawDescData = file_gpt_proto_rawDesc
)

func file_gpt_proto_rawDescGZIP() []byte {
	file_gpt_proto_rawDescOnce.Do(func() {
		file_gpt_proto_rawDescData = protoimpl.X.CompressGZIP(file_gpt_proto_rawDescData)
	})
	return file_gpt_proto_rawDescData
}

var file_gpt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gpt_proto_goTypes = []interface{}{
	(*Prompt)(nil),              // 0: myapp.Prompt
	(*InputAndInstruction)(nil), // 1: myapp.InputAndInstruction
	(*PersonaResponse)(nil),     // 2: myapp.PersonaResponse
}
var file_gpt_proto_depIdxs = []int32{
	1, // 0: myapp.GeneratePersonaEditService.GeneratePersonaEdit:input_type -> myapp.InputAndInstruction
	0, // 1: myapp.GeneratePersonaCompletionService.GeneratePersonaCompletion:input_type -> myapp.Prompt
	2, // 2: myapp.GeneratePersonaEditService.GeneratePersonaEdit:output_type -> myapp.PersonaResponse
	2, // 3: myapp.GeneratePersonaCompletionService.GeneratePersonaCompletion:output_type -> myapp.PersonaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gpt_proto_init() }
func file_gpt_proto_init() {
	if File_gpt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gpt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prompt); i {
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
		file_gpt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputAndInstruction); i {
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
		file_gpt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonaResponse); i {
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
			RawDescriptor: file_gpt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_gpt_proto_goTypes,
		DependencyIndexes: file_gpt_proto_depIdxs,
		MessageInfos:      file_gpt_proto_msgTypes,
	}.Build()
	File_gpt_proto = out.File
	file_gpt_proto_rawDesc = nil
	file_gpt_proto_goTypes = nil
	file_gpt_proto_depIdxs = nil
}

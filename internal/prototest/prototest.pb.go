// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: internal/prototest/prototest.proto

package prototest

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

type Simple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldA string `protobuf:"bytes,1,opt,name=field_a,json=fieldA,proto3" json:"field_a,omitempty"`
	FieldB int32  `protobuf:"varint,2,opt,name=field_b,json=fieldB,proto3" json:"field_b,omitempty"`
}

func (x *Simple) Reset() {
	*x = Simple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simple) ProtoMessage() {}

func (x *Simple) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simple.ProtoReflect.Descriptor instead.
func (*Simple) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{0}
}

func (x *Simple) GetFieldA() string {
	if x != nil {
		return x.FieldA
	}
	return ""
}

func (x *Simple) GetFieldB() int32 {
	if x != nil {
		return x.FieldB
	}
	return 0
}

type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldOne  string  `protobuf:"bytes,1,opt,name=field_one,json=fieldOne,proto3" json:"field_one,omitempty"`
	SubSimple *Simple `protobuf:"bytes,2,opt,name=sub_simple,json=subSimple,proto3" json:"sub_simple,omitempty"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{1}
}

func (x *Nested) GetFieldOne() string {
	if x != nil {
		return x.FieldOne
	}
	return ""
}

func (x *Nested) GetSubSimple() *Simple {
	if x != nil {
		return x.SubSimple
	}
	return nil
}

type A1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B *B1     `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
	C []int32 `protobuf:"varint,2,rep,packed,name=c,proto3" json:"c,omitempty"`
}

func (x *A1) Reset() {
	*x = A1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A1) ProtoMessage() {}

func (x *A1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use A1.ProtoReflect.Descriptor instead.
func (*A1) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{2}
}

func (x *A1) GetB() *B1 {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *A1) GetC() []int32 {
	if x != nil {
		return x.C
	}
	return nil
}

type B1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	D int32 `protobuf:"varint,1,opt,name=d,proto3" json:"d,omitempty"`
	X int32 `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *B1) Reset() {
	*x = B1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*B1) ProtoMessage() {}

func (x *B1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use B1.ProtoReflect.Descriptor instead.
func (*B1) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{3}
}

func (x *B1) GetD() int32 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *B1) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

type R1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F *A1 `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
}

func (x *R1) Reset() {
	*x = R1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *R1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*R1) ProtoMessage() {}

func (x *R1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use R1.ProtoReflect.Descriptor instead.
func (*R1) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{4}
}

func (x *R1) GetF() *A1 {
	if x != nil {
		return x.F
	}
	return nil
}

type R2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F *B1 `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
}

func (x *R2) Reset() {
	*x = R2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *R2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*R2) ProtoMessage() {}

func (x *R2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use R2.ProtoReflect.Descriptor instead.
func (*R2) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{5}
}

func (x *R2) GetF() *B1 {
	if x != nil {
		return x.F
	}
	return nil
}

type A2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E map[string]int32 `protobuf:"bytes,1,rep,name=e,proto3" json:"e,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *A2) Reset() {
	*x = A2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_prototest_prototest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A2) ProtoMessage() {}

func (x *A2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_prototest_prototest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use A2.ProtoReflect.Descriptor instead.
func (*A2) Descriptor() ([]byte, []int) {
	return file_internal_prototest_prototest_proto_rawDescGZIP(), []int{6}
}

func (x *A2) GetE() map[string]int32 {
	if x != nil {
		return x.E
	}
	return nil
}

var File_internal_prototest_prototest_proto protoreflect.FileDescriptor

var file_internal_prototest_prototest_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x22,
	0x3a, 0x0a, 0x06, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x22, 0x57, 0x0a, 0x06, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x6e, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x75, 0x62, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x22, 0x2f, 0x0a, 0x02, 0x41, 0x31, 0x12, 0x1b, 0x0a, 0x01, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x42, 0x31, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x01, 0x63, 0x22, 0x20, 0x0a, 0x02, 0x42, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x22, 0x21, 0x0a, 0x02, 0x52, 0x31, 0x12, 0x1b, 0x0a,
	0x01, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x31, 0x52, 0x01, 0x66, 0x22, 0x21, 0x0a, 0x02, 0x52, 0x32,
	0x12, 0x1b, 0x0a, 0x01, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x31, 0x52, 0x01, 0x66, 0x22, 0x5e, 0x0a,
	0x02, 0x41, 0x32, 0x12, 0x22, 0x0a, 0x01, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x32, 0x2e, 0x45, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x01, 0x65, 0x1a, 0x34, 0x0a, 0x06, 0x45, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x63,
	0x75, 0x73, 0x69, 0x72, 0x67, 0x65, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x6b, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_prototest_prototest_proto_rawDescOnce sync.Once
	file_internal_prototest_prototest_proto_rawDescData = file_internal_prototest_prototest_proto_rawDesc
)

func file_internal_prototest_prototest_proto_rawDescGZIP() []byte {
	file_internal_prototest_prototest_proto_rawDescOnce.Do(func() {
		file_internal_prototest_prototest_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_prototest_prototest_proto_rawDescData)
	})
	return file_internal_prototest_prototest_proto_rawDescData
}

var file_internal_prototest_prototest_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_prototest_prototest_proto_goTypes = []interface{}{
	(*Simple)(nil), // 0: prototest.Simple
	(*Nested)(nil), // 1: prototest.Nested
	(*A1)(nil),     // 2: prototest.A1
	(*B1)(nil),     // 3: prototest.B1
	(*R1)(nil),     // 4: prototest.R1
	(*R2)(nil),     // 5: prototest.R2
	(*A2)(nil),     // 6: prototest.A2
	nil,            // 7: prototest.A2.EEntry
}
var file_internal_prototest_prototest_proto_depIdxs = []int32{
	0, // 0: prototest.Nested.sub_simple:type_name -> prototest.Simple
	3, // 1: prototest.A1.b:type_name -> prototest.B1
	2, // 2: prototest.R1.f:type_name -> prototest.A1
	3, // 3: prototest.R2.f:type_name -> prototest.B1
	7, // 4: prototest.A2.e:type_name -> prototest.A2.EEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_prototest_prototest_proto_init() }
func file_internal_prototest_prototest_proto_init() {
	if File_internal_prototest_prototest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_prototest_prototest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simple); i {
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
		file_internal_prototest_prototest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
		file_internal_prototest_prototest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*A1); i {
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
		file_internal_prototest_prototest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*B1); i {
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
		file_internal_prototest_prototest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*R1); i {
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
		file_internal_prototest_prototest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*R2); i {
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
		file_internal_prototest_prototest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*A2); i {
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
			RawDescriptor: file_internal_prototest_prototest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_prototest_prototest_proto_goTypes,
		DependencyIndexes: file_internal_prototest_prototest_proto_depIdxs,
		MessageInfos:      file_internal_prototest_prototest_proto_msgTypes,
	}.Build()
	File_internal_prototest_prototest_proto = out.File
	file_internal_prototest_prototest_proto_rawDesc = nil
	file_internal_prototest_prototest_proto_goTypes = nil
	file_internal_prototest_prototest_proto_depIdxs = nil
}

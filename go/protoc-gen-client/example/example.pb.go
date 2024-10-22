// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/example.proto

package example_pb

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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id       int32             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Nested   *Nested           `protobuf:"bytes,4,opt,name=nested,proto3" json:"nested,omitempty"`
	NestedV2 *Example_NestedV2 `protobuf:"bytes,5,opt,name=nestedV2,proto3" json:"nestedV2,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_proto_example_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Example) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Example) GetNested() *Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *Example) GetNestedV2() *Example_NestedV2 {
	if x != nil {
		return x.NestedV2
	}
	return nil
}

type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_proto_msgTypes[1]
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
	return file_proto_example_proto_rawDescGZIP(), []int{1}
}

func (x *Nested) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Example_NestedV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Example_NestedV2) Reset() {
	*x = Example_NestedV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example_NestedV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example_NestedV2) ProtoMessage() {}

func (x *Example_NestedV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example_NestedV2.ProtoReflect.Descriptor instead.
func (*Example_NestedV2) Descriptor() ([]byte, []int) {
	return file_proto_example_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Example_NestedV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_example_proto protoreflect.FileDescriptor

var file_proto_example_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xad,
	0x01, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52,
	0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x56, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x56, 0x32, 0x52, 0x08, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x56, 0x32, 0x1a, 0x1e,
	0x0a, 0x08, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c,
	0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x3e, 0x0a, 0x0e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14,
	0x2e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_example_proto_rawDescOnce sync.Once
	file_proto_example_proto_rawDescData = file_proto_example_proto_rawDesc
)

func file_proto_example_proto_rawDescGZIP() []byte {
	file_proto_example_proto_rawDescOnce.Do(func() {
		file_proto_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_example_proto_rawDescData)
	})
	return file_proto_example_proto_rawDescData
}

var file_proto_example_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_example_proto_goTypes = []interface{}{
	(*Example)(nil),          // 0: example.Example
	(*Nested)(nil),           // 1: example.Nested
	(*Example_NestedV2)(nil), // 2: example.Example.NestedV2
}
var file_proto_example_proto_depIdxs = []int32{
	1, // 0: example.Example.nested:type_name -> example.Nested
	2, // 1: example.Example.nestedV2:type_name -> example.Example.NestedV2
	0, // 2: example.ExampleService.Echo:input_type -> example.Example
	0, // 3: example.ExampleService.Echo:output_type -> example.Example
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_example_proto_init() }
func file_proto_example_proto_init() {
	if File_proto_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_proto_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example_NestedV2); i {
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
			RawDescriptor: file_proto_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_example_proto_goTypes,
		DependencyIndexes: file_proto_example_proto_depIdxs,
		MessageInfos:      file_proto_example_proto_msgTypes,
	}.Build()
	File_proto_example_proto = out.File
	file_proto_example_proto_rawDesc = nil
	file_proto_example_proto_goTypes = nil
	file_proto_example_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.0
// source: processor_message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand         string  `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumberCores   uint32  `protobuf:"varint,3,opt,name=number_cores,json=numberCores,proto3" json:"number_cores,omitempty"`
	NumberThreads uint32  `protobuf:"varint,4,opt,name=number_threads,json=numberThreads,proto3" json:"number_threads,omitempty"`
	MinGhz        float64 `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz        float64 `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
}

func (x *CPU) Reset() {
	*x = CPU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processor_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPU) ProtoMessage() {}

func (x *CPU) ProtoReflect() protoreflect.Message {
	mi := &file_processor_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
func (*CPU) Descriptor() ([]byte, []int) {
	return file_processor_message_proto_rawDescGZIP(), []int{0}
}

func (x *CPU) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CPU) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CPU) GetNumberCores() uint32 {
	if x != nil {
		return x.NumberCores
	}
	return 0
}

func (x *CPU) GetNumberThreads() uint32 {
	if x != nil {
		return x.NumberThreads
	}
	return 0
}

func (x *CPU) GetMinGhz() float64 {
	if x != nil {
		return x.MinGhz
	}
	return 0
}

func (x *CPU) GetMaxGhz() float64 {
	if x != nil {
		return x.MaxGhz
	}
	return 0
}

type GPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand  string  `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz float64 `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz float64 `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory *Memory `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *GPU) Reset() {
	*x = GPU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processor_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPU) ProtoMessage() {}

func (x *GPU) ProtoReflect() protoreflect.Message {
	mi := &file_processor_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPU.ProtoReflect.Descriptor instead.
func (*GPU) Descriptor() ([]byte, []int) {
	return file_processor_message_proto_rawDescGZIP(), []int{1}
}

func (x *GPU) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *GPU) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GPU) GetMinGhz() float64 {
	if x != nil {
		return x.MinGhz
	}
	return 0
}

func (x *GPU) GetMaxGhz() float64 {
	if x != nil {
		return x.MaxGhz
	}
	return 0
}

func (x *GPU) GetMemory() *Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

var File_processor_message_proto protoreflect.FileDescriptor

var file_processor_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65, 0x63, 0x68, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x14, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x6d, 0x69, 0x6e, 0x47, 0x68, 0x7a, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x67,
	0x68, 0x7a, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x47, 0x68, 0x7a,
	0x22, 0x94, 0x01, 0x0a, 0x03, 0x47, 0x50, 0x55, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x47, 0x68, 0x7a, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6d, 0x61,
	0x78, 0x47, 0x68, 0x7a, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x29, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x62, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_processor_message_proto_rawDescOnce sync.Once
	file_processor_message_proto_rawDescData = file_processor_message_proto_rawDesc
)

func file_processor_message_proto_rawDescGZIP() []byte {
	file_processor_message_proto_rawDescOnce.Do(func() {
		file_processor_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_processor_message_proto_rawDescData)
	})
	return file_processor_message_proto_rawDescData
}

var file_processor_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_processor_message_proto_goTypes = []interface{}{
	(*CPU)(nil),    // 0: techschool.pcbook.CPU
	(*GPU)(nil),    // 1: techschool.pcbook.GPU
	(*Memory)(nil), // 2: techschool.pcbook.Memory
}
var file_processor_message_proto_depIdxs = []int32{
	2, // 0: techschool.pcbook.GPU.memory:type_name -> techschool.pcbook.Memory
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_processor_message_proto_init() }
func file_processor_message_proto_init() {
	if File_processor_message_proto != nil {
		return
	}
	file_memory_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_processor_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPU); i {
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
		file_processor_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPU); i {
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
			RawDescriptor: file_processor_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_processor_message_proto_goTypes,
		DependencyIndexes: file_processor_message_proto_depIdxs,
		MessageInfos:      file_processor_message_proto_msgTypes,
	}.Build()
	File_processor_message_proto = out.File
	file_processor_message_proto_rawDesc = nil
	file_processor_message_proto_goTypes = nil
	file_processor_message_proto_depIdxs = nil
}

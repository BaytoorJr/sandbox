// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.26.1
// source: sse.proto

package events

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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_sse_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_sse_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_sse_proto_rawDescGZIP(), []int{0}
}

type PaymentStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Amount        string                 `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentStatus) Reset() {
	*x = PaymentStatus{}
	mi := &file_sse_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentStatus) ProtoMessage() {}

func (x *PaymentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_sse_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentStatus.ProtoReflect.Descriptor instead.
func (*PaymentStatus) Descriptor() ([]byte, []int) {
	return file_sse_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentStatus) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_sse_proto protoreflect.FileDescriptor

var file_sse_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4f, 0x0a, 0x0d,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x46, 0x0a,
	0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0d, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sse_proto_rawDescOnce sync.Once
	file_sse_proto_rawDescData = file_sse_proto_rawDesc
)

func file_sse_proto_rawDescGZIP() []byte {
	file_sse_proto_rawDescOnce.Do(func() {
		file_sse_proto_rawDescData = protoimpl.X.CompressGZIP(file_sse_proto_rawDescData)
	})
	return file_sse_proto_rawDescData
}

var file_sse_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sse_proto_goTypes = []any{
	(*Empty)(nil),         // 0: events.Empty
	(*PaymentStatus)(nil), // 1: events.PaymentStatus
}
var file_sse_proto_depIdxs = []int32{
	0, // 0: events.EventService.StreamEvents:input_type -> events.Empty
	1, // 1: events.EventService.StreamEvents:output_type -> events.PaymentStatus
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sse_proto_init() }
func file_sse_proto_init() {
	if File_sse_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sse_proto_goTypes,
		DependencyIndexes: file_sse_proto_depIdxs,
		MessageInfos:      file_sse_proto_msgTypes,
	}.Build()
	File_sse_proto = out.File
	file_sse_proto_rawDesc = nil
	file_sse_proto_goTypes = nil
	file_sse_proto_depIdxs = nil
}

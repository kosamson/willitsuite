// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: proto/v1/willitsuite.proto

package willitsuitev1

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

type Protocol int32

const (
	Protocol_PROTOCOL_UNSPECIFIED Protocol = 0
	Protocol_PROTOCOL_IP          Protocol = 1
	Protocol_PROTOCOL_TCP         Protocol = 2
	Protocol_PROTOCOL_UDP         Protocol = 3
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "PROTOCOL_UNSPECIFIED",
		1: "PROTOCOL_IP",
		2: "PROTOCOL_TCP",
		3: "PROTOCOL_UDP",
	}
	Protocol_value = map[string]int32{
		"PROTOCOL_UNSPECIFIED": 0,
		"PROTOCOL_IP":          1,
		"PROTOCOL_TCP":         2,
		"PROTOCOL_UDP":         3,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1_willitsuite_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_proto_v1_willitsuite_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1_willitsuite_proto_rawDescGZIP(), []int{0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=proto.v1.Protocol" json:"protocol,omitempty"`
	Address  string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port     int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	mi := &file_proto_v1_willitsuite_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_willitsuite_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_willitsuite_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_PROTOCOL_UNSPECIFIED
}

func (x *ConnectRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	mi := &file_proto_v1_willitsuite_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_willitsuite_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_willitsuite_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_v1_willitsuite_proto protoreflect.FileDescriptor

var file_proto_v1_willitsuite_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x6c, 0x6c, 0x69,
	0x74, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x6e, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x59, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x54, 0x4f,
	0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54,
	0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x54, 0x43, 0x50, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x55, 0x44, 0x50, 0x10, 0x03, 0x32, 0x56, 0x0a, 0x12,
	0x57, 0x69, 0x6c, 0x6c, 0x69, 0x74, 0x53, 0x75, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x73, 0x61, 0x6d, 0x73, 0x6f, 0x6e, 0x2f, 0x77, 0x69, 0x6c, 0x6c,
	0x69, 0x74, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x74, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_willitsuite_proto_rawDescOnce sync.Once
	file_proto_v1_willitsuite_proto_rawDescData = file_proto_v1_willitsuite_proto_rawDesc
)

func file_proto_v1_willitsuite_proto_rawDescGZIP() []byte {
	file_proto_v1_willitsuite_proto_rawDescOnce.Do(func() {
		file_proto_v1_willitsuite_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_willitsuite_proto_rawDescData)
	})
	return file_proto_v1_willitsuite_proto_rawDescData
}

var file_proto_v1_willitsuite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_v1_willitsuite_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1_willitsuite_proto_goTypes = []any{
	(Protocol)(0),           // 0: proto.v1.Protocol
	(*ConnectRequest)(nil),  // 1: proto.v1.ConnectRequest
	(*ConnectResponse)(nil), // 2: proto.v1.ConnectResponse
}
var file_proto_v1_willitsuite_proto_depIdxs = []int32{
	0, // 0: proto.v1.ConnectRequest.protocol:type_name -> proto.v1.Protocol
	1, // 1: proto.v1.WillitSuiteService.Connect:input_type -> proto.v1.ConnectRequest
	2, // 2: proto.v1.WillitSuiteService.Connect:output_type -> proto.v1.ConnectResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_v1_willitsuite_proto_init() }
func file_proto_v1_willitsuite_proto_init() {
	if File_proto_v1_willitsuite_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_willitsuite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_willitsuite_proto_goTypes,
		DependencyIndexes: file_proto_v1_willitsuite_proto_depIdxs,
		EnumInfos:         file_proto_v1_willitsuite_proto_enumTypes,
		MessageInfos:      file_proto_v1_willitsuite_proto_msgTypes,
	}.Build()
	File_proto_v1_willitsuite_proto = out.File
	file_proto_v1_willitsuite_proto_rawDesc = nil
	file_proto_v1_willitsuite_proto_goTypes = nil
	file_proto_v1_willitsuite_proto_depIdxs = nil
}

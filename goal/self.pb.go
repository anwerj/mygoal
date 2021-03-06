// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: proto/self.proto

package goal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First string `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Last  string `protobuf:"bytes,2,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_self_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_proto_self_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_proto_self_proto_rawDescGZIP(), []int{0}
}

func (x *Name) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *Name) GetLast() string {
	if x != nil {
		return x.Last
	}
	return ""
}

type WelcomeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *WelcomeMessage) Reset() {
	*x = WelcomeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_self_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeMessage) ProtoMessage() {}

func (x *WelcomeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_self_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeMessage.ProtoReflect.Descriptor instead.
func (*WelcomeMessage) Descriptor() ([]byte, []int) {
	return file_proto_self_proto_rawDescGZIP(), []int{1}
}

func (x *WelcomeMessage) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_proto_self_proto protoreflect.FileDescriptor

var file_proto_self_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6c, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x57,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x3c, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_self_proto_rawDescOnce sync.Once
	file_proto_self_proto_rawDescData = file_proto_self_proto_rawDesc
)

func file_proto_self_proto_rawDescGZIP() []byte {
	file_proto_self_proto_rawDescOnce.Do(func() {
		file_proto_self_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_self_proto_rawDescData)
	})
	return file_proto_self_proto_rawDescData
}

var file_proto_self_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_self_proto_goTypes = []interface{}{
	(*Name)(nil),           // 0: proto.Name
	(*WelcomeMessage)(nil), // 1: proto.WelcomeMessage
}
var file_proto_self_proto_depIdxs = []int32{
	0, // 0: proto.SelfService.Welcome:input_type -> proto.Name
	1, // 1: proto.SelfService.Welcome:output_type -> proto.WelcomeMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_self_proto_init() }
func file_proto_self_proto_init() {
	if File_proto_self_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_self_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_proto_self_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeMessage); i {
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
			RawDescriptor: file_proto_self_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_self_proto_goTypes,
		DependencyIndexes: file_proto_self_proto_depIdxs,
		MessageInfos:      file_proto_self_proto_msgTypes,
	}.Build()
	File_proto_self_proto = out.File
	file_proto_self_proto_rawDesc = nil
	file_proto_self_proto_goTypes = nil
	file_proto_self_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SelfServiceClient is the client API for SelfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SelfServiceClient interface {
	Welcome(ctx context.Context, in *Name, opts ...grpc.CallOption) (*WelcomeMessage, error)
}

type selfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSelfServiceClient(cc grpc.ClientConnInterface) SelfServiceClient {
	return &selfServiceClient{cc}
}

func (c *selfServiceClient) Welcome(ctx context.Context, in *Name, opts ...grpc.CallOption) (*WelcomeMessage, error) {
	out := new(WelcomeMessage)
	err := c.cc.Invoke(ctx, "/proto.SelfService/Welcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SelfServiceServer is the server API for SelfService service.
type SelfServiceServer interface {
	Welcome(context.Context, *Name) (*WelcomeMessage, error)
}

// UnimplementedSelfServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSelfServiceServer struct {
}

func (*UnimplementedSelfServiceServer) Welcome(context.Context, *Name) (*WelcomeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Welcome not implemented")
}

func RegisterSelfServiceServer(s *grpc.Server, srv SelfServiceServer) {
	s.RegisterService(&_SelfService_serviceDesc, srv)
}

func _SelfService_Welcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfServiceServer).Welcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SelfService/Welcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfServiceServer).Welcome(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

var _SelfService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SelfService",
	HandlerType: (*SelfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Welcome",
			Handler:    _SelfService_Welcome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/self.proto",
}

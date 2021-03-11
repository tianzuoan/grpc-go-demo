// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: tina.proto

//声明 包名

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//参数结构
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tina_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tina_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_tina_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//参数结构
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tina_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_tina_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_tina_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LanguageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *LanguageRequest) Reset() {
	*x = LanguageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tina_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageRequest) ProtoMessage() {}

func (x *LanguageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tina_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageRequest.ProtoReflect.Descriptor instead.
func (*LanguageRequest) Descriptor() ([]byte, []int) {
	return file_tina_proto_rawDescGZIP(), []int{2}
}

func (x *LanguageRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type LanguageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score uint32 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *LanguageReply) Reset() {
	*x = LanguageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tina_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageReply) ProtoMessage() {}

func (x *LanguageReply) ProtoReflect() protoreflect.Message {
	mi := &file_tina_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageReply.ProtoReflect.Descriptor instead.
func (*LanguageReply) Descriptor() ([]byte, []int) {
	return file_tina_proto_rawDescGZIP(), []int{3}
}

func (x *LanguageReply) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_tina_proto protoreflect.FileDescriptor

var file_tina_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x69, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x25, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x32,
	0x8c, 0x01, 0x0a, 0x04, 0x74, 0x69, 0x6e, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x14, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tina_proto_rawDescOnce sync.Once
	file_tina_proto_rawDescData = file_tina_proto_rawDesc
)

func file_tina_proto_rawDescGZIP() []byte {
	file_tina_proto_rawDescOnce.Do(func() {
		file_tina_proto_rawDescData = protoimpl.X.CompressGZIP(file_tina_proto_rawDescData)
	})
	return file_tina_proto_rawDescData
}

var file_tina_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tina_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),    // 0: service.HelloRequest
	(*HelloReply)(nil),      // 1: service.HelloReply
	(*LanguageRequest)(nil), // 2: service.LanguageRequest
	(*LanguageReply)(nil),   // 3: service.LanguageReply
}
var file_tina_proto_depIdxs = []int32{
	0, // 0: service.tina.SayHello:input_type -> service.HelloRequest
	2, // 1: service.tina.LearnForeignLanguage:input_type -> service.LanguageRequest
	1, // 2: service.tina.SayHello:output_type -> service.HelloReply
	3, // 3: service.tina.LearnForeignLanguage:output_type -> service.LanguageReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tina_proto_init() }
func file_tina_proto_init() {
	if File_tina_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tina_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_tina_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_tina_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageRequest); i {
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
		file_tina_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageReply); i {
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
			RawDescriptor: file_tina_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tina_proto_goTypes,
		DependencyIndexes: file_tina_proto_depIdxs,
		MessageInfos:      file_tina_proto_msgTypes,
	}.Build()
	File_tina_proto = out.File
	file_tina_proto_rawDesc = nil
	file_tina_proto_goTypes = nil
	file_tina_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TinaClient is the client API for Tina service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TinaClient interface {
	//
	//声名一个 sayhello 接口
	//它的参数 是 HelloRequest
	//它的返回参数是 HelloReply
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	LearnForeignLanguage(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*LanguageReply, error)
}

type tinaClient struct {
	cc grpc.ClientConnInterface
}

func NewTinaClient(cc grpc.ClientConnInterface) TinaClient {
	return &tinaClient{cc}
}

func (c *tinaClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/service.tina/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinaClient) LearnForeignLanguage(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*LanguageReply, error) {
	out := new(LanguageReply)
	err := c.cc.Invoke(ctx, "/service.tina/LearnForeignLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TinaServer is the server API for Tina service.
type TinaServer interface {
	//
	//声名一个 sayhello 接口
	//它的参数 是 HelloRequest
	//它的返回参数是 HelloReply
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	LearnForeignLanguage(context.Context, *LanguageRequest) (*LanguageReply, error)
}

// UnimplementedTinaServer can be embedded to have forward compatible implementations.
type UnimplementedTinaServer struct {
}

func (*UnimplementedTinaServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedTinaServer) LearnForeignLanguage(context.Context, *LanguageRequest) (*LanguageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LearnForeignLanguage not implemented")
}

func RegisterTinaServer(s *grpc.Server, srv TinaServer) {
	s.RegisterService(&_Tina_serviceDesc, srv)
}

func _Tina_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinaServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.tina/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinaServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tina_LearnForeignLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinaServer).LearnForeignLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.tina/LearnForeignLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinaServer).LearnForeignLanguage(ctx, req.(*LanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tina_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.tina",
	HandlerType: (*TinaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Tina_SayHello_Handler,
		},
		{
			MethodName: "LearnForeignLanguage",
			Handler:    _Tina_LearnForeignLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tina.proto",
}
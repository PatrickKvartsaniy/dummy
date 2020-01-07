// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "authorisationservice.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "authorisationservice.HelloResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x49, 0x2c, 0x2d, 0xc9, 0xc8,
	0x2f, 0xca, 0x2c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0xca, 0x29, 0x29, 0x71, 0xf1, 0x78, 0xa4,
	0xe6, 0xe4, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xf8, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x9a, 0x5c, 0xbc, 0x50,
	0x35, 0xc5, 0x05, 0x20, 0x9d, 0x42, 0x12, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x30,
	0x75, 0x30, 0xae, 0x51, 0x32, 0x17, 0x8f, 0x4b, 0x69, 0x6e, 0x6e, 0x65, 0x30, 0xc4, 0x78, 0xa1,
	0x60, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x6e, 0x21, 0x25, 0x3d, 0x6c, 0x2e, 0xd0, 0x43, 0xb6,
	0x5e, 0x4a, 0x19, 0xaf, 0x1a, 0x88, 0xf5, 0x4e, 0xac, 0x51, 0xcc, 0x89, 0x05, 0x99, 0x49, 0x6c,
	0x60, 0x7f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x52, 0xd9, 0x65, 0xe8, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type dummyServiceClient struct {
	cc *grpc.ClientConn
}

func NewDummyServiceClient(cc *grpc.ClientConn) DummyServiceClient {
	return &dummyServiceClient{cc}
}

func (c *dummyServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/authorisationservice.DummyService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServiceServer is the server API for DummyService service.
type DummyServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedDummyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

func (*UnimplementedDummyServiceServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterDummyServiceServer(s *grpc.Server, srv DummyServiceServer) {
	s.RegisterService(&_DummyService_serviceDesc, srv)
}

func _DummyService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorisationservice.DummyService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DummyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorisationservice.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _DummyService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

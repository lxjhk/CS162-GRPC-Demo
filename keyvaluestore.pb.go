// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyvaluestore.proto

package main

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

// The request message containing the key
type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f733d5cfeb68adf, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// The request message containing the key and value
type StoreRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f733d5cfeb68adf, []int{1}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// The response message containing the value associated with the key
type Response struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f733d5cfeb68adf, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "keyvaluestore.GetRequest")
	proto.RegisterType((*StoreRequest)(nil), "keyvaluestore.StoreRequest")
	proto.RegisterType((*Response)(nil), "keyvaluestore.Response")
}

func init() { proto.RegisterFile("keyvaluestore.proto", fileDescriptor_7f733d5cfeb68adf) }

var fileDescriptor_7f733d5cfeb68adf = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4e, 0xad, 0x2c,
	0x4b, 0xcc, 0x29, 0x4d, 0x2d, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0x92, 0xe3, 0xe2, 0x72, 0x4f, 0x2d, 0x09, 0x4a, 0x2d, 0x04, 0x09, 0x08,
	0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a,
	0x66, 0x5c, 0x3c, 0xc1, 0x20, 0x85, 0x38, 0x55, 0x08, 0x89, 0x70, 0xb1, 0x82, 0xcd, 0x93, 0x60,
	0x02, 0x8b, 0x41, 0x38, 0x4a, 0x0a, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0x08, 0x15, 0x8c, 0x48, 0x2a, 0x8c, 0x56, 0x32, 0x71, 0xf1, 0x7a, 0xa7, 0x56, 0x86, 0x81, 0x38,
	0x60, 0x2b, 0x84, 0x5c, 0xb9, 0x38, 0xdd, 0x53, 0x4b, 0xc0, 0x02, 0xc5, 0x42, 0x92, 0x7a, 0xa8,
	0xae, 0x47, 0xb8, 0x52, 0x4a, 0x1c, 0x4d, 0x0a, 0x66, 0x91, 0x12, 0x83, 0x06, 0xa3, 0x01, 0xa3,
	0x90, 0x27, 0x17, 0x37, 0xd8, 0x3c, 0xa8, 0x41, 0xd2, 0x68, 0xaa, 0x91, 0xbd, 0x43, 0xc8, 0x28,
	0x07, 0x2e, 0x0e, 0x98, 0x8b, 0xc8, 0x73, 0x90, 0x90, 0x0b, 0x17, 0x17, 0xc2, 0x31, 0xe4, 0xba,
	0x25, 0x89, 0x0d, 0x1c, 0x77, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x42, 0xe1, 0xdd,
	0xd2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyValueStoreClient is the client API for KeyValueStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyValueStoreClient interface {
	// Provides a value for each key request
	GetValues(ctx context.Context, opts ...grpc.CallOption) (KeyValueStore_GetValuesClient, error)
	// Stores a value for each key-value request
	StoreValues(ctx context.Context, opts ...grpc.CallOption) (KeyValueStore_StoreValuesClient, error)
	// Provides a value for a key request
	GetValue(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	// Stores a value for a key-value request
	StoreValue(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Response, error)
}

type keyValueStoreClient struct {
	cc *grpc.ClientConn
}

func NewKeyValueStoreClient(cc *grpc.ClientConn) KeyValueStoreClient {
	return &keyValueStoreClient{cc}
}

func (c *keyValueStoreClient) GetValues(ctx context.Context, opts ...grpc.CallOption) (KeyValueStore_GetValuesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeyValueStore_serviceDesc.Streams[0], "/keyvaluestore.KeyValueStore/GetValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyValueStoreGetValuesClient{stream}
	return x, nil
}

type KeyValueStore_GetValuesClient interface {
	Send(*GetRequest) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type keyValueStoreGetValuesClient struct {
	grpc.ClientStream
}

func (x *keyValueStoreGetValuesClient) Send(m *GetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keyValueStoreGetValuesClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keyValueStoreClient) StoreValues(ctx context.Context, opts ...grpc.CallOption) (KeyValueStore_StoreValuesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeyValueStore_serviceDesc.Streams[1], "/keyvaluestore.KeyValueStore/StoreValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &keyValueStoreStoreValuesClient{stream}
	return x, nil
}

type KeyValueStore_StoreValuesClient interface {
	Send(*StoreRequest) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type keyValueStoreStoreValuesClient struct {
	grpc.ClientStream
}

func (x *keyValueStoreStoreValuesClient) Send(m *StoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keyValueStoreStoreValuesClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keyValueStoreClient) GetValue(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/keyvaluestore.KeyValueStore/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) StoreValue(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/keyvaluestore.KeyValueStore/StoreValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueStoreServer is the server API for KeyValueStore service.
type KeyValueStoreServer interface {
	// Provides a value for each key request
	GetValues(KeyValueStore_GetValuesServer) error
	// Stores a value for each key-value request
	StoreValues(KeyValueStore_StoreValuesServer) error
	// Provides a value for a key request
	GetValue(context.Context, *GetRequest) (*Response, error)
	// Stores a value for a key-value request
	StoreValue(context.Context, *StoreRequest) (*Response, error)
}

// UnimplementedKeyValueStoreServer can be embedded to have forward compatible implementations.
type UnimplementedKeyValueStoreServer struct {
}

func (*UnimplementedKeyValueStoreServer) GetValues(srv KeyValueStore_GetValuesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetValues not implemented")
}
func (*UnimplementedKeyValueStoreServer) StoreValues(srv KeyValueStore_StoreValuesServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreValues not implemented")
}
func (*UnimplementedKeyValueStoreServer) GetValue(ctx context.Context, req *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (*UnimplementedKeyValueStoreServer) StoreValue(ctx context.Context, req *StoreRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreValue not implemented")
}

func RegisterKeyValueStoreServer(s *grpc.Server, srv KeyValueStoreServer) {
	s.RegisterService(&_KeyValueStore_serviceDesc, srv)
}

func _KeyValueStore_GetValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeyValueStoreServer).GetValues(&keyValueStoreGetValuesServer{stream})
}

type KeyValueStore_GetValuesServer interface {
	Send(*Response) error
	Recv() (*GetRequest, error)
	grpc.ServerStream
}

type keyValueStoreGetValuesServer struct {
	grpc.ServerStream
}

func (x *keyValueStoreGetValuesServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keyValueStoreGetValuesServer) Recv() (*GetRequest, error) {
	m := new(GetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KeyValueStore_StoreValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeyValueStoreServer).StoreValues(&keyValueStoreStoreValuesServer{stream})
}

type KeyValueStore_StoreValuesServer interface {
	Send(*Response) error
	Recv() (*StoreRequest, error)
	grpc.ServerStream
}

type keyValueStoreStoreValuesServer struct {
	grpc.ServerStream
}

func (x *keyValueStoreStoreValuesServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keyValueStoreStoreValuesServer) Recv() (*StoreRequest, error) {
	m := new(StoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KeyValueStore_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvaluestore.KeyValueStore/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).GetValue(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_StoreValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).StoreValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyvaluestore.KeyValueStore/StoreValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).StoreValue(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyValueStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keyvaluestore.KeyValueStore",
	HandlerType: (*KeyValueStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _KeyValueStore_GetValue_Handler,
		},
		{
			MethodName: "StoreValue",
			Handler:    _KeyValueStore_StoreValue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetValues",
			Handler:       _KeyValueStore_GetValues_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StoreValues",
			Handler:       _KeyValueStore_StoreValues_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "keyvaluestore.proto",
}

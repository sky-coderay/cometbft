// Code generated by protoc-gen-go. DO NOT EDIT.
// source: defs.proto

/*
Package protodb is a generated protocol buffer package.

It is generated from these files:
	defs.proto

It has these top-level messages:
	Entity
	Nothing
	Domain
	Iterator
	Stats
	Init
*/
package protodb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Entity struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Key       []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Exists    bool   `protobuf:"varint,4,opt,name=exists" json:"exists,omitempty"`
	Start     []byte `protobuf:"bytes,5,opt,name=start,proto3" json:"start,omitempty"`
	End       []byte `protobuf:"bytes,6,opt,name=end,proto3" json:"end,omitempty"`
	Err       string `protobuf:"bytes,7,opt,name=err" json:"err,omitempty"`
	CreatedAt int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entity) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Entity) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entity) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Entity) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *Entity) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Entity) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Entity) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Entity) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Nothing struct {
}

func (m *Nothing) Reset()                    { *m = Nothing{} }
func (m *Nothing) String() string            { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()               {}
func (*Nothing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Domain struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Domain) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Domain) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

type Iterator struct {
	Domain *Domain `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Valid  bool    `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Key    []byte  `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte  `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Iterator) Reset()                    { *m = Iterator{} }
func (m *Iterator) String() string            { return proto.CompactTextString(m) }
func (*Iterator) ProtoMessage()               {}
func (*Iterator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Iterator) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *Iterator) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Iterator) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Iterator) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Stats struct {
	Data   map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TimeAt int64             `protobuf:"varint,2,opt,name=time_at,json=timeAt" json:"time_at,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Stats) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Stats) GetTimeAt() int64 {
	if m != nil {
		return m.TimeAt
	}
	return 0
}

type Init struct {
	Type string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Dir  string `protobuf:"bytes,3,opt,name=Dir" json:"Dir,omitempty"`
}

func (m *Init) Reset()                    { *m = Init{} }
func (m *Init) String() string            { return proto.CompactTextString(m) }
func (*Init) ProtoMessage()               {}
func (*Init) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Init) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Init) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Init) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func init() {
	proto.RegisterType((*Entity)(nil), "protodb.Entity")
	proto.RegisterType((*Nothing)(nil), "protodb.Nothing")
	proto.RegisterType((*Domain)(nil), "protodb.Domain")
	proto.RegisterType((*Iterator)(nil), "protodb.Iterator")
	proto.RegisterType((*Stats)(nil), "protodb.Stats")
	proto.RegisterType((*Init)(nil), "protodb.Init")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DB service

type DBClient interface {
	Init(ctx context.Context, in *Init, opts ...grpc.CallOption) (*Entity, error)
	Get(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error)
	GetStream(ctx context.Context, opts ...grpc.CallOption) (DB_GetStreamClient, error)
	Has(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error)
	Set(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error)
	SetSync(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error)
	Delete(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error)
	DeleteSync(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error)
	Iterator(ctx context.Context, in *Entity, opts ...grpc.CallOption) (DB_IteratorClient, error)
	ReverseIterator(ctx context.Context, in *Entity, opts ...grpc.CallOption) (DB_ReverseIteratorClient, error)
	// rpc print(Nothing) returns (Entity) {}
	Stats(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Stats, error)
}

type dBClient struct {
	cc *grpc.ClientConn
}

func NewDBClient(cc *grpc.ClientConn) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) Init(ctx context.Context, in *Init, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := grpc.Invoke(ctx, "/protodb.DB/init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Get(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := grpc.Invoke(ctx, "/protodb.DB/get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) GetStream(ctx context.Context, opts ...grpc.CallOption) (DB_GetStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[0], c.cc, "/protodb.DB/getStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBGetStreamClient{stream}
	return x, nil
}

type DB_GetStreamClient interface {
	Send(*Entity) error
	Recv() (*Entity, error)
	grpc.ClientStream
}

type dBGetStreamClient struct {
	grpc.ClientStream
}

func (x *dBGetStreamClient) Send(m *Entity) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dBGetStreamClient) Recv() (*Entity, error) {
	m := new(Entity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) Has(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := grpc.Invoke(ctx, "/protodb.DB/has", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Set(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/protodb.DB/set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) SetSync(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/protodb.DB/setSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Delete(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/protodb.DB/delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) DeleteSync(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := grpc.Invoke(ctx, "/protodb.DB/deleteSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Iterator(ctx context.Context, in *Entity, opts ...grpc.CallOption) (DB_IteratorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[1], c.cc, "/protodb.DB/iterator", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBIteratorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_IteratorClient interface {
	Recv() (*Iterator, error)
	grpc.ClientStream
}

type dBIteratorClient struct {
	grpc.ClientStream
}

func (x *dBIteratorClient) Recv() (*Iterator, error) {
	m := new(Iterator)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) ReverseIterator(ctx context.Context, in *Entity, opts ...grpc.CallOption) (DB_ReverseIteratorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DB_serviceDesc.Streams[2], c.cc, "/protodb.DB/reverseIterator", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBReverseIteratorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DB_ReverseIteratorClient interface {
	Recv() (*Iterator, error)
	grpc.ClientStream
}

type dBReverseIteratorClient struct {
	grpc.ClientStream
}

func (x *dBReverseIteratorClient) Recv() (*Iterator, error) {
	m := new(Iterator)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dBClient) Stats(ctx context.Context, in *Nothing, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := grpc.Invoke(ctx, "/protodb.DB/stats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DB service

type DBServer interface {
	Init(context.Context, *Init) (*Entity, error)
	Get(context.Context, *Entity) (*Entity, error)
	GetStream(DB_GetStreamServer) error
	Has(context.Context, *Entity) (*Entity, error)
	Set(context.Context, *Entity) (*Nothing, error)
	SetSync(context.Context, *Entity) (*Nothing, error)
	Delete(context.Context, *Entity) (*Nothing, error)
	DeleteSync(context.Context, *Entity) (*Nothing, error)
	Iterator(*Entity, DB_IteratorServer) error
	ReverseIterator(*Entity, DB_ReverseIteratorServer) error
	// rpc print(Nothing) returns (Entity) {}
	Stats(context.Context, *Nothing) (*Stats, error)
}

func RegisterDBServer(s *grpc.Server, srv DBServer) {
	s.RegisterService(&_DB_serviceDesc, srv)
}

func _DB_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Init)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Init(ctx, req.(*Init))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Get(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DBServer).GetStream(&dBGetStreamServer{stream})
}

type DB_GetStreamServer interface {
	Send(*Entity) error
	Recv() (*Entity, error)
	grpc.ServerStream
}

type dBGetStreamServer struct {
	grpc.ServerStream
}

func (x *dBGetStreamServer) Send(m *Entity) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dBGetStreamServer) Recv() (*Entity, error) {
	m := new(Entity)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DB_Has_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Has(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/Has",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Has(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Set(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_SetSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).SetSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/SetSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).SetSync(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Delete(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_DeleteSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).DeleteSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/DeleteSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).DeleteSync(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Iterator_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).Iterator(m, &dBIteratorServer{stream})
}

type DB_IteratorServer interface {
	Send(*Iterator) error
	grpc.ServerStream
}

type dBIteratorServer struct {
	grpc.ServerStream
}

func (x *dBIteratorServer) Send(m *Iterator) error {
	return x.ServerStream.SendMsg(m)
}

func _DB_ReverseIterator_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DBServer).ReverseIterator(m, &dBReverseIteratorServer{stream})
}

type DB_ReverseIteratorServer interface {
	Send(*Iterator) error
	grpc.ServerStream
}

type dBReverseIteratorServer struct {
	grpc.ServerStream
}

func (x *dBReverseIteratorServer) Send(m *Iterator) error {
	return x.ServerStream.SendMsg(m)
}

func _DB_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodb.DB/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Stats(ctx, req.(*Nothing))
	}
	return interceptor(ctx, in, info, handler)
}

var _DB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protodb.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "init",
			Handler:    _DB_Init_Handler,
		},
		{
			MethodName: "get",
			Handler:    _DB_Get_Handler,
		},
		{
			MethodName: "has",
			Handler:    _DB_Has_Handler,
		},
		{
			MethodName: "set",
			Handler:    _DB_Set_Handler,
		},
		{
			MethodName: "setSync",
			Handler:    _DB_SetSync_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _DB_Delete_Handler,
		},
		{
			MethodName: "deleteSync",
			Handler:    _DB_DeleteSync_Handler,
		},
		{
			MethodName: "stats",
			Handler:    _DB_Stats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getStream",
			Handler:       _DB_GetStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "iterator",
			Handler:       _DB_Iterator_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "reverseIterator",
			Handler:       _DB_ReverseIterator_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "defs.proto",
}

func init() { proto.RegisterFile("defs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0xda, 0x8e, 0x13, 0x4f, 0xa1, 0x2d, 0x2b, 0x04, 0xab, 0x4a, 0x48, 0x96, 0x2f, 0x98,
	0x3f, 0x2b, 0xa4, 0x07, 0x7e, 0x4e, 0x14, 0xa5, 0x87, 0x5c, 0x7a, 0x70, 0xb8, 0xa3, 0x6d, 0x3d,
	0xa4, 0x2b, 0x1a, 0xbb, 0xec, 0x0e, 0x15, 0x7e, 0x02, 0x1e, 0x80, 0x27, 0xe1, 0x0d, 0xd1, 0xae,
	0x7f, 0x42, 0x69, 0x0e, 0xe6, 0xe4, 0x99, 0xd9, 0xef, 0xfb, 0x66, 0xf6, 0xf3, 0x2c, 0x40, 0x81,
	0x5f, 0x4c, 0x76, 0xad, 0x2b, 0xaa, 0xf8, 0xc4, 0x7d, 0x8a, 0xf3, 0xe4, 0x37, 0x83, 0xf0, 0xb4,
	0x24, 0x45, 0x35, 0xdf, 0x07, 0x4f, 0x15, 0x82, 0xc5, 0x2c, 0x1d, 0xe7, 0x9e, 0x2a, 0xf8, 0x21,
	0xf8, 0x5f, 0xb1, 0x16, 0x5e, 0xcc, 0xd2, 0x7b, 0xb9, 0x0d, 0xf9, 0x43, 0x18, 0xdf, 0xc8, 0xab,
	0xef, 0x28, 0x7c, 0x57, 0x6b, 0x12, 0xfe, 0x08, 0x42, 0xfc, 0xa1, 0x0c, 0x19, 0x11, 0xc4, 0x2c,
	0x9d, 0xe6, 0x6d, 0x66, 0xd1, 0x86, 0xa4, 0x26, 0x31, 0x6e, 0xd0, 0x2e, 0xb1, 0xaa, 0x58, 0x16,
	0x22, 0x6c, 0x54, 0xb1, 0x74, 0x7d, 0x50, 0x6b, 0x31, 0x89, 0x59, 0x1a, 0xe5, 0x36, 0xe4, 0x4f,
	0x00, 0x2e, 0x34, 0x4a, 0xc2, 0xe2, 0xb3, 0x24, 0x31, 0x8d, 0x59, 0xea, 0xe7, 0x51, 0x5b, 0x39,
	0xa1, 0x24, 0x82, 0xc9, 0x59, 0x45, 0x97, 0xaa, 0x5c, 0x27, 0x33, 0x08, 0x17, 0xd5, 0x46, 0xaa,
	0x72, 0xdb, 0x8d, 0xed, 0xe8, 0xe6, 0xf5, 0xdd, 0x92, 0x6f, 0x30, 0x5d, 0x12, 0x6a, 0x49, 0x95,
	0xe6, 0x4f, 0x21, 0x2c, 0x1c, 0xdb, 0x91, 0xf6, 0xe6, 0x07, 0x59, 0x6b, 0x4b, 0xd6, 0x88, 0xe6,
	0xed, 0x71, 0x7b, 0x71, 0xd5, 0x08, 0x4d, 0xf3, 0x26, 0xe9, 0x0c, 0xf2, 0x77, 0x18, 0x14, 0xfc,
	0x65, 0x50, 0xf2, 0x93, 0xc1, 0x78, 0x45, 0x92, 0x0c, 0x7f, 0x09, 0x41, 0x21, 0x49, 0x0a, 0x16,
	0xfb, 0xe9, 0xde, 0x5c, 0xf4, 0xed, 0xdc, 0x69, 0xb6, 0x90, 0x24, 0x4f, 0x4b, 0xd2, 0x75, 0xee,
	0x50, 0xfc, 0x31, 0x4c, 0x48, 0x6d, 0xd0, 0x7a, 0xe0, 0x39, 0x0f, 0x42, 0x9b, 0x9e, 0xd0, 0xd1,
	0x1b, 0x88, 0x7a, 0x6c, 0x37, 0x05, 0x6b, 0xec, 0xbb, 0x35, 0x85, 0xe7, 0x6a, 0x4d, 0xf2, 0xde,
	0x7b, 0xcb, 0x92, 0x0f, 0x10, 0x2c, 0x4b, 0x45, 0x9c, 0x43, 0xf0, 0xa9, 0xbe, 0xc6, 0x96, 0xe4,
	0x62, 0x5b, 0x3b, 0x93, 0x9b, 0x8e, 0xe4, 0x62, 0xab, 0xbd, 0x50, 0xda, 0xdd, 0x30, 0xca, 0x6d,
	0x38, 0xff, 0x15, 0x80, 0xb7, 0xf8, 0xc8, 0x53, 0x08, 0x94, 0x15, 0xba, 0xdf, 0x5f, 0xc1, 0xea,
	0x1e, 0x6d, 0x0d, 0x6c, 0x76, 0x2a, 0x19, 0xf1, 0x67, 0xe0, 0xaf, 0x91, 0xf8, 0xbf, 0x27, 0xbb,
	0xa0, 0xc7, 0x10, 0xad, 0x91, 0x56, 0xa4, 0x51, 0x6e, 0x86, 0x10, 0x52, 0x36, 0x63, 0x56, 0xff,
	0x52, 0x9a, 0x41, 0xfa, 0xcf, 0xc1, 0x37, 0xbb, 0x46, 0x39, 0xec, 0x0b, 0xdd, 0x5a, 0x8d, 0x78,
	0x06, 0x13, 0x83, 0xb4, 0xaa, 0xcb, 0x8b, 0x61, 0xf8, 0x57, 0x10, 0x16, 0x78, 0x85, 0x84, 0xc3,
	0xe0, 0xaf, 0xed, 0x6b, 0xb4, 0xf0, 0xe1, 0x1d, 0xe6, 0x30, 0x55, 0xdd, 0xe2, 0xde, 0x21, 0x3c,
	0xd8, 0xfe, 0x87, 0x16, 0x93, 0x8c, 0x66, 0x8c, 0xbf, 0x83, 0x03, 0x8d, 0x37, 0xa8, 0x0d, 0x2e,
	0xff, 0x97, 0xfa, 0xc2, 0xbd, 0x27, 0x32, 0xfc, 0xce, 0x2c, 0x47, 0xfb, 0xb7, 0xf7, 0x36, 0x19,
	0x9d, 0x87, 0xae, 0x70, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xfe, 0x6a, 0xcc, 0x63, 0x04,
	0x00, 0x00,
}

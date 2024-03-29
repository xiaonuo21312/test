// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract_service.proto

package pbrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pb "github.com/xuperchain/xupercore/kernel/contract/bridge/pb"
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

func init() { proto.RegisterFile("contract_service.proto", fileDescriptor_e663a77702825514) }

var fileDescriptor_e663a77702825514 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x85, 0x98, 0x18, 0x7d, 0x9b, 0x38, 0x78, 0x12, 0x87, 0x4a, 0x88, 0x01, 0x63, 0x83,
	0x8b, 0x83, 0xc6, 0x95, 0x4b, 0x57, 0xa6, 0x82, 0x40, 0x5d, 0x61, 0x45, 0x93, 0x2a, 0x10, 0x4a,
	0x9c, 0x47, 0x17, 0x9a, 0xc6, 0xc1, 0x7e, 0x2e, 0xdd, 0xd7, 0xe2, 0xb3, 0xf1, 0x01, 0x50, 0x5d,
	0x3b, 0xab, 0x98, 0x93, 0xf6, 0xb2, 0x5b, 0xdb, 0xff, 0xef, 0xff, 0xb3, 0xe3, 0xbc, 0x26, 0xf0,
	0x50, 0xc8, 0x82, 0x54, 0x2c, 0xe8, 0xbb, 0x46, 0x35, 0xcb, 0x04, 0xf2, 0x52, 0x49, 0x92, 0x6c,
	0x6f, 0x2e, 0x2e, 0xe3, 0xac, 0xe0, 0x3e, 0xe6, 0x7a, 0x26, 0xda, 0x0f, 0xaa, 0x6f, 0x16, 0x3a,
	0xfe, 0x73, 0x07, 0xa0, 0x1f, 0x53, 0x36, 0xc3, 0xae, 0x4c, 0x91, 0x5d, 0xc0, 0x56, 0x37, 0xce,
	0x73, 0x76, 0xc8, 0x6f, 0x94, 0xd3, 0x09, 0x77, 0x60, 0x9c, 0xe7, 0x9f, 0xf1, 0x97, 0x41, 0x4d,
	0xed, 0xa3, 0xb5, 0x9c, 0x2e, 0x65, 0xa1, 0x91, 0x7d, 0x80, 0xad, 0x41, 0x56, 0x8c, 0xd9, 0x7e,
	0xb0, 0xb0, 0x88, 0xbc, 0xf2, 0x49, 0x03, 0xb1, 0x94, 0x1d, 0xff, 0x05, 0xd8, 0x3e, 0xbf, 0xd2,
	0x62, 0xb1, 0xd3, 0x3e, 0xb4, 0x06, 0x86, 0xce, 0x92, 0x9f, 0x28, 0x88, 0x3d, 0x0e, 0x77, 0x0d,
	0x79, 0xf9, 0x7e, 0x3d, 0xe0, 0x36, 0xda, 0x87, 0x56, 0x0f, 0x9b, 0x7d, 0x3d, 0x5c, 0xe3, 0xb3,
	0x80, 0xf3, 0x5d, 0xc0, 0xee, 0x5b, 0xcc, 0x91, 0xd0, 0x29, 0x9f, 0x06, 0x1b, 0x4b, 0xc4, 0x5b,
	0x9f, 0x35, 0x32, 0x4e, 0x3c, 0x82, 0x9d, 0x3e, 0xfe, 0x7e, 0x4f, 0xa8, 0x62, 0x92, 0x8a, 0x1d,
	0x04, 0x3b, 0x3e, 0xf6, 0xe6, 0xe7, 0x6b, 0x28, 0xe7, 0x1e, 0xc2, 0xf6, 0x27, 0x83, 0xea, 0x6a,
	0x38, 0x67, 0xe1, 0xbd, 0xb8, 0xd4, 0x6b, 0x0f, 0x9a, 0x21, 0x67, 0xfd, 0x06, 0x60, 0x7f, 0x3a,
	0xc9, 0xa5, 0x98, 0xd4, 0x8c, 0xd8, 0x35, 0xd0, 0x3c, 0x62, 0xab, 0x5c, 0x75, 0xd2, 0xf7, 0x87,
	0x2a, 0x2e, 0xf4, 0x0f, 0xac, 0x3b, 0x0d, 0x1f, 0x37, 0x9f, 0xc6, 0x35, 0xe5, 0xc4, 0x02, 0x76,
	0xbb, 0x0e, 0xb0, 0x7f, 0x8e, 0x17, 0xc1, 0xda, 0x2a, 0xe2, 0x17, 0x78, 0xb9, 0x01, 0xe9, 0x16,
	0x31, 0xc0, 0xba, 0x4a, 0x6a, 0xed, 0x43, 0x7b, 0x81, 0x8c, 0x87, 0x05, 0x37, 0x40, 0xbf, 0x60,
	0xb4, 0x31, 0xef, 0x96, 0x9d, 0xc3, 0x5e, 0x0f, 0xa9, 0x23, 0x84, 0x34, 0x05, 0x75, 0xd2, 0x54,
	0xa1, 0xd6, 0xa8, 0x59, 0x54, 0x37, 0xd7, 0xff, 0x93, 0x7e, 0xe1, 0x57, 0x9b, 0x17, 0x6e, 0xe1,
	0x89, 0xb0, 0x18, 0xd8, 0x81, 0xd4, 0xf4, 0x51, 0x8e, 0x6b, 0x06, 0xd6, 0xa5, 0xcd, 0x03, 0x5b,
	0x41, 0xce, 0xfa, 0x05, 0x76, 0x7a, 0x68, 0x6f, 0x53, 0x47, 0x8d, 0x35, 0x3b, 0xaa, 0xbb, 0x46,
	0x4f, 0x78, 0xfb, 0xa3, 0xf0, 0x5d, 0xf0, 0x9e, 0x11, 0xb4, 0xce, 0x91, 0xce, 0x0c, 0x95, 0x86,
	0x58, 0x78, 0x06, 0xab, 0xdc, 0x2b, 0x0f, 0xd7, 0x61, 0xd5, 0x53, 0xa1, 0x75, 0x3a, 0xcd, 0xe8,
	0x74, 0x86, 0x45, 0x9d, 0xbb, 0xca, 0x9b, 0xdd, 0x2b, 0xd8, 0xd2, 0x7d, 0xf2, 0x15, 0xda, 0x42,
	0x4e, 0x79, 0x12, 0x67, 0xa9, 0xe1, 0x73, 0x53, 0xa2, 0xaa, 0x1a, 0x65, 0xf2, 0xee, 0xee, 0xe8,
	0xcd, 0x38, 0xa3, 0x4b, 0x93, 0x70, 0x21, 0xa7, 0x91, 0x8d, 0xad, 0xd5, 0x7d, 0x94, 0x0a, 0xa3,
	0x09, 0xaa, 0x02, 0xf3, 0xc8, 0x97, 0xa2, 0x44, 0x65, 0xe9, 0x18, 0xa3, 0x32, 0x51, 0xa5, 0x48,
	0xee, 0xd9, 0x17, 0xd2, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x68, 0xbe, 0xa6, 0xcf,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NativeCodeClient is the client API for NativeCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NativeCodeClient interface {
	Call(ctx context.Context, in *pb.NativeCallRequest, opts ...grpc.CallOption) (*pb.NativeCallResponse, error)
	Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error)
}

type nativeCodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNativeCodeClient(cc grpc.ClientConnInterface) NativeCodeClient {
	return &nativeCodeClient{cc}
}

func (c *nativeCodeClient) Call(ctx context.Context, in *pb.NativeCallRequest, opts ...grpc.CallOption) (*pb.NativeCallResponse, error) {
	out := new(pb.NativeCallResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.NativeCode/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nativeCodeClient) Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error) {
	out := new(pb.PingResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.NativeCode/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NativeCodeServer is the server API for NativeCode service.
type NativeCodeServer interface {
	Call(context.Context, *pb.NativeCallRequest) (*pb.NativeCallResponse, error)
	Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error)
}

// UnimplementedNativeCodeServer can be embedded to have forward compatible implementations.
type UnimplementedNativeCodeServer struct {
}

func (*UnimplementedNativeCodeServer) Call(ctx context.Context, req *pb.NativeCallRequest) (*pb.NativeCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedNativeCodeServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterNativeCodeServer(s *grpc.Server, srv NativeCodeServer) {
	s.RegisterService(&_NativeCode_serviceDesc, srv)
}

func _NativeCode_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.NativeCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NativeCodeServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.NativeCode/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NativeCodeServer).Call(ctx, req.(*pb.NativeCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NativeCode_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NativeCodeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.NativeCode/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NativeCodeServer).Ping(ctx, req.(*pb.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NativeCode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xchain.contract.svc.NativeCode",
	HandlerType: (*NativeCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _NativeCode_Call_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _NativeCode_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract_service.proto",
}

// SyscallClient is the client API for Syscall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyscallClient interface {
	// KV service
	PutObject(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (*pb.PutResponse, error)
	GetObject(ctx context.Context, in *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error)
	DeleteObject(ctx context.Context, in *pb.DeleteRequest, opts ...grpc.CallOption) (*pb.DeleteResponse, error)
	NewIterator(ctx context.Context, in *pb.IteratorRequest, opts ...grpc.CallOption) (*pb.IteratorResponse, error)
	// Chain service
	QueryTx(ctx context.Context, in *pb.QueryTxRequest, opts ...grpc.CallOption) (*pb.QueryTxResponse, error)
	QueryBlock(ctx context.Context, in *pb.QueryBlockRequest, opts ...grpc.CallOption) (*pb.QueryBlockResponse, error)
	Transfer(ctx context.Context, in *pb.TransferRequest, opts ...grpc.CallOption) (*pb.TransferResponse, error)
	ContractCall(ctx context.Context, in *pb.ContractCallRequest, opts ...grpc.CallOption) (*pb.ContractCallResponse, error)
	CrossContractQuery(ctx context.Context, in *pb.CrossContractQueryRequest, opts ...grpc.CallOption) (*pb.CrossContractQueryResponse, error)
	GetAccountAddresses(ctx context.Context, in *pb.GetAccountAddressesRequest, opts ...grpc.CallOption) (*pb.GetAccountAddressesResponse, error)
	// Heartbeat
	Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error)
	// Post log
	PostLog(ctx context.Context, in *pb.PostLogRequest, opts ...grpc.CallOption) (*pb.PostLogResponse, error)
	GetCallArgs(ctx context.Context, in *pb.GetCallArgsRequest, opts ...grpc.CallOption) (*pb.CallArgs, error)
	SetOutput(ctx context.Context, in *pb.SetOutputRequest, opts ...grpc.CallOption) (*pb.SetOutputResponse, error)
	// Send Event
	EmitEvent(ctx context.Context, in *pb.EmitEventRequest, opts ...grpc.CallOption) (*pb.EmitEventResponse, error)
}

type syscallClient struct {
	cc grpc.ClientConnInterface
}

func NewSyscallClient(cc grpc.ClientConnInterface) SyscallClient {
	return &syscallClient{cc}
}

func (c *syscallClient) PutObject(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (*pb.PutResponse, error) {
	out := new(pb.PutResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/PutObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) GetObject(ctx context.Context, in *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	out := new(pb.GetResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) DeleteObject(ctx context.Context, in *pb.DeleteRequest, opts ...grpc.CallOption) (*pb.DeleteResponse, error) {
	out := new(pb.DeleteResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) NewIterator(ctx context.Context, in *pb.IteratorRequest, opts ...grpc.CallOption) (*pb.IteratorResponse, error) {
	out := new(pb.IteratorResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/NewIterator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) QueryTx(ctx context.Context, in *pb.QueryTxRequest, opts ...grpc.CallOption) (*pb.QueryTxResponse, error) {
	out := new(pb.QueryTxResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/QueryTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) QueryBlock(ctx context.Context, in *pb.QueryBlockRequest, opts ...grpc.CallOption) (*pb.QueryBlockResponse, error) {
	out := new(pb.QueryBlockResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/QueryBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) Transfer(ctx context.Context, in *pb.TransferRequest, opts ...grpc.CallOption) (*pb.TransferResponse, error) {
	out := new(pb.TransferResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) ContractCall(ctx context.Context, in *pb.ContractCallRequest, opts ...grpc.CallOption) (*pb.ContractCallResponse, error) {
	out := new(pb.ContractCallResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/ContractCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) CrossContractQuery(ctx context.Context, in *pb.CrossContractQueryRequest, opts ...grpc.CallOption) (*pb.CrossContractQueryResponse, error) {
	out := new(pb.CrossContractQueryResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/CrossContractQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) GetAccountAddresses(ctx context.Context, in *pb.GetAccountAddressesRequest, opts ...grpc.CallOption) (*pb.GetAccountAddressesResponse, error) {
	out := new(pb.GetAccountAddressesResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/GetAccountAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) Ping(ctx context.Context, in *pb.PingRequest, opts ...grpc.CallOption) (*pb.PingResponse, error) {
	out := new(pb.PingResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) PostLog(ctx context.Context, in *pb.PostLogRequest, opts ...grpc.CallOption) (*pb.PostLogResponse, error) {
	out := new(pb.PostLogResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/PostLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) GetCallArgs(ctx context.Context, in *pb.GetCallArgsRequest, opts ...grpc.CallOption) (*pb.CallArgs, error) {
	out := new(pb.CallArgs)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/GetCallArgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) SetOutput(ctx context.Context, in *pb.SetOutputRequest, opts ...grpc.CallOption) (*pb.SetOutputResponse, error) {
	out := new(pb.SetOutputResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/SetOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syscallClient) EmitEvent(ctx context.Context, in *pb.EmitEventRequest, opts ...grpc.CallOption) (*pb.EmitEventResponse, error) {
	out := new(pb.EmitEventResponse)
	err := c.cc.Invoke(ctx, "/xchain.contract.svc.Syscall/EmitEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyscallServer is the server API for Syscall service.
type SyscallServer interface {
	// KV service
	PutObject(context.Context, *pb.PutRequest) (*pb.PutResponse, error)
	GetObject(context.Context, *pb.GetRequest) (*pb.GetResponse, error)
	DeleteObject(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error)
	NewIterator(context.Context, *pb.IteratorRequest) (*pb.IteratorResponse, error)
	// Chain service
	QueryTx(context.Context, *pb.QueryTxRequest) (*pb.QueryTxResponse, error)
	QueryBlock(context.Context, *pb.QueryBlockRequest) (*pb.QueryBlockResponse, error)
	Transfer(context.Context, *pb.TransferRequest) (*pb.TransferResponse, error)
	ContractCall(context.Context, *pb.ContractCallRequest) (*pb.ContractCallResponse, error)
	CrossContractQuery(context.Context, *pb.CrossContractQueryRequest) (*pb.CrossContractQueryResponse, error)
	GetAccountAddresses(context.Context, *pb.GetAccountAddressesRequest) (*pb.GetAccountAddressesResponse, error)
	// Heartbeat
	Ping(context.Context, *pb.PingRequest) (*pb.PingResponse, error)
	// Post log
	PostLog(context.Context, *pb.PostLogRequest) (*pb.PostLogResponse, error)
	GetCallArgs(context.Context, *pb.GetCallArgsRequest) (*pb.CallArgs, error)
	SetOutput(context.Context, *pb.SetOutputRequest) (*pb.SetOutputResponse, error)
	// Send Event
	EmitEvent(context.Context, *pb.EmitEventRequest) (*pb.EmitEventResponse, error)
}

// UnimplementedSyscallServer can be embedded to have forward compatible implementations.
type UnimplementedSyscallServer struct {
}

func (*UnimplementedSyscallServer) PutObject(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (*UnimplementedSyscallServer) GetObject(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (*UnimplementedSyscallServer) DeleteObject(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (*UnimplementedSyscallServer) NewIterator(ctx context.Context, req *pb.IteratorRequest) (*pb.IteratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewIterator not implemented")
}
func (*UnimplementedSyscallServer) QueryTx(ctx context.Context, req *pb.QueryTxRequest) (*pb.QueryTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTx not implemented")
}
func (*UnimplementedSyscallServer) QueryBlock(ctx context.Context, req *pb.QueryBlockRequest) (*pb.QueryBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBlock not implemented")
}
func (*UnimplementedSyscallServer) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (*UnimplementedSyscallServer) ContractCall(ctx context.Context, req *pb.ContractCallRequest) (*pb.ContractCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCall not implemented")
}
func (*UnimplementedSyscallServer) CrossContractQuery(ctx context.Context, req *pb.CrossContractQueryRequest) (*pb.CrossContractQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrossContractQuery not implemented")
}
func (*UnimplementedSyscallServer) GetAccountAddresses(ctx context.Context, req *pb.GetAccountAddressesRequest) (*pb.GetAccountAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountAddresses not implemented")
}
func (*UnimplementedSyscallServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedSyscallServer) PostLog(ctx context.Context, req *pb.PostLogRequest) (*pb.PostLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostLog not implemented")
}
func (*UnimplementedSyscallServer) GetCallArgs(ctx context.Context, req *pb.GetCallArgsRequest) (*pb.CallArgs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCallArgs not implemented")
}
func (*UnimplementedSyscallServer) SetOutput(ctx context.Context, req *pb.SetOutputRequest) (*pb.SetOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOutput not implemented")
}
func (*UnimplementedSyscallServer) EmitEvent(ctx context.Context, req *pb.EmitEventRequest) (*pb.EmitEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitEvent not implemented")
}

func RegisterSyscallServer(s *grpc.Server, srv SyscallServer) {
	s.RegisterService(&_Syscall_serviceDesc, srv)
}

func _Syscall_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/PutObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).PutObject(ctx, req.(*pb.PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).GetObject(ctx, req.(*pb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).DeleteObject(ctx, req.(*pb.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_NewIterator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.IteratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).NewIterator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/NewIterator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).NewIterator(ctx, req.(*pb.IteratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_QueryTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.QueryTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).QueryTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/QueryTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).QueryTx(ctx, req.(*pb.QueryTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_QueryBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.QueryBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).QueryBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/QueryBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).QueryBlock(ctx, req.(*pb.QueryBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).Transfer(ctx, req.(*pb.TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_ContractCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.ContractCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).ContractCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/ContractCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).ContractCall(ctx, req.(*pb.ContractCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_CrossContractQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.CrossContractQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).CrossContractQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/CrossContractQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).CrossContractQuery(ctx, req.(*pb.CrossContractQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_GetAccountAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetAccountAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).GetAccountAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/GetAccountAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).GetAccountAddresses(ctx, req.(*pb.GetAccountAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).Ping(ctx, req.(*pb.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_PostLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.PostLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).PostLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/PostLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).PostLog(ctx, req.(*pb.PostLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_GetCallArgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.GetCallArgsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).GetCallArgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/GetCallArgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).GetCallArgs(ctx, req.(*pb.GetCallArgsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_SetOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.SetOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).SetOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/SetOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).SetOutput(ctx, req.(*pb.SetOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Syscall_EmitEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.EmitEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyscallServer).EmitEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xchain.contract.svc.Syscall/EmitEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyscallServer).EmitEvent(ctx, req.(*pb.EmitEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Syscall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xchain.contract.svc.Syscall",
	HandlerType: (*SyscallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutObject",
			Handler:    _Syscall_PutObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Syscall_GetObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Syscall_DeleteObject_Handler,
		},
		{
			MethodName: "NewIterator",
			Handler:    _Syscall_NewIterator_Handler,
		},
		{
			MethodName: "QueryTx",
			Handler:    _Syscall_QueryTx_Handler,
		},
		{
			MethodName: "QueryBlock",
			Handler:    _Syscall_QueryBlock_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Syscall_Transfer_Handler,
		},
		{
			MethodName: "ContractCall",
			Handler:    _Syscall_ContractCall_Handler,
		},
		{
			MethodName: "CrossContractQuery",
			Handler:    _Syscall_CrossContractQuery_Handler,
		},
		{
			MethodName: "GetAccountAddresses",
			Handler:    _Syscall_GetAccountAddresses_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Syscall_Ping_Handler,
		},
		{
			MethodName: "PostLog",
			Handler:    _Syscall_PostLog_Handler,
		},
		{
			MethodName: "GetCallArgs",
			Handler:    _Syscall_GetCallArgs_Handler,
		},
		{
			MethodName: "SetOutput",
			Handler:    _Syscall_SetOutput_Handler,
		},
		{
			MethodName: "EmitEvent",
			Handler:    _Syscall_EmitEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract_service.proto",
}

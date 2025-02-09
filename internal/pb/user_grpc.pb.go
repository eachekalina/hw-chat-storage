// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageStorageClient is the client API for MessageStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageStorageClient interface {
	GetLastMessages(ctx context.Context, in *GetLastMessagesRequest, opts ...grpc.CallOption) (*GetLastMessagesResponse, error)
}

type messageStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageStorageClient(cc grpc.ClientConnInterface) MessageStorageClient {
	return &messageStorageClient{cc}
}

func (c *messageStorageClient) GetLastMessages(ctx context.Context, in *GetLastMessagesRequest, opts ...grpc.CallOption) (*GetLastMessagesResponse, error) {
	out := new(GetLastMessagesResponse)
	err := c.cc.Invoke(ctx, "/storage.MessageStorage/GetLastMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageStorageServer is the server API for MessageStorage service.
// All implementations must embed UnimplementedMessageStorageServer
// for forward compatibility
type MessageStorageServer interface {
	GetLastMessages(context.Context, *GetLastMessagesRequest) (*GetLastMessagesResponse, error)
	mustEmbedUnimplementedMessageStorageServer()
}

// UnimplementedMessageStorageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageStorageServer struct {
}

func (UnimplementedMessageStorageServer) GetLastMessages(context.Context, *GetLastMessagesRequest) (*GetLastMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastMessages not implemented")
}
func (UnimplementedMessageStorageServer) mustEmbedUnimplementedMessageStorageServer() {}

// UnsafeMessageStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageStorageServer will
// result in compilation errors.
type UnsafeMessageStorageServer interface {
	mustEmbedUnimplementedMessageStorageServer()
}

func RegisterMessageStorageServer(s grpc.ServiceRegistrar, srv MessageStorageServer) {
	s.RegisterService(&MessageStorage_ServiceDesc, srv)
}

func _MessageStorage_GetLastMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStorageServer).GetLastMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.MessageStorage/GetLastMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStorageServer).GetLastMessages(ctx, req.(*GetLastMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageStorage_ServiceDesc is the grpc.ServiceDesc for MessageStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.MessageStorage",
	HandlerType: (*MessageStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastMessages",
			Handler:    _MessageStorage_GetLastMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

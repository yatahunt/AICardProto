// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: AICardProto/AICardProto.proto

package AIPokerProto

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

// GameServicesClient is the client API for GameServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServicesClient interface {
	Interact(ctx context.Context, in *GameInteraction, opts ...grpc.CallOption) (*GameHistory, error)
	Join(ctx context.Context, in *PlayerDefinition, opts ...grpc.CallOption) (*PlayerInscription, error)
}

type gameServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServicesClient(cc grpc.ClientConnInterface) GameServicesClient {
	return &gameServicesClient{cc}
}

func (c *gameServicesClient) Interact(ctx context.Context, in *GameInteraction, opts ...grpc.CallOption) (*GameHistory, error) {
	out := new(GameHistory)
	err := c.cc.Invoke(ctx, "/AIPokerProto.GameServices/Interact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServicesClient) Join(ctx context.Context, in *PlayerDefinition, opts ...grpc.CallOption) (*PlayerInscription, error) {
	out := new(PlayerInscription)
	err := c.cc.Invoke(ctx, "/AIPokerProto.GameServices/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServicesServer is the server API for GameServices service.
// All implementations must embed UnimplementedGameServicesServer
// for forward compatibility
type GameServicesServer interface {
	Interact(context.Context, *GameInteraction) (*GameHistory, error)
	Join(context.Context, *PlayerDefinition) (*PlayerInscription, error)
	mustEmbedUnimplementedGameServicesServer()
}

// UnimplementedGameServicesServer must be embedded to have forward compatible implementations.
type UnimplementedGameServicesServer struct {
}

func (UnimplementedGameServicesServer) Interact(context.Context, *GameInteraction) (*GameHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interact not implemented")
}
func (UnimplementedGameServicesServer) Join(context.Context, *PlayerDefinition) (*PlayerInscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedGameServicesServer) mustEmbedUnimplementedGameServicesServer() {}

// UnsafeGameServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServicesServer will
// result in compilation errors.
type UnsafeGameServicesServer interface {
	mustEmbedUnimplementedGameServicesServer()
}

func RegisterGameServicesServer(s grpc.ServiceRegistrar, srv GameServicesServer) {
	s.RegisterService(&GameServices_ServiceDesc, srv)
}

func _GameServices_Interact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInteraction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServicesServer).Interact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AIPokerProto.GameServices/Interact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServicesServer).Interact(ctx, req.(*GameInteraction))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServices_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServicesServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AIPokerProto.GameServices/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServicesServer).Join(ctx, req.(*PlayerDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

// GameServices_ServiceDesc is the grpc.ServiceDesc for GameServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AIPokerProto.GameServices",
	HandlerType: (*GameServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Interact",
			Handler:    _GameServices_Interact_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _GameServices_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AICardProto/AICardProto.proto",
}
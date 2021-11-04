// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// StartServerClient is the client API for StartServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StartServerClient interface {
	AgregarJugador(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Status, error)
	Juego1(ctx context.Context, in *Playermove, opts ...grpc.CallOption) (*Status, error)
	Siguientejuego(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Nextgame, error)
	PedirPozo(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Amount, error)
	SeSolicitoPozo(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Amount, error)
	EstadoLider(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Status, error)
	MandarALider(ctx context.Context, in *Playermove, opts ...grpc.CallOption) (*Status, error)
}

type startServerClient struct {
	cc grpc.ClientConnInterface
}

func NewStartServerClient(cc grpc.ClientConnInterface) StartServerClient {
	return &startServerClient{cc}
}

func (c *startServerClient) AgregarJugador(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/lider.StartServer/AgregarJugador", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) Juego1(ctx context.Context, in *Playermove, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/lider.StartServer/Juego1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) Siguientejuego(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Nextgame, error) {
	out := new(Nextgame)
	err := c.cc.Invoke(ctx, "/lider.StartServer/Siguientejuego", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) PedirPozo(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Amount, error) {
	out := new(Amount)
	err := c.cc.Invoke(ctx, "/lider.StartServer/PedirPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) SeSolicitoPozo(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Amount, error) {
	out := new(Amount)
	err := c.cc.Invoke(ctx, "/lider.StartServer/SeSolicitoPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) EstadoLider(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/lider.StartServer/EstadoLider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startServerClient) MandarALider(ctx context.Context, in *Playermove, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/lider.StartServer/MandarALider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartServerServer is the server API for StartServer service.
// All implementations must embed UnimplementedStartServerServer
// for forward compatibility
type StartServerServer interface {
	AgregarJugador(context.Context, *Name) (*Status, error)
	Juego1(context.Context, *Playermove) (*Status, error)
	Siguientejuego(context.Context, *Name) (*Nextgame, error)
	PedirPozo(context.Context, *Msg) (*Amount, error)
	SeSolicitoPozo(context.Context, *Name) (*Amount, error)
	EstadoLider(context.Context, *Name) (*Status, error)
	MandarALider(context.Context, *Playermove) (*Status, error)
	mustEmbedUnimplementedStartServerServer()
}

// UnimplementedStartServerServer must be embedded to have forward compatible implementations.
type UnimplementedStartServerServer struct {
}

func (UnimplementedStartServerServer) AgregarJugador(context.Context, *Name) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgregarJugador not implemented")
}
func (UnimplementedStartServerServer) Juego1(context.Context, *Playermove) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Juego1 not implemented")
}
func (UnimplementedStartServerServer) Siguientejuego(context.Context, *Name) (*Nextgame, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Siguientejuego not implemented")
}
func (UnimplementedStartServerServer) PedirPozo(context.Context, *Msg) (*Amount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirPozo not implemented")
}
func (UnimplementedStartServerServer) SeSolicitoPozo(context.Context, *Name) (*Amount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeSolicitoPozo not implemented")
}
func (UnimplementedStartServerServer) EstadoLider(context.Context, *Name) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoLider not implemented")
}
func (UnimplementedStartServerServer) MandarALider(context.Context, *Playermove) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MandarALider not implemented")
}
func (UnimplementedStartServerServer) mustEmbedUnimplementedStartServerServer() {}

// UnsafeStartServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StartServerServer will
// result in compilation errors.
type UnsafeStartServerServer interface {
	mustEmbedUnimplementedStartServerServer()
}

func RegisterStartServerServer(s grpc.ServiceRegistrar, srv StartServerServer) {
	s.RegisterService(&StartServer_ServiceDesc, srv)
}

func _StartServer_AgregarJugador_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).AgregarJugador(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/AgregarJugador",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).AgregarJugador(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_Juego1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playermove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).Juego1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/Juego1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).Juego1(ctx, req.(*Playermove))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_Siguientejuego_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).Siguientejuego(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/Siguientejuego",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).Siguientejuego(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_PedirPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).PedirPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/PedirPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).PedirPozo(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_SeSolicitoPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).SeSolicitoPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/SeSolicitoPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).SeSolicitoPozo(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_EstadoLider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).EstadoLider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/EstadoLider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).EstadoLider(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartServer_MandarALider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playermove)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServerServer).MandarALider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lider.StartServer/MandarALider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServerServer).MandarALider(ctx, req.(*Playermove))
	}
	return interceptor(ctx, in, info, handler)
}

// StartServer_ServiceDesc is the grpc.ServiceDesc for StartServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StartServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lider.StartServer",
	HandlerType: (*StartServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AgregarJugador",
			Handler:    _StartServer_AgregarJugador_Handler,
		},
		{
			MethodName: "Juego1",
			Handler:    _StartServer_Juego1_Handler,
		},
		{
			MethodName: "Siguientejuego",
			Handler:    _StartServer_Siguientejuego_Handler,
		},
		{
			MethodName: "PedirPozo",
			Handler:    _StartServer_PedirPozo_Handler,
		},
		{
			MethodName: "SeSolicitoPozo",
			Handler:    _StartServer_SeSolicitoPozo_Handler,
		},
		{
			MethodName: "EstadoLider",
			Handler:    _StartServer_EstadoLider_Handler,
		},
		{
			MethodName: "MandarALider",
			Handler:    _StartServer_MandarALider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lider.proto",
}

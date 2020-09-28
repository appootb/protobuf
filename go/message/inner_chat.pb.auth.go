// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: inner_chat.proto
package message

import (
	"context"
	"net/http"

	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
	"github.com/appootb/protobuf/go/webstream"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = http.StatusOK
var _ = runtime.String
var _ = context.TODO()
var _ = grpc.ServiceDesc{}
var _ = webstream.WebStream{}
var _ = websocket.UnknownFrame
var _ = permission.Subject_NONE
var _ = service.UnaryServerInterceptor

var _innerChatServiceSubjects = map[string][]permission.Subject{
	"/appootb.message.InnerChat/Deliver": {
		permission.Subject_SERVER,
	},
	"/appootb.message.InnerChat/Kick": {
		permission.Subject_SERVER,
	},
}

type wrapperInnerChatServer struct {
	InnerChatServer
	service.Implementor
}

func (w *wrapperInnerChatServer) Deliver(ctx context.Context, req *Package) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerChatServer.Deliver(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerChatServer,
		FullMethod: "/appootb.message.InnerChat/Deliver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerChatServer.Deliver(ctx, req.(*Package))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

func (w *wrapperInnerChatServer) Kick(ctx context.Context, req *Connections) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerChatServer.Kick(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerChatServer,
		FullMethod: "/appootb.message.InnerChat/Kick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerChatServer.Kick(ctx, req.(*Connections))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

// Register scoped server.
func RegisterInnerChatScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv InnerChatServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _innerChatServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_SERVER) {
		RegisterInnerChatServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperInnerChatServer{
		InnerChatServer: srv,
		Implementor:     impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_SERVER) {
		// Register gateway handler.
		if err := RegisterInnerChatHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
	}

	return nil
}

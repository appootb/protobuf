// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: auth.proto
package account

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

var _authServiceSubjects = map[string][]permission.Subject{
	"/appootb.account.Auth/Login": {
		permission.Subject_NONE,
	},
	"/appootb.account.Auth/OAuth": {
		permission.Subject_NONE,
	},
	"/appootb.account.Auth/Refresh": {
		permission.Subject_WEB,
		permission.Subject_PC,
		permission.Subject_MOBILE,
	},
}

type wrapperAuthServer struct {
	AuthServer
	service.Implementor
}

func (w *wrapperAuthServer) Login(ctx context.Context, req *LoginRequest) (*Info, error) {
	if w.UnaryInterceptor() == nil {
		return w.AuthServer.Login(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.AuthServer,
		FullMethod: "/appootb.account.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.AuthServer.Login(ctx, req.(*LoginRequest))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Info), nil
}

func (w *wrapperAuthServer) OAuth(ctx context.Context, req *OAuthRequest) (*Info, error) {
	if w.UnaryInterceptor() == nil {
		return w.AuthServer.OAuth(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.AuthServer,
		FullMethod: "/appootb.account.Auth/OAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.AuthServer.OAuth(ctx, req.(*OAuthRequest))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Info), nil
}

func (w *wrapperAuthServer) Refresh(ctx context.Context, req *empty.Empty) (*Info, error) {
	if w.UnaryInterceptor() == nil {
		return w.AuthServer.Refresh(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.AuthServer,
		FullMethod: "/appootb.account.Auth/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.AuthServer.Refresh(ctx, req.(*empty.Empty))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Info), nil
}

// Register scoped server.
func RegisterAuthScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv AuthServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _authServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_CLIENT) {
		RegisterAuthServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperAuthServer{
		AuthServer:  srv,
		Implementor: impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_CLIENT) {
		// Register gateway handler.
		if err := RegisterAuthHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
	}

	return nil
}

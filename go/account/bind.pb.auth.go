// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: bind.proto
package account

import (
	"context"

	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO()
var _ = grpc.ServiceDesc{}
var _ = permission.TokenLevel_NONE_TOKEN
var _ = service.UnaryServerInterceptor

var _levelBind = map[string]permission.TokenLevel{
	"/appootb.account.Bind/Apply":  permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Bind/Cancel": permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Bind/Gets":   permission.TokenLevel_MIDDLE_TOKEN,
}

type wrapperBindServer struct {
	BindServer
	service.Implementor
}

func (w *wrapperBindServer) Apply(ctx context.Context, req *BindRequest) (*empty.Empty, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.BindServer.Apply(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.BindServer,
		FullMethod: "/appootb.account.Bind/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.BindServer.Apply(ctx, req.(*BindRequest))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

func (w *wrapperBindServer) Cancel(ctx context.Context, req *BindRequest) (*empty.Empty, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.BindServer.Cancel(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.BindServer,
		FullMethod: "/appootb.account.Bind/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.BindServer.Cancel(ctx, req.(*BindRequest))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

func (w *wrapperBindServer) Gets(ctx context.Context, req *empty.Empty) (*AccountBinds, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.BindServer.Gets(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.BindServer,
		FullMethod: "/appootb.account.Bind/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.BindServer.Gets(ctx, req.(*empty.Empty))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*AccountBinds), nil
}

// Register scoped server.
func RegisterBindScopeServer(auth service.Authenticator, impl service.Implementor, srv BindServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelBind)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetScopedGRPCServer(permission.VisibleScope_DEFAULT_SCOPE) {
		RegisterBindServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperBindServer{
		BindServer:  srv,
		Implementor: impl,
	}
	for _, mux := range impl.GetScopedGatewayMux(permission.VisibleScope_DEFAULT_SCOPE) {
		err := RegisterBindHandlerServer(impl.Context(), mux, &wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

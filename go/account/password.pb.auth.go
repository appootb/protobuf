// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: password.proto
package account

import (
	"context"

	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = context.TODO()
var _ = grpc.ServiceDesc{}
var _ = permission.TokenLevel_NONE_TOKEN
var _ = service.UnaryServerInterceptor

var _levelPassword = map[string]permission.TokenLevel{
	"/appootb.account.Password/Reset":  permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Password/Set":    permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Password/Update": permission.TokenLevel_MIDDLE_TOKEN,
}

type wrapperPasswordServer struct {
	PasswordServer
	service.Implementor
}

func (w *wrapperPasswordServer) Set(ctx context.Context, req *PasswordRequest) (*Secret, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.PasswordServer.Set(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.PasswordServer,
		FullMethod: "/appootb.account.Password/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.PasswordServer.Set(ctx, req.(*PasswordRequest))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Secret), nil
}

func (w *wrapperPasswordServer) Update(ctx context.Context, req *PasswordRequest) (*Secret, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.PasswordServer.Update(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.PasswordServer,
		FullMethod: "/appootb.account.Password/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.PasswordServer.Update(ctx, req.(*PasswordRequest))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Secret), nil
}

func (w *wrapperPasswordServer) Reset(ctx context.Context, req *PasswordRequest) (*Secret, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.PasswordServer.Reset(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.PasswordServer,
		FullMethod: "/appootb.account.Password/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.PasswordServer.Reset(ctx, req.(*PasswordRequest))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Secret), nil
}

// Register scoped server.
func RegisterPasswordScopeServer(auth service.Authenticator, impl service.Implementor, srv PasswordServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelPassword)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetScopedGRPCServer(permission.VisibleScope_DEFAULT_SCOPE) {
		RegisterPasswordServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperPasswordServer{
		PasswordServer: srv,
		Implementor:    impl,
	}
	for _, mux := range impl.GetScopedGatewayMux(permission.VisibleScope_DEFAULT_SCOPE) {
		err := RegisterPasswordHandlerServer(impl.Context(), mux, &wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

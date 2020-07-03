// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: profile.proto
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
var _ = permission.Audience_NONE
var _ = service.UnaryServerInterceptor

var _levelProfile = map[string][]permission.Audience{
	"/appootb.account.Profile/Get": {
		permission.Audience_PC,
		permission.Audience_MOBILE,
		permission.Audience_WEB,
	},
	"/appootb.account.Profile/Gets": {
		permission.Audience_WEB,
		permission.Audience_PC,
		permission.Audience_MOBILE,
	},
	"/appootb.account.Profile/Set": {
		permission.Audience_WEB,
		permission.Audience_PC,
		permission.Audience_MOBILE,
	},
}

type wrapperProfileServer struct {
	ProfileServer
	service.Implementor
}

func (w *wrapperProfileServer) Set(ctx context.Context, req *Property) (*empty.Empty, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.ProfileServer.Set(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.ProfileServer,
		FullMethod: "/appootb.account.Profile/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.ProfileServer.Set(ctx, req.(*Property))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

func (w *wrapperProfileServer) Get(ctx context.Context, req *Property) (*Property, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.ProfileServer.Get(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.ProfileServer,
		FullMethod: "/appootb.account.Profile/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.ProfileServer.Get(ctx, req.(*Property))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Property), nil
}

func (w *wrapperProfileServer) Gets(ctx context.Context, req *empty.Empty) (*Properties, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.ProfileServer.Gets(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.ProfileServer,
		FullMethod: "/appootb.account.Profile/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.ProfileServer.Gets(ctx, req.(*empty.Empty))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Properties), nil
}

// Register scoped server.
func RegisterProfileScopeServer(auth service.Authenticator, impl service.Implementor, srv ProfileServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelProfile)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetScopedGRPCServer(permission.VisibleScope_CLIENT) {
		RegisterProfileServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperProfileServer{
		ProfileServer: srv,
		Implementor:   impl,
	}
	for _, mux := range impl.GetScopedGatewayMux(permission.VisibleScope_CLIENT) {
		err := RegisterProfileHandlerServer(impl.Context(), mux, &wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: inner_filter.proto
package strainer

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

var _levelInnerFilter = map[string][]permission.Audience{
	"/appootb.strainer.InnerFilter/Filter": {
		permission.Audience_SERVER,
	},
}

type wrapperInnerFilterServer struct {
	InnerFilterServer
	service.Implementor
}

func (w *wrapperInnerFilterServer) Filter(ctx context.Context, req *FilterRequest) (*empty.Empty, error) {
	if w.UnaryServerInterceptor() == nil {
		return w.InnerFilterServer.Filter(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerFilterServer,
		FullMethod: "/appootb.strainer.InnerFilter/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerFilterServer.Filter(ctx, req.(*FilterRequest))
	}
	resp, err := w.UnaryServerInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

// Register scoped server.
func RegisterInnerFilterScopeServer(auth service.Authenticator, impl service.Implementor, srv InnerFilterServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelInnerFilter)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetScopedGRPCServer(permission.VisibleScope_SERVER) {
		RegisterInnerFilterServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperInnerFilterServer{
		InnerFilterServer: srv,
		Implementor:       impl,
	}
	for _, mux := range impl.GetScopedGatewayMux(permission.VisibleScope_SERVER) {
		err := RegisterInnerFilterHandlerServer(impl.Context(), mux, &wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

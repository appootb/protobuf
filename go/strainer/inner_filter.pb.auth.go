// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: inner_filter.proto
package strainer

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

var _innerFilterServiceSubjects = map[string][]permission.Subject{
	"/appootb.strainer.InnerFilter/Filter": {
		permission.Subject_SERVER,
	},
}

type wrapperInnerFilterServer struct {
	InnerFilterServer
	service.Implementor
}

func (w *wrapperInnerFilterServer) Filter(ctx context.Context, req *FilterRequest) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerFilterServer.Filter(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerFilterServer,
		FullMethod: "/appootb.strainer.InnerFilter/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerFilterServer.Filter(ctx, req.(*FilterRequest))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

// Register scoped server.
func RegisterInnerFilterScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv InnerFilterServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _innerFilterServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_SERVER) {
		RegisterInnerFilterServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperInnerFilterServer{
		InnerFilterServer: srv,
		Implementor:       impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_SERVER) {
		// Register gateway handler.
		if err := RegisterInnerFilterHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
	}

	return nil
}

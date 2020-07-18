// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: inner_code.proto
package captcha

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
var _ = permission.Subject_NONE
var _ = service.UnaryServerInterceptor

var _innerCodeServiceSubjects = map[string][]permission.Subject{
	"/appootb.captcha.InnerCode/Verify": {
		permission.Subject_SERVER,
	},
}

type wrapperInnerCodeServer struct {
	InnerCodeServer
	service.Implementor
}

func (w *wrapperInnerCodeServer) Verify(ctx context.Context, req *VerifyRequest) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerCodeServer.Verify(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerCodeServer,
		FullMethod: "/appootb.captcha.InnerCode/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerCodeServer.Verify(ctx, req.(*VerifyRequest))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

// Register scoped server.
func RegisterInnerCodeScopeServer(auth service.Authenticator, impl service.Implementor, srv InnerCodeServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(_innerCodeServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_SERVER) {
		RegisterInnerCodeServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperInnerCodeServer{
		InnerCodeServer: srv,
		Implementor:     impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_SERVER) {
		err := RegisterInnerCodeHandlerServer(impl.Context(), mux, &wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

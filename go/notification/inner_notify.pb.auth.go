// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: inner_notify.proto
package notification

import (
	"context"
	"net/http"

	"github.com/appootb/protobuf/go/common"
	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
	"github.com/appootb/protobuf/go/webstream"
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

var _innerNotifyServiceSubjects = map[string][]permission.Subject{
	"/appootb.notification.InnerNotify/Post": {
		permission.Subject_SERVER,
	},
}

type wrapperInnerNotifyServer struct {
	InnerNotifyServer
	service.Implementor
}

func (w *wrapperInnerNotifyServer) Post(ctx context.Context, req *Notification) (*common.UUID, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerNotifyServer.Post(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerNotifyServer,
		FullMethod: "/appootb.notification.InnerNotify/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerNotifyServer.Post(ctx, req.(*Notification))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*common.UUID), nil
}

// Register scoped server.
func RegisterInnerNotifyScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv InnerNotifyServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _innerNotifyServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_SERVER) {
		RegisterInnerNotifyServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperInnerNotifyServer{
		InnerNotifyServer: srv,
		Implementor:       impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_SERVER) {
		// Register gateway handler.
		if err := RegisterInnerNotifyHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
	}

	return nil
}
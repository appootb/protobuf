// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: channel.proto
package relation

import (
	"context"
	"net/http"

	"github.com/appootb/protobuf/go/common"
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

var _channelServiceSubjects = map[string][]permission.Subject{
	"/appootb.relation.Channel/Add": {
		permission.Subject_WEB,
		permission.Subject_PC,
		permission.Subject_MOBILE,
	},
	"/appootb.relation.Channel/Cancel": {
		permission.Subject_WEB,
		permission.Subject_PC,
		permission.Subject_MOBILE,
	},
}

type wrapperChannelServer struct {
	ChannelServer
	service.Implementor
}

func (w *wrapperChannelServer) Add(ctx context.Context, req *ApplyRequest) (*ChannelInfo, error) {
	if w.UnaryInterceptor() == nil {
		return w.ChannelServer.Add(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.ChannelServer,
		FullMethod: "/appootb.relation.Channel/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.ChannelServer.Add(ctx, req.(*ApplyRequest))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*ChannelInfo), nil
}

func (w *wrapperChannelServer) Cancel(ctx context.Context, req *common.UniqueId) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.ChannelServer.Cancel(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.ChannelServer,
		FullMethod: "/appootb.relation.Channel/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.ChannelServer.Cancel(ctx, req.(*common.UniqueId))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

// Register scoped server.
func RegisterChannelScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv ChannelServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _channelServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_CLIENT) {
		RegisterChannelServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperChannelServer{
		ChannelServer: srv,
		Implementor:   impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_CLIENT) {
		// Register gateway handler.
		if err := RegisterChannelHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
	}

	return nil
}

// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: inner_hash.proto
package counter

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

var _innerHashServiceSubjects = map[string][]permission.Subject{
	"/appootb.counter.InnerHash/Get": {
		permission.Subject_SERVER,
	},
	"/appootb.counter.InnerHash/Gets": {
		permission.Subject_SERVER,
	},
	"/appootb.counter.InnerHash/Increase": {
		permission.Subject_SERVER,
	},
	"/appootb.counter.InnerHash/MultiGets": {
		permission.Subject_SERVER,
	},
	"/appootb.counter.InnerHash/MultiIncrease": {
		permission.Subject_SERVER,
	},
	"/appootb.counter.InnerHash/Set": {
		permission.Subject_SERVER,
	},
	"/appootb.counter.InnerHash/Sets": {
		permission.Subject_SERVER,
	},
}

type wrapperInnerHashServer struct {
	InnerHashServer
	service.Implementor
}

func (w *wrapperInnerHashServer) Increase(ctx context.Context, req *HashField) (*HashField, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.Increase(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/Increase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.Increase(ctx, req.(*HashField))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*HashField), nil
}

func (w *wrapperInnerHashServer) MultiIncrease(ctx context.Context, req *HashFields) (*HashFields, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.MultiIncrease(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/MultiIncrease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.MultiIncrease(ctx, req.(*HashFields))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*HashFields), nil
}

func (w *wrapperInnerHashServer) Get(ctx context.Context, req *HashField) (*HashField, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.Get(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.Get(ctx, req.(*HashField))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*HashField), nil
}

func (w *wrapperInnerHashServer) Gets(ctx context.Context, req *HashFields) (*HashFields, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.Gets(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/Gets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.Gets(ctx, req.(*HashFields))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*HashFields), nil
}

func (w *wrapperInnerHashServer) Set(ctx context.Context, req *HashField) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.Set(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.Set(ctx, req.(*HashField))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

func (w *wrapperInnerHashServer) Sets(ctx context.Context, req *HashFields) (*empty.Empty, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.Sets(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/Sets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.Sets(ctx, req.(*HashFields))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*empty.Empty), nil
}

func (w *wrapperInnerHashServer) MultiGets(ctx context.Context, req *HashKeys) (*HashValues, error) {
	if w.UnaryInterceptor() == nil {
		return w.InnerHashServer.MultiGets(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.InnerHashServer,
		FullMethod: "/appootb.counter.InnerHash/MultiGets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.InnerHashServer.MultiGets(ctx, req.(*HashKeys))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*HashValues), nil
}

// Register scoped server.
func RegisterInnerHashScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv InnerHashServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _innerHashServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_SERVER) {
		RegisterInnerHashServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperInnerHashServer{
		InnerHashServer: srv,
		Implementor:     impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_SERVER) {
		// Register gateway handler.
		if err := RegisterInnerHashHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
	}

	return nil
}

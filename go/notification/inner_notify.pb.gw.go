// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: inner_notify.proto

/*
Package notification is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package notification

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_InnerNotify_Post_0(ctx context.Context, marshaler runtime.Marshaler, client InnerNotifyClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Notification
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Post(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InnerNotify_Post_0(ctx context.Context, marshaler runtime.Marshaler, server InnerNotifyServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Notification
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Post(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterInnerNotifyHandlerServer registers the http handlers for service InnerNotify to "mux".
// UnaryRPC     :call InnerNotifyServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterInnerNotifyHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InnerNotifyServer) error {

	mux.Handle("POST", pattern_InnerNotify_Post_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InnerNotify_Post_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InnerNotify_Post_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterInnerNotifyHandlerFromEndpoint is same as RegisterInnerNotifyHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInnerNotifyHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInnerNotifyHandler(ctx, mux, conn)
}

// RegisterInnerNotifyHandler registers the http handlers for service InnerNotify to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInnerNotifyHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInnerNotifyHandlerClient(ctx, mux, NewInnerNotifyClient(conn))
}

// RegisterInnerNotifyHandlerClient registers the http handlers for service InnerNotify
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InnerNotifyClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InnerNotifyClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InnerNotifyClient" to call the correct interceptors.
func RegisterInnerNotifyHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InnerNotifyClient) error {

	mux.Handle("POST", pattern_InnerNotify_Post_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InnerNotify_Post_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InnerNotify_Post_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_InnerNotify_Post_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"notification", "inner", "notify", "queue"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_InnerNotify_Post_0 = runtime.ForwardResponseMessage
)

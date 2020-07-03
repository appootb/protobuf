package service

import (
	"context"

	"github.com/appootb/protobuf/go/permission"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implementor interface.
type Implementor interface {
	// Return server context.
	Context() context.Context

	// Return the unary server interceptor for local gateway handler server.
	UnaryServerInterceptor() grpc.UnaryServerInterceptor

	// Return the stream server interceptor for local gateway handler server.
	StreamServerInterceptor() grpc.StreamServerInterceptor

	// Get gRPC server of the specified visible scope.
	GetScopedGRPCServer(scope permission.VisibleScope) []*grpc.Server

	// Get gateway mux of the specified visible scope.
	GetScopedGatewayMux(scope permission.VisibleScope) []*runtime.ServeMux
}

// Server authenticator interface.
type Authenticator interface {
	// Register required token level of the service.
	// The map key of the parameter is the full url path of the method.
	RegisterServiceTokenLevel(fullMethodTokenLevels map[string][]permission.Audience)

	// Authenticate a request specified by the full url path of the method.
	Authenticate(ctx context.Context, fullMethod string) (*permission.Secret, error)
}

type secretKey struct{}

// UnaryServerInterceptor returns a new unary server interceptor that authenticates incoming messages.
//
// Invalid messages will be rejected with `PermissionDenied` before reaching any userspace handlers.
func UnaryServerInterceptor(v Authenticator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		secret, err := v.Authenticate(ctx, info.FullMethod)
		if err != nil {
			if _, ok := err.(interface {
				GRPCStatus() *status.Status
			}); ok {
				return nil, err
			}
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}
		return handler(context.WithValue(ctx, secretKey{}, secret), req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that authenticates incoming messages.
//
// The stage at which unauthenticated messages will be rejected with `PermissionDenied` varies based on the
// type of the RPC. For `ServerStream` (1:m) requests, it will happen before reaching any userspace
// handlers. For `ClientStream` (n:1) or `BidiStream` (n:m) RPCs, the messages will be rejected on
// calls to `stream.Recv()`.
func StreamServerInterceptor(v Authenticator) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		secret, err := v.Authenticate(stream.Context(), info.FullMethod)
		if err != nil {
			if _, ok := err.(interface {
				GRPCStatus() *status.Status
			}); ok {
				return err
			}
			return status.Errorf(codes.PermissionDenied, err.Error())
		}
		wrapper := &ctxWrapper{stream, secret}
		return handler(srv, wrapper)
	}
}

type ctxWrapper struct {
	grpc.ServerStream
	secret *permission.Secret
}

func (s *ctxWrapper) Context() context.Context {
	ctx := s.ServerStream.Context()
	return context.WithValue(ctx, secretKey{}, s.secret)
}

func AccountSecretFromContext(ctx context.Context) *permission.Secret {
	if secret := ctx.Value(secretKey{}); secret != nil {
		return secret.(*permission.Secret)
	}
	return nil
}

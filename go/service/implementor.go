package service

import (
	"context"

	"github.com/appootb/protobuf/go/permission"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GatewayHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error

// Service implementor interface.
type Implementor interface {
	// Get gRPC server of the specified visible scope.
	GetScopedGRPCServer(scope permission.VisibleScope) []*grpc.Server

	// Register gateway handler to the specified visible scope.
	RegisterGateway(scope permission.VisibleScope, handler GatewayHandler) error
}

// The secret information of the account token.
type Authorization struct {
	Account  uint64                 // Account ID
	Device   string                 // Account device ID
	Issuer   string                 // Token issuer
	Level    permission.TokenLevel  // Token access level
	Secret   []byte                 // Token secret data
	Roles    []string               // Roles associated with the token
	Metadata map[string]interface{} // Any other associated metadata
}

// Server authenticator interface.
type Authenticator interface {
	// Register required token level of the service.
	// The map key of the parameter is the full url path of the method.
	RegisterServiceTokenLevel(fullMethodTokenLevels map[string]permission.TokenLevel)

	// Authenticate a request specified by the full url path of the method.
	Authenticate(ctx context.Context, fullMethod string) (*Authorization, error)
}

type authenticatorKey struct{}

// UnaryServerInterceptor returns a new unary server interceptor that authenticates incoming messages.
//
// Invalid messages will be rejected with `PermissionDenied` before reaching any userspace handlers.
func UnaryServerInterceptor(v Authenticator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		data, err := v.Authenticate(ctx, info.FullMethod)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, err.Error())
		}
		return handler(context.WithValue(ctx, authenticatorKey{}, data), req)
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
		data, err := v.Authenticate(stream.Context(), info.FullMethod)
		if err != nil {
			return status.Errorf(codes.PermissionDenied, err.Error())
		}
		wrapper := &ctxWrapper{stream, data}
		return handler(srv, wrapper)
	}
}

type ctxWrapper struct {
	grpc.ServerStream
	data *Authorization
}

func (s *ctxWrapper) Context() context.Context {
	ctx := s.ServerStream.Context()
	return context.WithValue(ctx, authenticatorKey{}, s.data)
}

func RequestAuthorization(ctx context.Context) *Authorization {
	if data := ctx.Value(authenticatorKey{}); data != nil {
		return data.(*Authorization)
	}
	return nil
}

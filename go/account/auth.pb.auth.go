// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: auth.proto
package account

import (
	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = permission.TokenLevel_NONE_TOKEN
var _ = service.UnaryServerInterceptor

var _levelAuth = map[string]permission.TokenLevel{
	"/appootb.account.Auth/GetCode":    permission.TokenLevel_NONE_TOKEN,
	"/appootb.account.Auth/GetRegions": permission.TokenLevel_NONE_TOKEN,
	"/appootb.account.Auth/Login":      permission.TokenLevel_NONE_TOKEN,
	"/appootb.account.Auth/OAuth":      permission.TokenLevel_NONE_TOKEN,
	"/appootb.account.Auth/Refresh":    permission.TokenLevel_LOW_TOKEN,
}

// Register scoped server.
func RegisterAuthScopeServer(auth service.Authenticator, impl service.Implementor, srv AuthServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelAuth)

	// Register scoped gRPC server.
	for _, grpc := range impl.GetScopedGRPCServer(permission.VisibleScope_DEFAULT_SCOPE) {
		RegisterAuthServer(grpc, srv)
	}
	// Register scoped gateway handler.
	return impl.RegisterGateway(permission.VisibleScope_DEFAULT_SCOPE, RegisterAuthHandler)
}
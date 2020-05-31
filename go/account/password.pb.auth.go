// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: password.proto
package account

import (
	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = permission.TokenLevel_NONE_TOKEN
var _ = service.UnaryServerInterceptor

var _levelPassword = map[string]permission.TokenLevel{
	"/appootb.account.Password/Reset":  permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Password/Set":    permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Password/Update": permission.TokenLevel_MIDDLE_TOKEN,
}

// Register scoped server.
func RegisterPasswordScopeServer(auth service.Authenticator, impl service.Implementor, srv PasswordServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelPassword)

	// Register scoped gRPC server.
	for _, grpc := range impl.GetScopedGRPCServer(permission.VisibleScope_DEFAULT_SCOPE) {
		RegisterPasswordServer(grpc, srv)
	}
	// Register scoped gateway handler.
	return impl.RegisterGateway(permission.VisibleScope_DEFAULT_SCOPE, RegisterPasswordHandler)
}

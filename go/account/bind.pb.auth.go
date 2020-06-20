// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: bind.proto
package account

import (
	"github.com/appootb/protobuf/go/permission"
	"github.com/appootb/protobuf/go/service"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = permission.TokenLevel_NONE_TOKEN
var _ = service.UnaryServerInterceptor

var _levelBind = map[string]permission.TokenLevel{
	"/appootb.account.Bind/Apply":  permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Bind/Cancel": permission.TokenLevel_MIDDLE_TOKEN,
	"/appootb.account.Bind/Gets":   permission.TokenLevel_MIDDLE_TOKEN,
}

// Register scoped server.
func RegisterBindScopeServer(auth service.Authenticator, impl service.Implementor, srv BindServer) error {
	// Register service required token level.
	auth.RegisterServiceTokenLevel(_levelBind)

	// Register scoped gRPC server.
	for _, grpc := range impl.GetScopedGRPCServer(permission.VisibleScope_DEFAULT_SCOPE) {
		RegisterBindServer(grpc, srv)
	}
	// Register scoped gateway handler server.
	for _, mux := range impl.GetScopedGatewayMux(permission.VisibleScope_DEFAULT_SCOPE) {
		err := RegisterBindHandlerServer(impl.Context(), mux, srv)
		if err != nil {
			return err
		}
	}
	return nil
}

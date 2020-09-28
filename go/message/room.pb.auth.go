// Code generated by protoc-gen-auth. DO NOT EDIT!
// source: room.proto
package message

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

var _roomServiceSubjects = map[string][]permission.Subject{
	"/appootb.message.Room/Dispatch": {
		permission.Subject_GUEST,
		permission.Subject_WEB,
		permission.Subject_PC,
		permission.Subject_MOBILE,
	},
	"/appootb.message.Room/Interact": {
		permission.Subject_GUEST,
		permission.Subject_WEB,
		permission.Subject_PC,
		permission.Subject_MOBILE,
	},
	"/appootb.message.Room/Send": {
		permission.Subject_WEB,
		permission.Subject_PC,
		permission.Subject_MOBILE,
	},
}

type wrapperRoomServer struct {
	RoomServer
	service.Implementor
}

func (w *wrapperRoomServer) Dispatch(ctx context.Context, req *common.UniqueId) (*RoomServerOption, error) {
	if w.UnaryInterceptor() == nil {
		return w.RoomServer.Dispatch(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.RoomServer,
		FullMethod: "/appootb.message.Room/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.RoomServer.Dispatch(ctx, req.(*common.UniqueId))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*RoomServerOption), nil
}

func (w *wrapperRoomServer) Interact(srv Room_InteractServer) error {
	return w.RoomServer.Interact(srv)
}

func (w *wrapperRoomServer) Send(ctx context.Context, req *Post) (*Postmark, error) {
	if w.UnaryInterceptor() == nil {
		return w.RoomServer.Send(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     w.RoomServer,
		FullMethod: "/appootb.message.Room/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return w.RoomServer.Send(ctx, req.(*Post))
	}
	resp, err := w.UnaryInterceptor()(ctx, req, info, handler)
	if err != nil {
		return nil, err
	}
	return resp.(*Postmark), nil
}

// Register scoped server.
func RegisterRoomScopeServer(component string, auth service.Authenticator, impl service.Implementor, srv RoomServer) error {
	// Register service required subjects.
	auth.RegisterServiceSubjects(component, _roomServiceSubjects)

	// Register scoped gRPC server.
	for _, gRPC := range impl.GetGRPCServer(permission.VisibleScope_CLIENT) {
		RegisterRoomServer(gRPC, srv)
	}
	// Register scoped gateway handler server.
	wrapper := wrapperRoomServer{
		RoomServer:  srv,
		Implementor: impl,
	}
	for _, mux := range impl.GetGatewayMux(permission.VisibleScope_CLIENT) {
		// Register gateway handler.
		if err := RegisterRoomHandlerServer(impl.Context(), mux, &wrapper); err != nil {
			return err
		}
		// Register websocket handler.
		if err := RegisterRoomWsHandlerServer(mux, srv, impl.StreamInterceptor()); err != nil {
			return err
		}
	}

	return nil
}

func RegisterRoomWsHandlerServer(mux *runtime.ServeMux, srv RoomServer, streamInterceptor grpc.StreamServerInterceptor) error {

	mux.Handle("GET", ws_pattern_Room_Interact_0, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		fn := func(c *websocket.Conn) {
			ctx, err := runtime.AnnotateIncomingContext(r.Context(), mux, r)
			if err != nil {
				_ = c.WriteClose(http.StatusBadRequest)
				return
			}
			inbound, outbound := runtime.MarshalerForRequest(mux, r)
			stream := webstream.NewWebsocketStream(ctx, c, inbound, outbound)
			if streamInterceptor == nil {
				err := srv.Interact(&roomInteractServer{stream})
				if err != nil {
					_ = c.WriteClose(http.StatusInternalServerError)
				}
				return
			}
			handler := func(_ interface{}, stream grpc.ServerStream) error {
				return srv.Interact(&roomInteractServer{stream})
			}
			info := &grpc.StreamServerInfo{
				FullMethod:     "/appootb.message.Room/Interact",
				IsClientStream: true,
				IsServerStream: true,
			}
			if err := streamInterceptor(srv, stream, info, handler); err != nil {
				_ = c.WriteClose(http.StatusInternalServerError)
			}
		}
		websocket.Handler(fn).ServeHTTP(w, r)
	})

	return nil
}

var (
	ws_pattern_Room_Interact_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"message", "room"}, "", runtime.AssumeColonVerbOpt(true)))
)

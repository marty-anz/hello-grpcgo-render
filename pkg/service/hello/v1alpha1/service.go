package v1alpha1

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	"github.com/marty-anz/hello-grpcgo-render/pkg/hello/service/v1alpha1"
)

type Service struct {
}

func (s *Service) SayHello(context.Context, *connect.Request[v1alpha1.SayHelloRequest]) (*connect.Response[v1alpha1.SayHelloResponse], error) {
	return connect.NewResponse(
		&v1alpha1.SayHelloResponse{Message: "hello world", Number: 1},
	), nil
}

func (s *Service) SayHelloError(ctx context.Context, req *connect.Request[v1alpha1.SayHelloErrorRequest]) (*connect.Response[v1alpha1.SayHelloResponse], error) {
	return nil, connect.NewError(connect.Code(req.Msg.GetCode()), errors.New("SayHelloError"))
}

package v1alpha

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/marty-anz/hello-grpcgo-render/pkg/hello/service/v1alpha1"
)

type Service struct {
}

func (s *Service) SayHello(context.Context, *connect_go.Request[v1alpha1.SayHelloRequest]) (*connect_go.Response[v1alpha1.SayHelloResponse], error) {
	return connect_go.NewResponse(
		&v1alpha1.SayHelloResponse{Message: "hello world"},
	), nil
}

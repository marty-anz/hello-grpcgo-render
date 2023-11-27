package v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/marty-anz/hello-grpcgo-render/pkg/hello/service/v1alpha1"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, req *connect.Request[v1alpha1.SayHelloRequest]) (*connect.Response[v1alpha1.SayHelloResponse], error) {
	// fmt.Printf("request headers: %v", req.Header())

	resp := &v1alpha1.SayHelloResponse{Number: 2, Message: "hello world", MessageId: "message id"}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("err: %w\n", err)
		return nil, err
	}

	fmt.Printf("encoding/json data: %s\n", string(data))

	data, err = protojson.Marshal(resp)
	if err != nil {
		fmt.Printf("err: %w\n", err)
		return nil, err
	}

	fmt.Printf("protojson data: %s\n", string(data))

	return connect.NewResponse(
		&v1alpha1.SayHelloResponse{Number: 2, Message: "hello world"},
	), nil
}

func (s *Service) SayHelloError(ctx context.Context,
	req *connect.Request[v1alpha1.SayHelloErrorRequest]) (*connect.Response[v1alpha1.SayHelloErrorResponse], error) {
	return nil, connect.NewError(connect.Code(2), errors.New("Say Hello Error"))
}

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/marty-anz/hello-grpcgo-render/pkg/hello/service/v1alpha1/v1alpha1connect"
	"github.com/marty-anz/hello-grpcgo-render/pkg/service/hello/v1alpha1"
)

func main() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serviceNames := []string{v1alpha1connect.HelloName}

	reflector := grpcreflect.NewStaticReflector(serviceNames...)
	checker := grpchealth.NewStaticChecker(serviceNames...)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	mux.Handle(grpchealth.NewHandler(checker))

	service := &v1alpha1.Service{}

	servicePath, serviceHandler := v1alpha1connect.NewHelloHandler(
		service,
	)

	mux.Handle(servicePath, serviceHandler)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Default().Println(fmt.Sprintf("HTTP listen and serve: %v", err))
	}
}

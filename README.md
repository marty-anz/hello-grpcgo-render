# hello-grpcgo-render

1. `cd helloapis`
1. `buf generate`
1. `cd ..`
1. `go run cmd/hello/main.go`
1. `grpcurl -plaintext localhost:8080 list`
1. `grpcurl -plaintext -d'{}' localhost:8080 localhost:8080 hello.service.v1alpha1.Hello/SayHello`

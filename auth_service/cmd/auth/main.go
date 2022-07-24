package main

import (
	"flag"
	"github.com/gxrlxv/musique/auth_service/internal/app"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

func main() {
	app.Run()
}

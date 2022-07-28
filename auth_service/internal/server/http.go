package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHTTPServer(endpoint string, log *logrus.Logger) *runtime.ServeMux {
	log.Info("new http server")
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := v1.RegisterAuthHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		log.Panic(err)
	}

	return mux
}

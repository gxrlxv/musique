package main

import (
	"context"
	"flag"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/config"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/internal/service"
	"github.com/gxrlxv/musique/auth_service/internal/usecase"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"google.golang.org/grpc"
	"net"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

func main() {
	log := logging.GetLogger()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatalf("%v", err)
	}
	authRepo := repository.NewRepository(postgreSQLClient, log)
	//authServ := service.NewAuthService(authRepo)

	//u, err := authRepo.FindByUsername(context.Background(), "gxr123")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(u.ID)
	hasher := hash.NewSHA1Hasher("salt")

	authUseCase := usecase.NewAuthUseCase(authRepo, hasher, log)

	s := grpc.NewServer()

	srv := service.NewAuthService(authUseCase, log)
	log.Info("register auth service")
	v1.RegisterAuthServer(s, srv)

	log.Info("listen :8080")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	log.Info("run server")
	err = s.Serve(lis)
	if err != nil {
		return
	}
}

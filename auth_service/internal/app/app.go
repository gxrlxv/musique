package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/config"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/internal/server"
	"github.com/gxrlxv/musique/auth_service/internal/service"
	"github.com/gxrlxv/musique/auth_service/internal/usecase"
	"github.com/gxrlxv/musique/auth_service/pkg/auth"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

func Run() {
	log := logging.GetLogger()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatalf("%v", err)
	}

	authRepo := repository.NewRepository(postgreSQLClient, log)

	hasher := hash.NewSHA1Hasher(cfg.Hasher.Salt)

	manager, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		log.Fatalf("%v", err)
	}

	authUseCase := usecase.NewAuthUseCase(authRepo, hasher, *manager, log, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL)

	authService := service.NewAuthService(authUseCase, log)
	log.Info("new grpc server")
	grpcServer := server.NewGRPCServer(authService, log)

	listen, err := net.Listen("tcp", cfg.Server.Grpc.Addr)
	if err != nil {
		panic(err)
	}

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = v1.RegisterAuthHandlerFromEndpoint(context.Background(), mux, cfg.Server.Grpc.Addr, opts)
	if err != nil {
		panic(err)
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(cfg.Server.Http.Addr, mux)
	})

	err = g.Wait()
	if err != nil {
		panic(err)
	}
}

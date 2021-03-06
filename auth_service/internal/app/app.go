package app

import (
	"context"
	"fmt"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/config"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/internal/service"
	"github.com/gxrlxv/musique/auth_service/internal/usecase"
	"github.com/gxrlxv/musique/auth_service/pkg/auth"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"google.golang.org/grpc"
	"net"
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

	s := grpc.NewServer()

	srv := service.NewAuthService(authUseCase, log)
	log.Info("register auth service")
	v1.RegisterAuthServer(s, srv)

	log.Infof("listen: %s", cfg.Listen.Port)
	lis, err := net.Listen(cfg.Listen.Type, fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		return
	}

	log.Info("run server")
	err = s.Serve(lis)
	if err != nil {
		return
	}
}

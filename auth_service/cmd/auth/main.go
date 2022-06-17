package main

import (
	"context"
	"fmt"
	"github.com/gxrlxv/musique/auth_service/internal/config"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/internal/service"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	authRepo := repository.NewRepository(postgreSQLClient, logger)
	authServ := service.NewAuthService(authRepo)

	u, err := authRepo.FindByUsername(context.Background(), "gxr123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.ID)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gxrlxv/musique/auth_service/internal/config"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	authRepo := repository.NewRepository(postgreSQLClient, logger)
	//authServ := service.NewAuthService(authRepo)

	u, err := authRepo.FindByUsername(context.Background(), "gxr123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.ID)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", "10000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.DialOption

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	//client := pb.NewRouteGuideClient(conn)

	grpcServer := grpc.NewServer(opts...)
	//pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

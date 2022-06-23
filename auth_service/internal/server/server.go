package server

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Deps struct {
	logger *logrus.Logger

	AuthHandler v1.AuthServer
}

type Server struct {
	Deps
	srv *grpc.Server
}

func NewGRPCServer(deps Deps, logger *logrus.Logger) *Server {
	log := logrus.NewEntry(logger)

	return &Server{
		srv: grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(log),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(log),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		),
		Deps: deps,
	}
}

func (s *Server) ListenAndServe(port string) error {
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	v1.RegisterAuthServer(s.srv, s.Deps.AuthHandler)

	if err := s.srv.Serve(lis); err != nil {
		return err
	}

	return nil
}

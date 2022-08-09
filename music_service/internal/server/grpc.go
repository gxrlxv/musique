package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	v1 "github.com/gxrlxv/musique/music_service/api/music/v1"
	"github.com/gxrlxv/musique/music_service/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGRPCServer(music *service.MusicService, logger *logrus.Logger) *grpc.Server {
	logger.Info("new grpc server")
	log := logrus.NewEntry(logger)
	var opts = []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(log),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterMusicServer(srv, music)

	return srv
}

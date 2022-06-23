package server

import (
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"google.golang.org/grpc"
)

type Deps struct {
	Logger logging.Logger

	serviceser
}

type Server struct {
	Deps
	srv *grpc.Server
}

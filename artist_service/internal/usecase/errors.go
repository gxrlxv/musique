package usecase

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrHaveNotPermission = status.Error(codes.Unauthenticated, "have not permission")
)

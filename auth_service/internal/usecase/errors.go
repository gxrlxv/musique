package usecase

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFoundEmail        = status.Error(codes.NotFound, "user with given email not found")
	ErrUserAlreadyExistEmail    = status.Error(codes.AlreadyExists, "user with given email already exist")
	ErrUserAlreadyExistUsername = status.Error(codes.AlreadyExists, "user with given username already exist")
	ErrUserAlreadyExistPhone    = status.Error(codes.AlreadyExists, "user with given phone already exist")
	ErrPasswordInvalid          = status.Error(codes.PermissionDenied, "invalid password")
	ErrPasswordDontMatch        = status.Error(codes.PermissionDenied, "password's don't match")
	ErrTokenInvalid             = status.Error(codes.Unauthenticated, "invalid token")
	ErrHaveNotPermission        = status.Error(codes.Unauthenticated, "have not permission")
)

func internalErr(err error) error {
	return status.Errorf(codes.Internal, "%v", err.Error())
}

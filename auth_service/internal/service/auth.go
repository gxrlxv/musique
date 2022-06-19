package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, dto domain.CreateUserDTO) error
	SignIn(ctx context.Context, dto domain.CreateUserDTO) (domain.User, error)
}

type authService struct {
	v1.UnimplementedAuthServer
	uc AuthUseCase
}

func NewAuthService(useCase AuthUseCase) *authService {
	return &authService{
		UnimplementedAuthServer: v1.UnimplementedAuthServer{},
		uc:                      useCase,
	}
}

func (a *authService) SignUp(ctx context.Context, in *v1.SignUpRequest) (*v1.SignUpReply, error) {
	userDTO := domain.CreateUserDTO{
		Username:       in.Username,
		Email:          in.Email,
		Password:       in.Password,
		RepeatPassword: in.RepeatPassword,
		FirstName:      in.FirstName,
		LastName:       in.LastName,
		Gender:         in.Gender,
		Country:        in.Country,
		City:           in.City,
		Phone:          in.Phone,
	}

	if err := a.uc.SignUp(ctx, userDTO); err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *authService) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInReply, error) {
	return nil, nil
}

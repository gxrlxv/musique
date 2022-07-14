package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, dto domain.CreateUserDTO) (string, error)
	SignIn(ctx context.Context, email, password string) (domain.User, error)
	NewTokens(ctx context.Context, userId, role string) (*v1.Tokens, error)
}

type AuthService struct {
	v1.UnimplementedAuthServer
	uc  AuthUseCase
	log *logging.Logger
}

func NewAuthService(useCase AuthUseCase, log *logging.Logger) *AuthService {
	return &AuthService{
		UnimplementedAuthServer: v1.UnimplementedAuthServer{},
		uc:                      useCase,
		log:                     log,
	}
}

func (a *AuthService) SignUp(ctx context.Context, in *v1.SignUpRequest) (*v1.SignUpReply, error) {
	a.log.Info("signUp service")
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

	userID, err := a.uc.SignUp(ctx, userDTO)
	if err != nil {
		return &v1.SignUpReply{}, nil
	}

	tokens, err := a.uc.NewTokens(ctx, userID, "user")
	if err != nil {
		return nil, err
	}

	return &v1.SignUpReply{Id: userID, Tokens: tokens}, nil
}

func (a *AuthService) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInReply, error) {
	_, err := a.uc.SignIn(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a AuthService) RefreshToken(ctx context.Context, in *v1.RefreshTokenRequest) (*v1.RefreshTokenReply, error) {
	return &v1.RefreshTokenReply{}, nil
}

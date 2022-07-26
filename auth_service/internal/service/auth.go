package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, dto domain.CreateUserDTO) (*domain.User, error)
	SignIn(ctx context.Context, email, password string) (*domain.User, error)
	NewTokens(ctx context.Context, userId, role string) (*v1.Tokens, error)
	GetIdFromRefresh(ctx context.Context, refresh string) (string, error)
	Identify(ctx context.Context, access string) (string, error)
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
	err := in.Validate()
	if err != nil {
		return &v1.SignUpReply{}, err
	}

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

	user, err := a.uc.SignUp(ctx, userDTO)
	if err != nil {
		return &v1.SignUpReply{}, err
	}

	tokens, err := a.uc.NewTokens(ctx, user.ID, user.Role)
	if err != nil {
		return &v1.SignUpReply{}, err
	}

	return &v1.SignUpReply{
		Id:     user.ID,
		Tokens: tokens,
	}, nil
}

func (a *AuthService) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInReply, error) {
	a.log.Info("signIn service")
	err := in.Validate()
	if err != nil {
		return &v1.SignInReply{}, err
	}

	user, err := a.uc.SignIn(ctx, in.Email, in.Password)
	if err != nil {
		return &v1.SignInReply{}, err
	}

	tokens, err := a.uc.NewTokens(ctx, user.ID, user.Role)
	if err != nil {
		return &v1.SignInReply{}, err
	}

	return &v1.SignInReply{
		Id:     user.ID,
		Tokens: tokens,
	}, nil
}

func (a *AuthService) RefreshToken(ctx context.Context, in *v1.RefreshTokenRequest) (*v1.RefreshTokenReply, error) {
	userId, err := a.uc.GetIdFromRefresh(ctx, in.RefreshToken)
	if err != nil {
		return &v1.RefreshTokenReply{}, err
	}

	tokens, err := a.uc.NewTokens(ctx, userId, "")
	if err != nil {
		return &v1.RefreshTokenReply{}, err
	}

	return &v1.RefreshTokenReply{Tokens: tokens}, nil
}

func (a *AuthService) Identify(ctx context.Context, in *v1.IdentityRequest) (*v1.IdentityReply, error) {
	userId, err := a.uc.Identify(ctx, in.AccessToken)
	if err != nil {
		return &v1.IdentityReply{}, err
	}
	return &v1.IdentityReply{Id: userId}, nil
}

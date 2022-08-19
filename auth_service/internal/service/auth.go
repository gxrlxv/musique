package service

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/internal/usecase"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	v1.UnimplementedAuthServer
	uc  *usecase.UseCase
	log *logrus.Logger
}

func NewAuthService(useCase *usecase.UseCase, log *logrus.Logger) *AuthService {
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

	err = a.uc.NewPlaylist(ctx, user.ID)
	if err != nil {
		return &v1.SignUpReply{}, err
	}

	err = a.uc.NewSubscription(ctx, user.ID)
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

	tokens, err := a.uc.NewTokens(ctx, userId, "user")
	if err != nil {
		return &v1.RefreshTokenReply{}, err
	}

	return &v1.RefreshTokenReply{Tokens: tokens}, nil
}

func (a *AuthService) Identify(ctx context.Context, in *v1.IdentifyRequest) (*v1.IdentifyReply, error) {
	userId, role, err := a.uc.Identify(ctx, in.AccessToken)
	if err != nil {
		return &v1.IdentifyReply{
			UserId: userId,
			Role:   role,
		}, err
	}

	return &v1.IdentifyReply{
		UserId: userId,
		Role:   role,
	}, nil
}

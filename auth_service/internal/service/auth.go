package service

import (
	"context"
	"fmt"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
)

type authService struct {
	v1.UnimplementedAuthServer
	repository AuthRepository
}

func NewAuthService(repo AuthRepository) *authService {
	return &authService{
		UnimplementedAuthServer: v1.UnimplementedAuthServer{},
		repository:              repo,
	}
}

func (a *authService) SignUp(ctx context.Context, in *v1.SignUpRequest) (*v1.SignUpReply, error) {
	if _, err := a.repository.FindByUsername(ctx, in.Username); err != nil {
		return nil, err
	}

	if _, err := a.repository.FindByEmail(ctx, in.Email); err != nil {
		return nil, err
	}

	if _, err := a.repository.FindByPhone(ctx, in.Phone); err != nil {
		return nil, err
	}

	if in.Password != in.RepeatPassword {
		return nil, fmt.Errorf("passwords don't match")
	}

	//passHash := PasswordHash(in.Password)
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

	err := a.repository.Create(ctx, domain.NewUser(userDTO))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *authService) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInReply, error) {
	return nil, nil
}

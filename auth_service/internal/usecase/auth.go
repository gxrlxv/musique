package usecase

import (
	"context"
	"fmt"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
)

type AuthRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
}

type authUseCase struct {
	repo   AuthRepository
	hasher hash.PasswordHasher
}

func NewAuthUseCase(repository AuthRepository, hasher hash.PasswordHasher) *authUseCase {
	return &authUseCase{repo: repository, hasher: hasher}
}

func (a *authUseCase) SignUp(ctx context.Context, dto domain.CreateUserDTO) error {
	if dto.Password != dto.RepeatPassword {
		return fmt.Errorf("password don't match")
	}

	if _, err := a.repo.FindByEmail(ctx, dto.Email); err != nil {
		return err
	}

	if _, err := a.repo.FindByUsername(ctx, dto.Username); err != nil {
		return err
	}

	if _, err := a.repo.FindByPhone(ctx, dto.Phone); err != nil {
		return err
	}

	passwordHash, err := a.hasher.Hash(dto.Password)
	if err != nil {
		return err
	}

	user := domain.NewUser(dto, passwordHash)

	err = a.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *authUseCase) SignIn(ctx context.Context, dto domain.CreateUserDTO) (domain.User, error) {
	return domain.User{}, nil
}

package usecase

import (
	"context"
	"fmt"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/auth"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/gxrlxv/musique/auth_service/pkg/logging"
	"time"
)

type AuthRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByPhone(ctx context.Context, phone string) (*domain.User, error)
	UpdateSession(ctx context.Context, session *domain.Session) error
	CreateSession(ctx context.Context, userID string) error
	GetIdByToken(ctx context.Context, refresh string) (string, error)
}

type authUseCase struct {
	repo            AuthRepository
	hasher          hash.PasswordHasher
	tokenManager    auth.Manager
	log             *logging.Logger
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthUseCase(repository AuthRepository, hasher hash.PasswordHasher, tokenManager auth.Manager, log *logging.Logger,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *authUseCase {
	return &authUseCase{
		repo:            repository,
		hasher:          hasher,
		tokenManager:    tokenManager,
		log:             log,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL}
}

func (a *authUseCase) SignUp(ctx context.Context, dto domain.CreateUserDTO) (*domain.User, error) {
	a.log.Info("signUp use case")
	if dto.Password != dto.RepeatPassword {
		return &domain.User{}, fmt.Errorf("password don't match")
	}

	if _, err := a.repo.GetByEmail(ctx, dto.Email); err == nil {
		return &domain.User{}, err
	}

	if _, err := a.repo.GetByUsername(ctx, dto.Username); err == nil {
		return &domain.User{}, err
	}

	if _, err := a.repo.GetByPhone(ctx, dto.Phone); err == nil {
		return &domain.User{}, err
	}
	a.log.Info("hash password")
	passwordHash, err := a.hasher.Hash(dto.Password)
	if err != nil {
		return &domain.User{}, err
	}

	model := domain.NewUser(dto, passwordHash)

	user, err := a.repo.Create(ctx, model)
	if err != nil {
		a.log.Info(err)
		return &domain.User{}, err
	}

	err = a.repo.CreateSession(ctx, user.ID)
	if err != nil {
		a.log.Info(err)
		return &domain.User{}, err
	}

	return user, nil
}

func (a *authUseCase) SignIn(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := a.repo.GetByEmail(ctx, email)
	if err != nil {
		return &domain.User{}, err
	}

	passwordHash, err := a.hasher.Hash(password)
	if err != nil {
		return &domain.User{}, err
	}

	if user.PasswordHash != passwordHash {
		return &domain.User{}, fmt.Errorf("incorrect password")
	}

	return user, nil
}

func (a *authUseCase) NewTokens(ctx context.Context, userId, role string) (*v1.Tokens, error) {
	refresh, err := a.tokenManager.NewRefreshToken()
	if err != nil {
		return &v1.Tokens{}, err
	}

	access, err := a.tokenManager.NewJWT(userId, role, a.accessTokenTTL)
	if err != nil {
		return &v1.Tokens{}, err
	}

	session := domain.Session{
		UserId:       userId,
		RefreshToken: refresh,
		ExpiresAt:    time.Now().Add(a.refreshTokenTTL),
	}

	if err := a.repo.UpdateSession(ctx, &session); err != nil {
		return &v1.Tokens{}, err
	}

	return &v1.Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

func (a *authUseCase) GetIdFromRefresh(ctx context.Context, refresh string) (string, error) {

	userId, err := a.repo.GetIdByToken(ctx, refresh)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (a *authUseCase) Identify(ctx context.Context, access string) (string, error) {
	userID, err := a.tokenManager.Parse(access)
	if err != nil {
		return "", err
	}

	return userID, nil
}

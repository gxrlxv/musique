package usecase

import (
	"context"
	v1 "github.com/gxrlxv/musique/auth_service/api/auth/v1"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/internal/repository"
	"github.com/gxrlxv/musique/auth_service/pkg/auth"
	"github.com/gxrlxv/musique/auth_service/pkg/hash"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, dto domain.CreateUserDTO) (domain.User, error)
	SignIn(ctx context.Context, email, password string) (domain.User, error)
	NewTokens(ctx context.Context, userId, role string) (*v1.Tokens, error)
	GetIdFromRefresh(ctx context.Context, refresh string) (string, error)
	Identify(ctx context.Context, access string) (string, string, error)
	NewPlaylist(ctx context.Context, userId string) error
	NewSubscription(ctx context.Context, userId string) error
}
type UseCase struct {
	AuthUseCase
}

func NewUseCase(repo *repository.Repository, log *logrus.Logger, hasher hash.PasswordHasher, tokenManager auth.Manager,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *UseCase {
	return &UseCase{
		AuthUseCase: NewAuthUseCase(repo.UserRepository, repo.SessionRepository, repo.PlaylistRepository,
			repo.SubscriptionRepository, log, hasher, tokenManager, accessTokenTTL, refreshTokenTTL),
	}
}

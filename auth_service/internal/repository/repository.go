package repository

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	GetByPhone(ctx context.Context, phone string) (domain.User, error)
}

type SessionRepository interface {
	UpdateSession(ctx context.Context, session domain.Session) error
	CreateSession(ctx context.Context, userID string) error
	GetIdByToken(ctx context.Context, refresh string) (string, error)
}

type PlaylistRepository interface {
	Create(ctx context.Context, userId string) error
}

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, subscription domain.Subscription) error
	GetBySubscriptionId(ctx context.Context, id int) (time.Duration, error)
}

type Repository struct {
	UserRepository
	SessionRepository
	PlaylistRepository
	SubscriptionRepository
}

func NewRepository(client postgresql.Client, log *logrus.Logger) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(client, log),
		SessionRepository:      NewSessionRepository(client, log),
		PlaylistRepository:     NewPlaylistRepository(client, log),
		SubscriptionRepository: NewSubscriptionRepository(client, log),
	}
}

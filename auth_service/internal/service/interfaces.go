package service

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
)

type (
	AuthRepository interface {
		Create(ctx context.Context, user *domain.User) error
		FindByUsername(ctx context.Context, username string) (domain.User, error)
		FindByEmail(ctx context.Context, email string) (domain.User, error)
		FindByPhone(ctx context.Context, phone string) (domain.User, error)
	}
)

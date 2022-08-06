package repository

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/internal/domain"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type sessionRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewSessionRepository(client postgresql.Client, log *logrus.Logger) *sessionRepository {
	return &sessionRepository{
		client: client,
		log:    log,
	}
}

func (sr *sessionRepository) CreateSession(ctx context.Context, userID string) error {
	q := `
			INSERT INTO public.session
				(user_id)
			VALUES
				($1)`

	_, err := sr.client.Exec(ctx, q, userID)
	if err != nil {
		sr.log.Error(err)
		return err
	}

	return nil
}

func (sr *sessionRepository) UpdateSession(ctx context.Context, session *domain.Session) error {
	q := `
			UPDATE public.session
			SET refresh_token = $1, expires_at = $2
			WHERE user_id = $3`

	_, err := sr.client.Exec(ctx, q, session.RefreshToken, session.ExpiresAt, session.UserId)
	if err != nil {
		sr.log.Error(err)
		return err
	}

	return nil
}

func (sr *sessionRepository) GetIdByToken(ctx context.Context, refresh string) (string, error) {
	q := `
			SELECT user_id 
			FROM public.session 
			WHERE refresh_token = $1`

	var userId string
	err := sr.client.QueryRow(ctx, q, refresh).Scan(&userId)
	if err != nil {
		sr.log.Error(err)
		return "", err
	}

	return userId, nil
}

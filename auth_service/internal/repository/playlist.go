package repository

import (
	"context"
	"github.com/gxrlxv/musique/auth_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type playlistRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewPlaylistRepository(client postgresql.Client, log *logrus.Logger) *playlistRepository {
	return &playlistRepository{
		client: client,
		log:    log,
	}
}

func (pr *playlistRepository) Create(ctx context.Context, userId string) error {
	q := `
			INSERT INTO public.playlist
				(user_id)
			VALUES
				($1)`

	_, err := pr.client.Exec(ctx, q, userId)
	if err != nil {
		pr.log.Error(err)
		return err
	}

	return nil
}

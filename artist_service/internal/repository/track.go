package repository

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/gxrlxv/musique/artist_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type trackRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewTrackRepository(client postgresql.Client, log *logrus.Logger) *trackRepository {
	return &trackRepository{
		client: client,
		log:    log,
	}
}

func (trackRepository) SaveTrack(ctx context.Context, track *domain.Track) error {

	//TODO implement me
	panic("implement me")
}

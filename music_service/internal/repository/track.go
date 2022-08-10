package repository

import (
	"github.com/gxrlxv/musique/music_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type trackRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewMusicRepository(client postgresql.Client, log *logrus.Logger) *trackRepository {
	return &trackRepository{
		client: client,
		log:    log,
	}
}

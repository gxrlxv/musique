package repository

import (
	"github.com/gxrlxv/musique/music_service/pkg/client/postgresql"
	"github.com/sirupsen/logrus"
)

type musicRepository struct {
	client postgresql.Client
	log    *logrus.Logger
}

func NewMusicRepository(client postgresql.Client, log *logrus.Logger) *musicRepository {
	return &musicRepository{
		client: client,
		log:    log,
	}
}

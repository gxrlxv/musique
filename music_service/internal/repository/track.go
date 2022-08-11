package repository

import (
	"context"
	"github.com/gxrlxv/musique/music_service/internal/domain"
	"github.com/gxrlxv/musique/music_service/pkg/client/postgresql"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrTrackNotFound = status.Error(codes.NotFound, "track with given id not found")
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

func (tr *trackRepository) GetTrack(ctx context.Context, trackId string) (domain.Track, error) {
	q := `
			SELECT title, bytes, milliseconds 
			FROM public.track 
			WHERE id = $1`

	var track domain.Track

	err := tr.client.QueryRow(ctx, q, trackId).Scan(&track.Title, &track.Bytes, &track.Milliseconds)
	if err != nil {
		tr.log.Error(err)
		if err == pgx.ErrNoRows {
			return domain.Track{}, ErrTrackNotFound
		}
		return domain.Track{}, err
	}

	return track, err
}

func (tr *trackRepository) GetTracksByAlbumId(ctx context.Context, albumId string) ([]domain.Track, error) {
	q := `
			SELECT id, title, bytes, milliseconds
			FROM public.track
			WHERE album_id = $1`

	var track domain.Track
	tracks := make([]domain.Track, 0)

	rows, err := tr.client.Query(ctx, q, albumId)
	if err != nil {
		tr.log.Error(err)
		return tracks, err
	}

	for rows.Next() {

		err := rows.Scan(&track.Id, &track.Title, &track.Bytes, &track.Milliseconds)
		if err != nil {
			return tracks, err
		}

		tracks = append(tracks, track)
	}

	return tracks, err
}

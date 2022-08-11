package repository

import (
	"context"
	"github.com/gxrlxv/musique/music_service/pkg/client/postgresql"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrPlaylistTrackNotFound = status.Error(codes.NotFound, "track with given id not found in playlist")
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

func (pr *playlistRepository) SaveTrack(ctx context.Context, playlistId, trackId string) error {
	q := `
			INSERT INTO public.playlist_tracks
				(playlist_id, track_id)
			VALUES ($1, $2)`

	_, err := pr.client.Exec(ctx, q, playlistId, trackId)
	if err != nil {
		pr.log.Error(err)
		return err
	}

	return err
}

func (pr *playlistRepository) DeleteTrack(ctx context.Context, playlistId, trackId string) error {
	q := `
			DELETE FROM public.playlist_tracks
			WHERE playlist_id = $1 and track_id = $2`

	_, err := pr.client.Exec(ctx, q, playlistId, trackId)
	if err != nil {
		pr.log.Error(err)
		return err
	}

	return err
}

func (pr *playlistRepository) GetTrack(ctx context.Context, playlistId, trackId string) error {
	q := `
			SELECT playlist_id FROM public.playlist_tracks
			WHERE playlist_id = $1 and track_id = $2`

	var exists string

	err := pr.client.QueryRow(ctx, q, playlistId, trackId).Scan(&exists)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrPlaylistTrackNotFound
		}
		return err
	}

	return err
}

func (pr *playlistRepository) GetAllTracks(ctx context.Context, playlistId string) ([]string, error) {
	q := `
			SELECT track_id FROM public.playlist_tracks
			WHERE playlist_id = $1`

	var trackId string
	tracks := make([]string, 0)

	rows, err := pr.client.Query(ctx, q, playlistId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&trackId)
		if err != nil {
			return nil, err
		}

		tracks = append(tracks, trackId)
	}
	
	return tracks, err
}

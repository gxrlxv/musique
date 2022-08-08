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

func (tr *trackRepository) SaveTrack(ctx context.Context, track *domain.Track) error {
	q := `
			INSERT INTO public.track
    			(title, artist_id, album_id, genre_id, release_year, milliseconds, bytes)
			VALUES
    			($1,$2,$3,$4,$5,$6,$7)`

	_, err := tr.client.Exec(ctx, q, track.Title, track.ArtistId, track.AlbumId, track.GenreId, track.ReleaseYear, track.Milliseconds, track.Bytes)
	if err != nil {
		tr.log.Error(err)
		return err
	}

	return nil
}

func (tr *trackRepository) DeleteTracks(ctx context.Context, albumId string) error {
	q := `
			DELETE FROM public.track
    		WHERE album_id = $1`

	_, err := tr.client.Exec(ctx, q, albumId)
	if err != nil {
		tr.log.Error(err)
		return err
	}

	return nil
}

func (tr *trackRepository) DeleteTrack(ctx context.Context, albumId, trackId string) error {
	q := `
			DELETE FROM public.track
    		WHERE album_id = $1 and id = $2`

	_, err := tr.client.Exec(ctx, q, albumId, trackId)
	if err != nil {
		tr.log.Error(err)
		return err
	}

	return nil
}

package usecase

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/sirupsen/logrus"
)

type AlbumRepository interface {
	CreateAlbum(ctx context.Context, album *domain.Album) (string, error)
	DeleteAlbum(ctx context.Context, albumId string) error
}

type TrackRepository interface {
	SaveTrack(ctx context.Context, track *domain.Track) error
	DeleteTracks(ctx context.Context, albumId string) error
}

type GenreRepository interface {
	GetByTitle(ctx context.Context, title string) (int, error)
}

type ArtistRepository interface {
	GetByName(ctx context.Context, name string) (string, error)
}

type albumUseCase struct {
	albumRepo  AlbumRepository
	trackRepo  TrackRepository
	genreRepo  GenreRepository
	artistRepo ArtistRepository
	log        *logrus.Logger
}

func NewAlbumUseCase(albumRepo AlbumRepository, trackRepo TrackRepository, genreRepo GenreRepository, artistRepo ArtistRepository, log *logrus.Logger) *albumUseCase {
	return &albumUseCase{
		albumRepo:  albumRepo,
		trackRepo:  trackRepo,
		genreRepo:  genreRepo,
		artistRepo: artistRepo,
		log:        log,
	}
}

func (a *albumUseCase) CreateAlbum(ctx context.Context, albumDTO domain.CreateAlbumDTO) error {
	artistId, err := a.artistRepo.GetByName(ctx, albumDTO.ArtistName)
	if err != nil {
		return err
	}

	album := domain.NewAlbum(albumDTO.Title, albumDTO.ReleaseYear, artistId)

	albumId, err := a.albumRepo.CreateAlbum(ctx, album)
	if err != nil {
		return err
	}

	for i := 0; i < albumDTO.NumberTracks; i++ {
		genreId, err := a.genreRepo.GetByTitle(ctx, albumDTO.Tracks[i].Genre)
		if err != nil {
			return err
		}

		track := domain.NewTrack(albumDTO.Tracks[i].Title, artistId, albumId, albumDTO.ReleaseYear, genreId, albumDTO.Tracks[i].Milliseconds, albumDTO.Tracks[i].Bytes)

		err = a.trackRepo.SaveTrack(ctx, track)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *albumUseCase) DeleteAlbum(ctx context.Context, albumId string) error {

	err := a.trackRepo.DeleteTracks(ctx, albumId)
	if err != nil {
		return err
	}

	err = a.albumRepo.DeleteAlbum(ctx, albumId)
	if err != nil {
		return err
	}

	return nil
}

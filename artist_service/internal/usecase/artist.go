package usecase

import (
	"context"
	"github.com/gxrlxv/musique/artist_service/internal/domain"
	"github.com/gxrlxv/musique/artist_service/internal/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

const artistRole = "artist"

type artistUseCase struct {
	albumRepo  repository.AlbumRepository
	trackRepo  repository.TrackRepository
	genreRepo  repository.GenreRepository
	artistRepo repository.ArtistRepository
	log        *logrus.Logger
}

func NewArtistUseCase(albumRepo repository.AlbumRepository, trackRepo repository.TrackRepository,
	genreRepo repository.GenreRepository, artistRepo repository.ArtistRepository, log *logrus.Logger) *artistUseCase {
	return &artistUseCase{
		albumRepo:  albumRepo,
		trackRepo:  trackRepo,
		genreRepo:  genreRepo,
		artistRepo: artistRepo,
		log:        log,
	}
}

func (a *artistUseCase) CreateAlbum(ctx context.Context, albumDTO domain.CreateAlbumDTO) (string, error) {
	a.log.Info("create album use case")

	artistId, err := a.artistRepo.GetByName(ctx, albumDTO.ArtistName)
	if err != nil {
		a.log.Error(err)
		return "", err
	}

	album := domain.NewAlbum(albumDTO.Title, artistId, albumDTO.ReleaseYear, albumDTO.NumberTracks)

	albumId, err := a.albumRepo.CreateAlbum(ctx, *album)
	if err != nil {
		a.log.Error(err)
		return "", err
	}

	for i := 0; i < albumDTO.NumberTracks; i++ {
		genreId, err := a.genreRepo.GetByTitle(ctx, albumDTO.Tracks[i].Genre)
		if err != nil {
			a.log.Error(err)
			return "", err
		}

		track := domain.NewTrack(albumDTO.Tracks[i].Title, artistId, albumId, albumDTO.ReleaseYear, genreId, albumDTO.Tracks[i].Milliseconds, albumDTO.Tracks[i].Bytes)

		err = a.trackRepo.SaveTrack(ctx, *track)
		if err != nil {
			a.log.Error(err)
			return "", err
		}
	}

	return albumId, err
}

func (a *artistUseCase) DeleteAlbum(ctx context.Context, albumId string) (bool, error) {
	a.log.Info("delete album use case")

	_, err := a.albumRepo.GetAlbumByID(ctx, albumId)
	if err != nil {
		return false, err
	}

	err = a.trackRepo.DeleteTracks(ctx, albumId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	err = a.albumRepo.DeleteAlbum(ctx, albumId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	return true, err
}

func (a *artistUseCase) DeleteTrack(ctx context.Context, albumId, trackId string) (bool, error) {
	a.log.Info("delete track use case")

	album, err := a.albumRepo.GetAlbumByID(ctx, albumId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	updateAlbum := domain.UpdateAlbumDTO{
		Id:           albumId,
		Title:        album.Title,
		NumberTracks: album.NumberTracks - 1,
	}

	err = a.albumRepo.UpdateAlbum(ctx, updateAlbum)
	if err != nil {
		return false, err
	}

	err = a.trackRepo.DeleteTrack(ctx, albumId, trackId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	return true, err
}

func (a *artistUseCase) AddTrack(ctx context.Context, albumId string, trackDTO domain.CreateTrackDTO) (bool, error) {
	a.log.Info("add track use case")

	album, err := a.albumRepo.GetAlbumByID(ctx, albumId)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	updateAlbum := domain.UpdateAlbumDTO{
		Id:           albumId,
		Title:        album.Title,
		NumberTracks: album.NumberTracks + 1,
	}

	err = a.albumRepo.UpdateAlbum(ctx, updateAlbum)
	if err != nil {
		return false, err
	}

	genreId, err := a.genreRepo.GetByTitle(ctx, trackDTO.Genre)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	track := domain.NewTrack(trackDTO.Title, album.ArtistId, albumId, album.ReleaseYear, genreId, trackDTO.Milliseconds, trackDTO.Bytes)

	err = a.trackRepo.SaveTrack(ctx, *track)
	if err != nil {
		a.log.Error(err)
		return false, err
	}

	return true, nil
}

func (a *artistUseCase) CheckPermission(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	role := md.Get("role")
	if role[0] != artistRole {
		return ErrHaveNotPermission
	}

	return nil
}

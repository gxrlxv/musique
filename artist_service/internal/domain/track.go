package domain

type Track struct {
	Id           string
	Title        string
	ArtistId     string
	AlbumId      string
	GenreId      int
	ReleaseYear  int32
	Milliseconds int64
	Bytes        int64
}

type CreateTrackDTO struct {
	Title        string
	Genre        string
	Milliseconds int64
	Bytes        int64
}

func NewTrack(title, artistId, albumId string, releaseYear int32, genreId int, milliseconds, bytes int64) *Track {
	return &Track{
		Id:           "",
		Title:        title,
		ArtistId:     artistId,
		AlbumId:      albumId,
		GenreId:      genreId,
		ReleaseYear:  releaseYear,
		Milliseconds: milliseconds,
		Bytes:        bytes,
	}
}

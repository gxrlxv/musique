package domain

type Album struct {
	Id           string
	Title        string
	ReleaseYear  int32
	ArtistId     string
	NumberTracks int
}

type CreateAlbumDTO struct {
	Title        string
	ReleaseYear  int32
	ArtistName   string
	NumberTracks int
	Tracks       []*CreateTrackDTO
}

type UpdateAlbumDTO struct {
	Id           string
	Title        string
	NumberTracks int
}

func NewAlbum(title, artistId string, releaseYear int32, numberTracks int) *Album {
	return &Album{
		Id:           "",
		Title:        title,
		ReleaseYear:  releaseYear,
		ArtistId:     artistId,
		NumberTracks: numberTracks,
	}
}

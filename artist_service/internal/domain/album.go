package domain

type Album struct {
	Id          string
	Title       string
	ReleaseYear string
	ArtistId    string
}

type CreateAlbumDTO struct {
	Title        string
	ReleaseYear  string
	ArtistName   string
	NumberTracks int
	Tracks       []*CreateTrackDTO
}

func NewAlbum(title, releaseYear, artistId string) *Album {
	return &Album{
		Id:          "",
		Title:       title,
		ReleaseYear: releaseYear,
		ArtistId:    artistId,
	}
}

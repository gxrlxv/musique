package domain

type Album struct {
	id          string
	title       string
	releaseYear string
	artistId    string
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
		id:          "",
		title:       title,
		releaseYear: releaseYear,
		artistId:    artistId,
	}
}

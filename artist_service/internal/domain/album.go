package domain

type Album struct {
	title       string
	releaseYear string
	artistId    string
}

type CreateAlbumDTO struct {
	Title       string
	ReleaseYear string
	ArtistName  string
}

func NewAlbum(title, releaseYear, artistId string) *Album {
	return &Album{
		title:       title,
		releaseYear: releaseYear,
		artistId:    artistId,
	}
}

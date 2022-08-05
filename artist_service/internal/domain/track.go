package domain

type Track struct {
	id           string
	title        string
	artistId     string
	albumId      string
	genreId      int
	releaseYear  string
	milliseconds int64
	bytes        int64
}

type CreateTrackDTO struct {
	Title        string
	Genre        string
	Milliseconds int64
	Bytes        int64
}

func NewTrack(title, artistId, albumId, releaseYear string, genreId int, milliseconds, bytes int64) *Track {
	return &Track{
		id:           "",
		title:        title,
		artistId:     artistId,
		albumId:      albumId,
		genreId:      genreId,
		releaseYear:  releaseYear,
		milliseconds: milliseconds,
		bytes:        bytes,
	}
}

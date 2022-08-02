package domain

type Playlist struct {
	id     int
	title  string
	userId string
}

func NewPlaylist(title, userId string) *Playlist {
	return &Playlist{
		id:     0,
		title:  title,
		userId: userId,
	}
}

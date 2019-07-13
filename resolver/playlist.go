package resolver

import (
	"fmt"
	"net/http"
)

type PlaylistResolver struct {
	playlistId string
}

func NewPlaylist(id string) (string, error) {
	url := fmt.Sprintf("https://www.youtube.com/playlist?list=%s", id)

	_, err := http.Get(url)

	if err != nil {
		return "", err
	}

	return "", nil
}

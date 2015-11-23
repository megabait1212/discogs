package discogs

import "fmt"

type ArtistReleasesService struct {
	client *Client
}

type ArtistReleases struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Items   int `json:"items"`
		Page    int `json:"page"`
		URLs    struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
		Pages int `json:"pages"`
	} `json:"pagination"`
	Releases []struct {
		Thumb       string `json:"thumb"`
		Artist      string `json:"artist"`
		MainRelease int    `json:"main_release"`
		Title       string `json:"title"`
		Role        string `json:"role"`
		Year        int    `json:"year"`
		ResourceURL string `json:"resource_url"`
		Type        string `json:"type"`
		ID          int    `json:"id"`
	} `json:"releases"`
}

// Get an artist by ID.
func (s *ArtistReleasesService) Get(id int) (*ArtistReleases, *Response, error) {
	u := fmt.Sprintf("artists/%d/releases", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	artistReleases := new(ArtistReleases)
	resp, err := s.client.Do(req, artistReleases)
	if err != nil {
		return nil, resp, err
	}

	return artistReleases, resp, err
}

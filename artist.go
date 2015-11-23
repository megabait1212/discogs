package discogs

import "fmt"

type ArtistService struct {
	client *Client
}

type Artist struct {
	Name            string   `json:"name"`
	Name_Variations []string `json:"namevariations"`
	Real_Name       string   `json:"realname"`
	Aliases         []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Resource_URL string `json:"resource_url"`
	} `json:"aliases"`
	Profile      string   `json:"profile"`
	Releases_URL string   `json:"release_url"`
	Resource_URL string   `json:"resource_url"`
	URI          string   `json:"uri"`
	URLs         []string `json:"urls"`
	Data_quality string   `json:"data_quality"`
	ID           int      `json:"id"`
	Images       []struct {
		Height       int    `json:"height"`
		Resource_URL string `json:"resource_url"`
		Type         string `json:"type"`
		URI          string `json:"uri"`
		URI_150      string `json:"uri150"`
		Width        int    `json:"width"`
	} `json:"images"`
	Members []struct {
		Active       bool   `json:"active"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Resource_URL string `json:"resource_url"`
	} `json:"members"`
}

// Get an artist by ID.
func (s *ArtistService) Get(id int) (*Artist, *Response, error) {
	u := fmt.Sprintf("artists/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	artist := new(Artist)
	resp, err := s.client.Do(req, artist)
	if err != nil {
		return nil, resp, err
	}

	return artist, resp, err
}

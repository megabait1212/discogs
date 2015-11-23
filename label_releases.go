package discogs

import "fmt"

type LabelReleasesService struct {
	client *Client
}

type LabelReleases struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Pages   int `json:"pages"`
		Page    int `json:"page"`
		URLs    struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
		Items int `json:"items"`
	} `json:"pagination"`
	Releases []struct {
		Status      string `json:"status"`
		Thumb       string `json:"thumb"`
		Title       string `json:"title"`
		Format      string `json:"format"`
		Catno       string `json:"catno"`
		ResourceURL string `json:"resource_url"`
		Artist      string `json:"artist"`
		ID          int    `json:"id"`
	} `json:"releases"`
}

// Get a label's releases by ID.
func (s *LabelReleasesService) Get(id int) (*LabelReleases, *Response, error) {
	u := fmt.Sprintf("labels/%d/releases", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	labelReleases := new(LabelReleases)
	resp, err := s.client.Do(req, labelReleases)
	if err != nil {
		return nil, resp, err
	}

	return labelReleases, resp, err
}

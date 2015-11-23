package discogs

import "fmt"

type MasterVersionsService struct {
	client *Client
}

type MasterVersions struct {
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
	Versions []struct {
		Status      string `json:"status"`
		Thumb       string `json:"thumb"`
		Title       string `json:"title"`
		Country     string `json:"country"`
		Format      string `json:"format"`
		Label       string `json:"label"`
		Released    string `json:"released"`
		CatNo       string `json:"catno"`
		ResourceUrl string `json:"resource_url"`
		ID          int    `json:"id"`
	} `json:"versions"`
}

// Get a version of a master by ID.
func (s *MasterVersionsService) Get(id int) (*MasterVersions, *Response, error) {
	u := fmt.Sprintf("masters/%d/versions", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	masterVersions := new(MasterVersions)
	resp, err := s.client.Do(req, masterVersions)
	if err != nil {
		return nil, resp, err
	}

	return masterVersions, resp, err
}

package discogs

import "fmt"

type LabelService struct {
	client *Client
}

type Label struct {
	Profile     string `json:"profile"`
	ReleasesURL string `json:"releases_url"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
	URI         string `json:"uri"`
	Sublabels   []struct {
		ResourceURL string `json:"resource_url"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
	} `json:"sublabels"`
	URLs   []string `json:"urls"`
	Images []struct {
		URI         string `json:"uri"`
		Height      int    `json:"height"`
		Width       int    `json:"width"`
		ResourceURL string `json:"resource_url"`
		Type        string `json:"type"`
		URI150      string `json:"uri150"`
	} `json:"images"`
	ResourceURL string `json:"resource_url"`
	Id          int    `json:"id"`
	DataQuality string `json:"data_quality"`
}

// Get a label by ID.
func (s *LabelService) Get(id int) (*Label, *Response, error) {
	u := fmt.Sprintf("labels/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	label := new(Label)
	resp, err := s.client.Do(req, label)
	if err != nil {
		return nil, resp, err
	}

	return label, resp, err
}

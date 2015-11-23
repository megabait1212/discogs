package discogs

import "fmt"

type MasterService struct {
	client *Client
}

type Master struct {
	Styles []string `json:"styles"`
	Genres []string `json:"genres"`
	Videos []struct {
		Duration    int    `json:"duration"`
		Description string `json:"description"`
		Embed       bool   `json:"embed"`
		URI         string `json:"uri"`
		Title       string `json:"title"`
	} `json:"videos"`
	Title          string `json:"title"`
	MainRelease    int    `json:"main_release"`
	MainReleaseURL string `json:"main_release_url"`
	URI            string `json:"uri"`
	Artists        []struct {
		Join        string `json:"join"`
		Name        string `json:"name"`
		ANV         string `json:"anv"`
		Tracks      string `json:"tracks"`
		Role        string `json:"role"`
		ResourceURL string `json:"resource_url"`
		ID          int    `json:"id"`
	} `json:"artists"`
	VersionsURL string `json:"versions_url"`
	Year        int    `json:"year"`
	Images      []struct {
		URI         string `json:"uri"`
		Height      int    `json:"height"`
		Width       int    `json:"width"`
		ResourceURL string `json:"resource_url"`
		Type        string `json:"type"`
		URI150      string `json:"uri150"`
	} `json:"images"`
	ResourceURL string `json:"resource_url"`
	Tracklist   []struct {
		Duration string `json:"duration"`
		Position string `json:"position"`
		Type     string `json:"type_"`
		Title    string `json:"title"`
	} `json:"tracklist"`
	ID          int    `json:"id"`
	DataQuality string `json:"data_quality"`
}

// Get a master by ID.
func (s *MasterService) Get(id int) (*Master, *Response, error) {
	u := fmt.Sprintf("masters/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	master := new(Master)
	resp, err := s.client.Do(req, master)
	if err != nil {
		return nil, resp, err
	}

	return master, resp, err
}

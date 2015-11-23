package discogs

import "fmt"

type ReleaseService struct {
	client *Client
}

type Release struct {
	Title   string `json:"title"`
	ID      int    `json:"id"`
	Artists []struct {
		ANV          string `json:"anv"` // Artist Name Variation
		ID           int    `json:"id"`
		Join         string `json:"join"`
		Name         string `json:"name"`
		Resource_URL string `json:"resource_url"`
		Role         string `json:"role"`
		Tracks       string `json:"tracks"` // ie "1, 3" or "1 to 9"
	} `json:"artists"`
	Data_Quality string `json:"data_quality"`
	Thumb        string `json:"thumb"` // Thumbnail
	Community    struct {
		Contributors []struct {
			Resource_URL string `json:"resource_url"`
			Username     string `json:"username"`
		} `json:"contributors"`
		Data_Quality string `json:"data_quality"`
		Have         int    `json:"have"`
		Rating       struct {
			Average float64 `json:"average"`
			Count   int     `json:"count"`
		} `json:"rating"`
		Status    string `json:"status"`
		submitter struct {
			Resource_URL string `json:"resource_url"`
			Username     string `json:"username"`
		}
		Want int `json:"want"`
	} `json:"community"`
	Companies []struct {
		CatNo            string `json:"catno"`
		Entity_Type      string `json:"entity_type"`
		Entity_Type_Name string `json:"entity_type_name"`
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Resource_URL     string `json:"resource_url"`
	} `json:"companies"`
	Country          string `json:"country"`
	Date_Added       string `json:"date_added"`
	Date_Changed     string `json:"date_changed"`
	Estimated_Weight int    `json:"estimated_weight"`
	ExtraArtists     []struct {
		ANV          string `json:"anv"`
		ID           int    `json:"id"`
		Join         string `json:"join"`
		Name         string `json:"name"`
		Resource_URL string `json:"resource_url"`
		Role         string `json:"role"`
		Tracks       string `json:"tracks"`
	}
	Format_Quantity int `json:"format_quantity"`
	Formats         []struct {
		Descriptions []string `json:"descriptions"`
		Name         string   `json:"name"`
		Qty          string   `json:"qty"`
	} `json:"formats"`
	Genres      []string `json:"genres"`
	Identifiers []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"identifiers"`
	Images []struct {
		Height       int    `json:"height"`
		Resource_URL string `json:"resource_url"`
		Type         string `json:"type"`
		URI          string `json:"uri"`
		URI150       string `json:"uri_150"`
		Width        int    `json:"width"`
	} `json:"images"`
	Labels []struct {
		CatNo        string `json:"catno"`
		Entity_Type  string `json:"entity_type"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Resource_URL string `json:"resource_url"`
	} `json:"labels"`
	Master_ID          int    `json:"master_id"`
	Master_URL         string `json:"master_url"`
	Notes              string `json:"notes"`
	Released           string `json:"released"`
	Released_Formatted string `json:"released_formatted"`
	Resource_URL       string `json:"resource_url"`
	Series             []struct {
		ID           int    `json:"id"`
		Resource_URL string `json:"resource_url"`
		CatNo        string `json:"catno"`
		Name         string `json:"name"`
		Entity_Type  string `json:"entity_type"`
	}
	Status    string   `json:"status"`
	Styles    []string `json:"styles"`
	Tracklist []struct {
		Duration string `json:"duration"`
		Position string `json:"position"`
		Title    string `json:"title"`
		Type_    string `json:"Type_"`
	} `json:"tracklist"`
	URI    string `json:"URI"`
	Videos []struct {
		Description string `json:"description"`
		Duration    int    `json:"duration"`
		Embed       bool   `json:"embed"`
		Title       string `json:"title"`
		URI         string `json:"uri"`
	} `json:"videos"`
	Year int `json:"year"`
}

// Get a release by ID.
func (s *ReleaseService) Get(id int) (*Release, *Response, error) {
	u := fmt.Sprintf("releases/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	release := new(Release)
	resp, err := s.client.Do(req, release)
	if err != nil {
		return nil, resp, err
	}

	return release, resp, err
}

package discogs

import "fmt"

type SearchService struct {
	client *Client
}
type Search struct {
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
	Results []struct {
		Style     []string `json:"style"`
		Thumb     string   `json:"thumb"`
		Title     string   `json:"title"`
		Country   string   `json:"country"`
		Format    []string `json:"format"`
		URI       string   `json:"uri"`
		Community struct {
			Want int `json:"want"`
			Have int `json:"have"`
		} `json:"community"`
		Label       []string `json:"label"`
		CatNo       string   `json:"catno"`
		Year        string   `json:"year"`
		Genre       []string `json:"genre"`
		ResourceURL string   `json:"resource_url"`
		Type        string   `json:"type"`
		ID          int      `json:"id"`
	} `json:"results"`
}

type SearchParams struct {
	Query         string
	Type          string
	Title         string
	Release_Title string
	Credit        string
	Artist        string
	ANV           string
	Label         string
	Genre         string
	Style         string
	Country       string
	Year          string
	Format        string
	CatNo         string
	Barcode       string
	Track         string
	Submitter     string
	Contributor   string
}

// Make a search.
func (s *SearchService) Get(params *SearchParams) (*Search, *Response, error) {
	id, release := "", ""
	u := fmt.Sprintf("/database/search?%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	search := new(Search)
	resp, err := s.client.Do(req, release)
	if err != nil {
		return nil, resp, err
	}

	return search, resp, err
}

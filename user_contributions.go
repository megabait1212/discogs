package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserContributionsService struct {
	client *Client
}

type UserContributions struct {
	Contributions []struct {
		Artists []struct {
			Anv         string `json:"anv"`
			Id          int    `json:"id"`
			Join        string `json:"join"`
			Name        string `json:"name"`
			ResourceUrl string `json:"resource_url"`
			Role        string `json:"role"`
			Tracks      string `json:"tracks"`
		} `json:"artists"`
		Community struct {
			Contributors []struct {
				ResourceUrl string `json:"resource_url"`
				Username    string `json:"username"`
			} `json:"contributors"`
			DataQuality string `json:"data_quality"`
			Have        int    `json:"have"`
			Rating      struct {
				Average int `json:"average"`
				Count   int `json:"count"`
			} `json:"rating"`
			Status    string `json:"status"`
			Submitter struct {
				ResourceUrl string `json:"resource_url"`
				Username    string `json:"username"`
			} `json:"submitter"`
			Want int `json:"want"`
		} `json:"community"`
		Companies       []interface{} `json:"companies"`
		Country         string        `json:"country"`
		DataQuality     string        `json:"data_quality"`
		DateAdded       string        `json:"date_added"`
		DateChanged     string        `json:"date_changed"`
		EstimatedWeight int           `json:"estimated_weight"`
		FormatQuantity  int           `json:"format_quantity"`
		Formats         []struct {
			Descriptions []string `json:"descriptions"`
			Name         string   `json:"name"`
			Qty          string   `json:"qty"`
		} `json:"formats"`
		Genres []string `json:"genres"`
		Id     int      `json:"id"`
		Images []struct {
			Height      int    `json:"height"`
			ResourceUrl string `json:"resource_url"`
			Type        string `json:"type"`
			Uri         string `json:"uri"`
			Uri150      string `json:"uri150"`
			Width       int    `json:"width"`
		} `json:"images"`
		Labels []struct {
			Catno       string `json:"catno"`
			EntityType  string `json:"entity_type"`
			Id          int    `json:"id"`
			Name        string `json:"name"`
			ResourceUrl string `json:"resource_url"`
		} `json:"labels"`
		MasterId          int           `json:"master_id"`
		MasterUrl         string        `json:"master_url"`
		Notes             string        `json:"notes"`
		Released          string        `json:"released"`
		ReleasedFormatted string        `json:"released_formatted"`
		ResourceUrl       string        `json:"resource_url"`
		Series            []interface{} `json:"series"`
		Status            string        `json:"status"`
		Styles            []string      `json:"styles"`
		Thumb             string        `json:"thumb"`
		Title             string        `json:"title"`
		Uri               string        `json:"uri"`
		Videos            []struct {
			Description string `json:"description"`
			Duration    int    `json:"duration"`
			Embed       bool   `json:"embed"`
			Title       string `json:"title"`
			Uri         string `json:"uri"`
		} `json:"videos"`
		Year int `json:"year"`
	} `json:"contributions"`
	Pagination struct {
		Items   int `json:"items"`
		Page    int `json:"page"`
		Pages   int `json:"pages"`
		PerPage int `json:"per_page"`
		Urls    struct {
		} `json:"urls"`
	} `json:"pagination"`
}

// Get a user's contributions.
func (s *UserContributionsService) Get(username string) (*UserContributions, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/contributions", username)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserContributions
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

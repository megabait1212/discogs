package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserWantsService struct {
	client *Client
}

type UserWants struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Pages   int `json:"pages"`
		Page    int `json:"page"`
		Items   int `json:"items"`
		Urls    struct {
		} `json:"urls"`
	} `json:"pagination"`
	Wants []struct {
		Rating           int `json:"rating"`
		BasicInformation struct {
			Formats []struct {
				Text         string   `json:"text"`
				Qty          string   `json:"qty"`
				Descriptions []string `json:"descriptions"`
				Name         string   `json:"name"`
			} `json:"formats"`
			Thumb  string `json:"thumb"`
			Title  string `json:"title"`
			Labels []struct {
				ResourceUrl string `json:"resource_url"`
				EntityType  string `json:"entity_type"`
				Catno       string `json:"catno"`
				Id          int    `json:"id"`
				Name        string `json:"name"`
			} `json:"labels"`
			Year    int `json:"year"`
			Artists []struct {
				Join        string `json:"join"`
				Name        string `json:"name"`
				Anv         string `json:"anv"`
				Tracks      string `json:"tracks"`
				Role        string `json:"role"`
				ResourceUrl string `json:"resource_url"`
				Id          int    `json:"id"`
			} `json:"artists"`
			ResourceUrl string `json:"resource_url"`
			Id          int    `json:"id"`
		} `json:"basic_information"`
		ResourceUrl string `json:"resource_url"`
		Id          int    `json:"id"`
	} `json:"wants"`
}

// Get a user's inventory.
func (s *UserWantsService) Get(username string) (*UserWants, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/wants", username)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserWants
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

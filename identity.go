package discogs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type IdentityService struct {
	client *Client
}

type Identity struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	ResourceUrl  string `json:"resource_url"`
	ConsumerName string `json:"consumer_name"`
}

// Get a label's releases by ID.
func (s *IdentityService) Get(id int) (*Identity, *Response, error) {
	GetAccessToken()

	response, err := oauthConsumer.Get("http://api.discogs.com/oauth/identity", nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *Identity
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

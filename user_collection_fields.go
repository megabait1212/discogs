package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserCollectionFieldsService struct {
	client *Client
}

type UserCollectionFields struct {
	Fields []struct {
		Name     string   `json:"name"`
		Options  []string `json:"options"`
		Id       int      `json:"id"`
		Position int      `json:"position"`
		Type     string   `json:"type"`
		Public   bool     `json:"public"`
	} `json:"fields"`
}

// Get a user's inventory.
func (s *UserCollectionFieldsService) Get(username string) (*UserCollectionFields, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/collection/fields", username)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserCollectionFields
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

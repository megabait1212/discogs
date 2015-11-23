package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserCollectionFoldersService struct {
	client *Client
}

type UserCollectionFolders struct {
	Folders []struct {
		Id          int    `json:"id"`
		Count       int    `json:"count"`
		Name        string `json:"name"`
		ResourceUrl string `json:"resource_url"`
	} `json:"folders"`
}

// Get a user's inventory.
func (s *UserCollectionFoldersService) Get(username string) (*UserCollectionFolders, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/collection/folders", username)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserCollectionFolders
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

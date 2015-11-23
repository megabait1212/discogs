package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserCollectionFolderService struct {
	client *Client
}

type UserCollectionFolder struct {
	Id          int    `json:"id"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	ResourceUrl string `json:"resource_url"`
}

// Get a user's inventory.
func (s *UserCollectionFolderService) Get(username string, folderID int) (*UserCollectionFolder, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/collection/folders/%d", username, folderID)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserCollectionFolder
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

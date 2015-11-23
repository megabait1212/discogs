package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserCollectionFolderReleasesService struct {
	client *Client
}

type UserCollectionFolderReleases struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Items   int `json:"items"`
		Page    int `json:"page"`
		Urls    struct {
		} `json:"urls"`
		Pages int `json:"pages"`
	} `json:"pagination"`
	Releases []struct {
		InstanceId       int `json:"instance_id"`
		Rating           int `json:"rating"`
		BasicInformation struct {
			Labels []struct {
				Id          int    `json:"id"`
				ResourceUrl string `json:"resource_url"`
				Catno       string `json:"catno"`
				Name        string `json:"name"`
				EntityType  string `json:"entity_type"`
			} `json:"labels"`
			Formats []struct {
				Descriptions []string `json:"descriptions"`
				Name         string   `json:"name"`
				Qty          string   `json:"qty"`
			} `json:"formats"`
			Thumb   string `json:"thumb"`
			Title   string `json:"title"`
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
			Year        int    `json:"year"`
			Id          int    `json:"id"`
		} `json:"basic_information"`
		Id int `json:"id"`
	} `json:"releases"`
}

// Get a user's inventory.
func (s *UserCollectionFolderReleasesService) Get(username string, folderID int) (*UserCollectionFolderReleases, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/collection/folders/%d/releases", username, folderID)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserCollectionFolderReleases
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

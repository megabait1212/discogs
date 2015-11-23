package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserService struct {
	client *Client
}

type User struct {
	Profile              string  `json:"profile"`
	WantlistUrl          string  `json:"wantlist_url"`
	Rank                 float64 `json:"rank"`
	NumPending           int     `json:"num_pending"`
	Id                   int     `json:"id"`
	NumForSale           int     `json:"num_for_sale"`
	HomePage             string  `json:"home_page"`
	Location             string  `json:"location"`
	CollectionFoldersUrl string  `json:"collection_folders_url"`
	Username             string  `json:"username"`
	CollectionFieldsUrl  string  `json:"collection_fields_url"`
	ReleasesContributed  int     `json:"releases_contributed"`
	Registered           string  `json:"registered"`
	RatingAvg            float64 `json:"rating_avg"`
	NumCollection        int     `json:"num_collection"`
	ReleasesRated        int     `json:"releases_rated"`
	NumLists             int     `json:"num_lists"`
	Name                 string  `json:"name"`
	NumWantlist          int     `json:"num_wantlist"`
	InventoryUrl         string  `json:"inventory_url"`
	AvatarUrl            string  `json:"avatar_url"`
	Uri                  string  `json:"uri"`
	ResourceUrl          string  `json:"resource_url"`
	Email                string  `json:"email"`
}

type UserPost struct {
	Username string `json:"username"`
	Name     string `json:"name,omitempty"`
	Homepage string `json:"home_page,omitempty"`
	Location string `json:"location,omitempty"`
	Profile  string `json:"profile,omitempty"`
}

// Get a user's information.
func (s *UserService) Get(username string) (*User, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s", username)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *User
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

// Post a user's information.
func (s *UserService) Post(username string, data UserPost) (*User, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s", username)

	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(body)
	response, err := oauthConsumer.Post(u, nil, accessToken, bodyString)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *User
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type UserInventoryService struct {
	client *Client
}

type UserInventory struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Items   int `json:"items"`
		Page    int `json:"page"`
		Urls    struct {
		} `json:"urls"`
		Pages int `json:"pages"`
	} `json:"pagination"`
	Listings []struct {
		Status string `json:"status"`
		Price  struct {
			Currency string  `json:"currency"`
			Value    float32 `json:"value"`
		} `json:"price"`
		AllowOffers     bool   `json:"allow_offers"`
		SleeveCondition string `json:"sleeve_condition"`
		Id              int    `json:"id"`
		Condition       string `json:"condition"`
		Posted          string `json:"posted"`
		ShipsFrom       string `json:"ships_from"`
		Uri             string `json:"uri"`
		Comments        string `json:"comments"`
		Seller          struct {
			Username    string `json:"username"`
			ResourceUrl string `json:"resource_url"`
			Id          int    `json:"id"`
		} `json:"seller"`
		Release struct {
			CatalogNumber string `json:"catalog_number"`
			ResourceUrl   string `json:"resource_url"`
			Year          int    `json:"year"`
			Id            int    `json:"id"`
			Description   string `json:"description"`
		} `json:"release"`
		ResourceUrl string `json:"resource_url"`
		Audio       bool   `json:"audio"`
	} `json:"listings"`
}

// Get a user's inventory.
func (s *UserInventoryService) Get(username string) (*UserInventory, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/users/%s/inventory", username)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *UserInventory
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MarketplaceFeeService struct {
	client *Client
}

type MarketplaceFee struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

// Get a user's inventory.
func (s *MarketplaceFeeService) Get(price float64, currency string) (*MarketplaceFee, *Response, error) {
	GetAccessToken()

	seperator := ""
	if currency != "" {
		seperator = "/"
	}

	u := fmt.Sprintf("http://api.discogs.com/marketplace/fee/%f%s%s", price, seperator, currency)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *MarketplaceFee
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

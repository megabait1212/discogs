package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MarketplacePriceSuggestionsService struct {
	client *Client
}

type MarketplacePriceSuggestions struct {
	VeryGood struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Very Good (VG)"`
	GoodPlus struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Good Plus (G+)"`
	NearMint struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Near Mint (NM or M-)"`
	Good struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Good (G)"`
	VeryGoodPlus struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Very Good Plus (VG+)"`
	Mint struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Mint (M)"`
	Fair struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Fair (F)"`
	Poor struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"Poor (P)"`
}

// Get a user's inventory.
func (s *MarketplacePriceSuggestionsService) Get(releaseID int) (*MarketplacePriceSuggestions, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/price_suggestions/%d", releaseID)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *MarketplacePriceSuggestions
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type condition struct{}

func (c condition) Mint() string {
	return "Mint (M)"
}
func (c condition) NearMint() string {
	return "Near Mint (NM or M-)"
}
func (c condition) VeryGoodPlus() string {
	return "Very Good Plus (VG+)"
}
func (c condition) VeryGood() string {
	return "Very Good (VG)"
}
func (c condition) GoodPlus() string {
	return "Good Plus (G+)"
}
func (c condition) Good() string {
	return "Good (G)"
}
func (c condition) Fair() string {
	return "Fair (F)"
}
func (c condition) Poor() string {
	return "Poor (P)"
}

var Condition = condition{}

type sleeveCondition struct {
}

func (c sleeveCondition) Mint() string {
	return "Mint (M)"
}
func (c sleeveCondition) NearMint() string {
	return "Near Mint (NM or M-)"
}
func (c sleeveCondition) VeryGoodPlus() string {
	return "Very Good Plus (VG+)"
}
func (c sleeveCondition) VeryGood() string {
	return "Very Good (VG)"
}
func (c sleeveCondition) GoodPlus() string {
	return "Good Plus (G+)"
}
func (c sleeveCondition) Good() string {
	return "Good (G)"
}
func (c sleeveCondition) Fair() string {
	return "Fair (F)"
}
func (c sleeveCondition) Poor() string {
	return "Poor (P)"
}
func (c sleeveCondition) Generic() string {
	return "Generic"
}
func (c sleeveCondition) NotGraded() string {
	return "Not Graded"
}
func (c sleeveCondition) NoCover() string {
	return "No Cover"
}

var SleeveCondition = sleeveCondition{}

type MarketplaceListingService struct {
	client *Client
}

type MarketplaceListing struct {
	Status string `json:"status"`
	Price  struct {
		Currency string  `json:"currency"`
		Value    float64 `json:"value"`
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
}

type MarketplaceListingEdit struct {
	ReleaseID       int     `json:"release_id"`
	Condition       string  `json:"condition"`                  // Must be from a set of strings, same for Sleeve Condition.
	SleeveCondition string  `json:"sleeve_condition,omitempty"` // http://www.discogs.com/developers/#page:marketplace,header:marketplace-listing-post
	Price           float64 `json:"price"`
	Comments        string  `json:"comments,omitempty"`
	AllowOffers     bool    `json:"allow_offers,omitempty"`
	Status          string  `json:"status"`
	ExternalID      int     `json:"external_id,omitempty"`
	Location        int     `json:"location,omitempty"`
	Weight          int     `json:"weight,omitempty"`
	FormatQuantity  int     `json:"format_quantity,omitempty"`
}

type MarketplaceListingNew struct {
	ListingID   int    `json:"listing_id"`
	ResourceURL string `json:"resource_url"`
}

type MarketplaceListingDelete struct {
	ListingID int `json:"listing_id"`
}

// Get a user's inventory.
func (s *MarketplaceListingService) Get(listingID int) (*MarketplaceListing, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/listings/%d", listingID)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *MarketplaceListing
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

// Edit data for a user's listing.
func (s *MarketplaceListingService) Edit(listingID int, data MarketplaceListingEdit) (*Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/listings/%d", listingID)

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

	return nil, err
}

// Delete a user's listing.
func (s *MarketplaceListingService) Delete(listingID int, data MarketplaceListingDelete) (*Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/listings/%d", listingID)

	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(body)
	response, err := oauthConsumer.Delete(u, nil, accessToken, bodyString)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return nil, err
}

// New data for a user's listing.
func (s *MarketplaceListingService) New(data MarketplaceListingEdit) (*MarketplaceListingNew, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/listings")

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

	var f *MarketplaceListingNew
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

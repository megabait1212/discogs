package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MarketplaceOrderService struct {
	client *Client
}

type MarketplaceOrder struct {
	Id          string   `json:"id"`
	ResourceUrl string   `json:"resource_url"`
	MessagesUrl string   `json:"messages_url"`
	Uri         string   `json:"uri"`
	Status      string   `json:"status"`
	NextStatus  []string `json:"next_status"`
	Fee         struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	} `json:"fee"`
	Created string `json:"created"`
	Items   []struct {
		Release struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
		} `json:"release"`
		Price struct {
			Currency string `json:"currency"`
			Value    int    `json:"value"`
		} `json:"price"`
		Id int `json:"id"`
	} `json:"items"`
	Shipping struct {
		Currency string `json:"currency"`
		Value    int    `json:"value"`
	} `json:"shipping"`
	ShippingAddress        string `json:"shipping_address"`
	AdditionalInstructions string `json:"additional_instructions"`
	Seller                 struct {
		ResourceUrl string `json:"resource_url"`
		Username    string `json:"username"`
		Id          int    `json:"id"`
	} `json:"seller"`
	LastActivity string `json:"last_activity"`
	Buyer        struct {
		ResourceUrl string `json:"resource_url"`
		Username    string `json:"username"`
		Id          int    `json:"id"`
	} `json:"buyer"`
	Total struct {
		Currency string `json:"currency"`
		Value    int    `json:"value"`
	} `json:"total"`
}

type MarketplaceOrderEdit struct {
	Status   string  `json:"status"`
	Shipping float64 `json:"shipping"`
}

// Get a user's inventory.
func (s *MarketplaceOrderService) Get(orderID int) (*MarketplaceOrder, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/orders/%d", orderID)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *MarketplaceOrder
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

// Edit data for a user's order.
func (s *MarketplaceOrderService) Edit(orderID int, data MarketplaceOrderEdit) (*Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/orders/%d", orderID)

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

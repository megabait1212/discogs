package discogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type MarketplaceOrderMessagesService struct {
	client *Client
}

type MarketplaceOrderMessages struct {
	Pagination struct {
		PerPage int `json:"per_page"`
		Pages   int `json:"pages"`
		Page    int `json:"page"`
		Urls    struct {
		} `json:"urls"`
		Items int `json:"items"`
	} `json:"pagination"`
	Messages []struct {
		From struct {
			Username    string `json:"username"`
			ResourceUrl string `json:"resource_url"`
		} `json:"from"`
		Message string `json:"message"`
		Order   struct {
			ResourceUrl string `json:"resource_url"`
			Id          string `json:"id"`
		} `json:"order"`
		Timestamp string `json:"timestamp"`
		Subject   string `json:"subject"`
	} `json:"messages"`
}

type MarketplaceOrderMessagesEdit struct {
	Status    string `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	ReleaseID int    `json:"release_id"`
}

// Get a user's inventory.
func (s *MarketplaceOrderMessagesService) Get(orderID int) (*MarketplaceOrderMessages, *Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/orders/%d/messages", orderID)
	response, err := oauthConsumer.Get(u, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *MarketplaceOrderMessages
	bits, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(bits, &f)

	return f, nil, err
}

// Add a message to the order's message log.
func (s *MarketplaceOrderMessagesService) Edit(orderID int, data MarketplaceOrderMessagesEdit) (*Response, error) {
	GetAccessToken()

	u := fmt.Sprintf("http://api.discogs.com/marketplace/orders/%d/messages", orderID)

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

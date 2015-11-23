package discogs

import (
	"io/ioutil"
	"log"
)

type ImageService struct {
	client *Client
}

type Image struct {
	ImageURL   string
	ImageBytes []byte
}

// Get an image.
func (s *ImageService) Get(imageURL string) (*Image, *Response, error) {
	GetAccessToken()

	response, err := oauthConsumer.Get(imageURL, nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var f *Image
	bits, _ := ioutil.ReadAll(response.Body)

	file := "./image.jpeg"
	err = ioutil.WriteFile(file, bits, 0644)
	if err != nil {
		panic(err)
	}

	return f, nil, err
}

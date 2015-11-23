package main

import (
	"encoding/json"
	"fmt"
)

type UserPost struct {
	Username string `json:"username"`
	Name     string `json:"name,omitempty"`
	Homepage string `json:"home_page,omitempty"`
	Location string `json:"location,omitempty"`
	Profile  string `json:"profile,omitempty"`
}

func main() {
	userPost := UserPost{
		Username: "teratomata",
		Name:     "David",
	}
	bolB, _ := json.Marshal(userPost)
	fmt.Println(string(bolB))
}

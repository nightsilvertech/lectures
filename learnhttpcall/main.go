package main

import (
	"encoding/json"
	"fmt"
	"github.com/nightsilvertech/lectures/learnhttpcall/jsonModel"
	"io/ioutil"
	"net/http"
)

func main() {
	const endpoint = "https://jsonplaceholder.typicode.com/photos"

	// create http client
	client := http.Client{}

	// create get request
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Println("error while create get request", err)
	}

	// execute get request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error while execute get request", err)
	}

	// extract response body from client get request
	jsonResponseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while extracting response body", err)
	}

	// unmarshal data from jsonplaceholder
	var photos []jsonModel.Photo
	err = json.Unmarshal(jsonResponseByte, &photos)
	if err != nil {
		fmt.Println("error while unmarshal json byte", err)
	}

	for _, photo := range photos {
		fmt.Println(photo)
	}
}

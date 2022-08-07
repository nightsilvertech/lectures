package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nightsilvertech/lectures/learnhttpmethod/model"
	"io/ioutil"
	"log"
	"net/http"
)

// MethodGet untuk mengambil data dari server
func MethodGet() {

}

// MethodPost untuk membuat data baru ke server
func MethodPost(endpoint string, requestBody []byte) ([]byte, error) {
	// create http client
	client := http.Client{}

	// create post request
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error while create get request %v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// execute post request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while execute get request %v", err)
	}

	// extract response body from client post request
	jsonResponseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error while extracting response body %v", err)
	}

	return jsonResponseByte, nil
}

func main() {
	// send post request to this url
	const endpoint = "https://318c8b57-5d53-4b6b-9ff2-23cb37072d60.mock.pstmn.io/user"

	user := model.User{
		FirstName: "DADANG",
		LastName:  "MANSYUR",
		IsCeo:     true,
		Email:     "udingg@outlook.com",
		Age:       70,
		Products:  []string{"Windows", "Microsoft Office", "Azure"},
		Address: model.Address{
			Province: "XXX",
			District: "xxx",
		},
	}
	bodyReq, _ := json.Marshal(user)

	jsonResponseByte, err := MethodPost(endpoint, bodyReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonResponseByte))
}

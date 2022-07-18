package main

import (
	"encoding/json"
	"fmt"
	"github.com/nightsilvertech/lectures/learnjson/jsonModel"
	"io/ioutil"
)

func main() {
	// JSON
	// JavaScript Object Notation
	// Format pertukaran data antara backend to backend atau backend to frontend
	// Paling populer digunakan saat ini

	// Contoh seolah2 data didapat dari server
	jsonByte, err := ioutil.ReadFile("/home/xoxo/Documents/Go/src/github.com/nightsilvertech/lectures/learnjson/users.json")
	if err != nil {
		fmt.Println("error while opening json file, detail", err)
	}

	// Unmarshal : gunanya mengkonversi data json []byte ke dalam struct
	var users jsonModel.DataUsers
	err = json.Unmarshal(jsonByte, &users)
	if err != nil {
		fmt.Println("error while unmarshal json data, detail", err)
	}

	// Marshal : gunanya mengkonversi struct ke dalam json []byte dan di konvert ke string juga
	jsonByteData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error while marshal json data, detail", err)
	}
	fmt.Println(string(jsonByteData))
}

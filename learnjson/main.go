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
	// Digunakan sebagai format pertukaran antara backend dan frontend

	// contoh nge get data dari server
	jsonByte, err := ioutil.ReadFile("/home/xoxo/Documents/Go/src/github.com/nightsilvertech/lectures/learnjson/users.json")
	if err != nil {
		fmt.Println("Error file tidak ketemu, detail", err)
	}

	// unmarshal : convert data json ke dalam struct
	var users jsonModel.DataUsers
	err = json.Unmarshal(jsonByte, &users)
	if err != nil {
		fmt.Println("Gagal binding ke struct, detail", err)
	}
	fmt.Println("Unmarshal", users)

	// marshal : convert struct ke dalam data json
	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Gagal binding ke json, detail", err)
	}
	fmt.Println("Marshal", string(jsonData))
}

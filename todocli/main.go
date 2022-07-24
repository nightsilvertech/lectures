package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetData(endpoint string) ([]byte, error) {
	// create http client
	client := http.Client{}

	// create get request
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error while create get request %v", err)
	}

	// execute get request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while execute get request %v", err)
	}

	// extract response body from client get request
	jsonResponseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error while extracting response body %v", err)
	}

	return jsonResponseByte, nil
}

type Todos []Todo

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *Todos) Incomplete() (total int64) {
	for _, todo := range *t {
		if todo.Completed == false {
			total += 1
		}
	}
	return
}

func (t *Todos) Complete() (total int64) {
	for _, todo := range *t {
		if todo.Completed == true {
			total += 1
		}
	}
	return
}

func (t *Todos) Count() (total int64) {
	for range *t {
		total += 1
	}
	return
}

func main() {

	const endpoint = "https://jsonplaceholder.typicode.com/todos?userId=%d"

	var userId int64
	flag.Int64Var(&userId, "userId", 0, "-userId=1")
	flag.Parse()

	url := fmt.Sprintf(endpoint, userId)
	jsonData, err := GetData(url)
	if err != nil {
		fmt.Println("error while call GetData", err)
	} else {
		var todos Todos
		err := json.Unmarshal(jsonData, &todos)
		if err != nil {
			fmt.Println("error while Unmarshal data", err)
		}
		fmt.Println("All todos", todos.Count())
		fmt.Println("Incomplete todos", todos.Incomplete())
		fmt.Println("Complete todos", todos.Complete())
	}

}

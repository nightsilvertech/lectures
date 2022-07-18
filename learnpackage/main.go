package main

import (
	"fmt"
	"github.com/nightsilvertech/lectures/learnpackage/model"
	"github.com/nightsilvertech/lectures/learnpackage/model/jsonModel"
)

func main() {
	user1 := model.User{
		Name: "Udin",
		Age:  30,
	}

	job1 := model.Job{
		Title:   "CEO",
		Company: "PT. GAJAH DUDUK",
	}

	json1 := jsonModel.JSON{}

	fmt.Println(user1, job1, json1)
}

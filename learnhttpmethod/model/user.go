package model

type User struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	IsCeo     bool     `json:"isCeo"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Products  []string `json:"products"`
	Address   Address  `json:"address"`
}

type Address struct {
	Province string `json:"province"`
	District string `json:"district"`
}

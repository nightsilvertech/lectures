package jsonModel

type DataUsers struct {
	Users []User `json:"users"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsCeo     bool   `json:"isCeo"`
	Age       int    `json:"age"`
}

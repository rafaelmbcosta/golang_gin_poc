package entity

type Person struct {
	FirstName string `json: "firstName" binding: "required"`
	LastName  string `json: "lastname" binding: "required"`
	Age       int8   `json: "age"`
	Email     string `json: "email"`
}

type Video struct {
	Title       string `json: "title" binding: "min=2, max=10 validate: "is-cool"`
	Description string `json: "description" binding: "max=20"`
	URL         string `json: "url" binding: "required, url"`
	Author      string `json: "author" binding: "required"`
}

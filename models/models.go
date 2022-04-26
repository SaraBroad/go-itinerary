package models

type Item struct {
	Category
}

type Category struct {
	Name string `json:"name"`
}

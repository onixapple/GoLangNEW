package types

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Stock  uint   `json:"stock"`
}

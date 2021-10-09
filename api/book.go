package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) ToJson() []byte {
	toJson, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	return toJson
}

func FromJson(data []byte) Book {
	book := Book{}

	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}

	return book
}

var Books = []Book{
	{Title: "Book 1", Author: "Author 1", ISBN: "00001"},
	{Title: "Book 2", Author: "Author 2", ISBN: "00002"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)

	if err != nil {
		panic(err)
	}

	w.Header().Add("content-type", "application/json; charset=utf-8")

	w.Write(b)

}

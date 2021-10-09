package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book{
	"00001": {Title: "Book 1", Author: "Author 1", ISBN: "00001"},
	"00002": {Title: "Book 2", Author: "Author 2", ISBN: "00002"},
	"00003": {Title: "Book 3", Author: "Author 3", ISBN: "00003"},
	"00004": {Title: "Book 4", Author: "Author 4", ISBN: "00004"},
	"00005": {Title: "Book 5", Author: "Author 5", ISBN: "00005"},
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {

	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:
		performGet(isbn, w)
	case http.MethodPut:
		performUpdate(r, w, isbn)

	case http.MethodDelete:
		performDelete(isbn, w)
	default:
		defaultHandleFunc(w)

	}
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {
	case http.MethodGet:
		performGetAll(w)
	case http.MethodPost:
		performPost(r, w)

	default:
		defaultHandleFunc(w)
	}
}

func defaultHandleFunc(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Unsupported request method."))
}

func performPost(r *http.Request, w http.ResponseWriter) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	book := FromJson(body)
	isbn, created := createBook(book)

	if created {
		w.Header().Add("Location", "/api/books/"+isbn)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}

func performGetAll(w http.ResponseWriter) {
	books := AllBooks()
	writeJson(w, books)
}

func performDelete(isbn string, w http.ResponseWriter) {
	DeleteBook(isbn)
	w.WriteHeader(http.StatusOK)
}

func performUpdate(r *http.Request, w http.ResponseWriter, isbn string) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	book := FromJson(body)
	exists := UpdateBook(isbn, book)

	if exists {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func performGet(isbn string, w http.ResponseWriter) {
	book, found := GetBook(isbn)
	if found {
		writeJson(w, book)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func createBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]

	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]

	if exists {
		books[isbn] = book
	}
	return exists
}

func AllBooks() []Book {
	values := make([]Book, len(books))

	idx := 0

	for _, book := range books {
		values[idx] = book
		idx++
	}

	return values
}

func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]

	return book, found

}

func DeleteBook(isbn string) {
	delete(books, isbn)
}

func (b Book) ToJson() []byte {
	toJson, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}

	return toJson
}

func writeJson(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)

	if err != nil {
		panic(err)
	}
	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.Write(b)
}

func FromJson(data []byte) Book {
	book := Book{}

	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}

	return book
}

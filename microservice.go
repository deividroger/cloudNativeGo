package main

import (
	"net/http"

	"github.com/deividroger/cloudNativeGo/api"
	"github.com/deividroger/cloudNativeGo/conf"
)

func main() {

	http.HandleFunc("/", api.Index)
	http.HandleFunc("/api/echo", api.Echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(conf.GetPort(), nil)
}

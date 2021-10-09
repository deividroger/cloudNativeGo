package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {

	book := Book{Title: "Cloud Native GO", Author: "Test", ISBN: "1234"}

	json := book.ToJson()

	assert.Equal(t, `{"title":"Cloud Native GO","author":"Test","isbn":"1234"}`, string(json), "Book JSON marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {

	json := []byte(`{"title":"Cloud Native GO","author":"Test","isbn":"1234"}`)

	book := FromJson(json)

	assert.Equal(t, Book{Title: "Cloud Native GO", Author: "Test", ISBN: "1234"}, book, "Book Json unmarshalling wrong")
}

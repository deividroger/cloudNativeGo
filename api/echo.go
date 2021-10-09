package api

import (
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("content-Type", "text/plain")
	w.Write([]byte(message))

}

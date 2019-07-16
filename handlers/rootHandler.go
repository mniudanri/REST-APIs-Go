package handlers

import (
	"net/http"
)

//root handler
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path != "/"){
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Url Not Found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1!\n"))
}
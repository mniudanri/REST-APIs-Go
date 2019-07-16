package handlers

import (
	"encoding/json"
	"net/http"
)

type jsonResponse map[string]interface{}

func postError(w http.ResponseWriter, code int){
	http.Error(w, http.StatusText(code), code)
}

func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse){
	if content != nil {
		js, err := json.Marshal(content)

		if err != nil {
			postError(w, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return 
	}
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
	return 
}
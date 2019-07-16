package handlers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

func UsersRouter (w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			employeeGetAll(w, r)
			return
		case http.MethodPost:
			createEmployee(w, r)
			return
		default: 
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}
	id := bson.ObjectIdHex(path)
	// // w.WriteHeader(http.StatusNotFound)
	// // w.Write([]byte(http.StatusText(400)))
	// // return 

	switch r.Method {
	case http.MethodGet:
		employeeGetById(w, r, id)
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}

}
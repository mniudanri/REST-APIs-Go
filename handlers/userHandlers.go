package handlers

import (
	"net/http"
	"REST-APIs-Go/user"
	"io/ioutil"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"github.com/asdine/storm"
)

func employeeGetAll(w http.ResponseWriter, r *http.Request){
	users, err := user.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

func bodyToUser(r *http.Request, u *user.Employee) error {
	if r.Body == nil {
		return errors.New("Request body empty")
	}
	if u == nil{
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)
}

func createEmployee(w http.ResponseWriter, r *http.Request){
	u := new(user.Employee)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		}else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location","/users/"+u.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func employeeGetById(w http.ResponseWriter, _ * http.Request, id bson.ObjectId) {
	u, err := user.GetById(id)
	if err != nil{
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}
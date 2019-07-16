package user

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/asdine/storm"
)

type Employee struct {
	ID bson.ObjectId `json:"id" storm"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

const (
	dbPath = "users.db"
)

var (
	ErrRecordInvalid = errors.New("Name is required")
)

func All()([]Employee, error){
	db, err := storm.Open(dbPath)

	if err != nil{
		return nil, err
	}

	defer db.Close()
	emps := []Employee{}

	err = db.All(&emps)

	if err != nil {
		return nil, err
	}

	return emps, nil
}

//retrive data by id user
func GetById(id bson.ObjectId) (*Employee, error) {
	db, err := storm.Open(dbPath)

	if err != nil{
		return nil, err
	}

	defer db.Close()
	u := new(Employee)
	err = db.One("ID", id, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)

	if err != nil{
		return err
	}

	defer db.Close()
	u := new(Employee)
	err = db.One("ID", id, u)

	if err != nil {
		return err
	}	

	return db.DeleteStruct(u)
}

//save updates or creates record to db
func (u *Employee) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)

	if err != nil{
		return err
	}

	defer db.Close()
	
	return db.Save(u)
}

//validate record
func (u *Employee) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}

	return nil
}
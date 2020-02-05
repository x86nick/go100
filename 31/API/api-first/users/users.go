package user

import (
	"errors"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// User holds data for a single user

type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "users.db"
)

var (
	ErrRecordInvalid = errors.New("Record is Invalid")
)

// All retrives all users from db
func All() ([]User, error) {
	// 1. open db
	db, err := storm.Open(dbPath)
	// in case error is not nil, then return it to caller
	if err != nil {
		return nil, err
	}
	// 2. make sure connection is close
	defer db.Close()
	// 3. create list of users
	users := []User{}
	err = db.All(&users) // using pointer the slice
	if err != nil {
		return nil, err
	}
	return users, nil
}

// add second read function, as an argument pass id, and return a pointer to the user struct
// and an error

// One returns a single user record from the database
func One(id bson.ObjectId) (*User, error) {
	// 1. open db
	db, err := storm.Open(dbPath)
	// in case error is not nil, then return it to caller
	if err != nil {
		return nil, err
	}
	// 2. make sure connection is close
	defer db.Close()
	// 3. create a single user
	u := new(User)
	// assign resuts to err variable
	// the u keyword does return a pointer to the structure, so there's no need to add the reference ampersand in here.
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil // if no error, then return user and no error
}

// Create and update functions will be handled by a single member of the user struct
// that we will call safe

// Delete removes a given record from the database
// it needs an ID as an argument, and returns an error.
func Delete(id bson.ObjectId) error {
	// 1. open db
	db, err := storm.Open(dbPath)
	// in case error is not nil, then return it to caller
	if err != nil {
		return err
	}
	// 2. make sure connection is close
	defer db.Close()
	// 3. create a single user
	u := new(User)
	// assign resuts to err variable
	// the u keyword does return a pointer to the structure, so there's no need to add the reference ampersand in here.
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// Save function takes care of creating and updating the record in database
// This function is going to be a member of the user struct and will just return an error.

func (u *User) Save() error {
	// this function do assignment upto ; then checks for condition, and execcute code inside {} if condition is true.
	if err := u.validate(); err != nil {
		return err
	}
	// 1. open db
	db, err := storm.Open(dbPath)
	// in case error is not nil, then return it to caller
	if err != nil {
		return err
	}
	// 2. make sure connection is close
	defer db.Close()
	return db.Save(u)
}

// create a user member function to validate some data
//  let's say that the username cannot be empty.
// this function will return an error and will not take an arguments.
func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid // declare it above
	}
	return nil
}

package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm"
	user "github.com/x86nick/go100/31/API/api-first/users"
	"gopkg.in/mgo.v2/bson"
)

// supply 2 argumetns, 1. request itself, 2. pointer to the user item
func bodyTouser(r *http.Request, u *user.User) error {
	// take care of edge cases
	if r.Body == nil { // if request body is nil, return an error
		return errors.New("request body is empty")
	}
	// if a user poiter is nil, return an error
	if u == nil {
		return errors.New("a user is required")
	}
	// now read the request body using helper function ioutil
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// And finally we can unmarshal what we read into the supplied structure. then write the post handler
	return json.Unmarshal(bd, u)
}

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	// Retrive all the request from the collection
	users, err := user.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

func usersPostOne(w http.ResponseWriter, r *http.Request) {
	// first we need a new user record
	u := new(user.User)
	err := bodyTouser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	// since this is a new records, we need to assign a new id
	u.ID = bson.NewObjectId()
	err = u.Save() // call save method of user object and it will return an error,
	//  2 types of errors here 1. associated with db, 2. assoisabled with failed validation
	if err != nil {
		//  If the error is failed validation we need to notify other about the bad request,
		// otherwise we need to send a 500 internal server error
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		// And in both cases, terminate dysfunction by calling return
		return
	}
	// if everything went well their code was added, the customer response is to send the location
	// of the new resource as part of the header and that to a one created code. First let's add the header.
	//The field is called Location, and the value is relative URL to the created record,
	//  in our case "/users/"+u.ID.Hex()
	w.Header().Set("Location", "/users/"+u.ID.Hex())
	// then just write the code in the header
	w.WriteHeader(http.StatusCreated)
	// now add the call to the handler in usersRouter
}

// Retrieving an ITEM
// this handler does not need any data from the request, you can put _ in place of r in `r *http.Request`
func usersGetOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	// retrieve the data from database
	// call the functio user.One and pass the id
	u, err := user.One(id)

	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		// anything else means internal server error code
		postError(w, http.StatusInternalServerError)
		return
	}
	// finally if no error is return, means user data is ready and we can post it as JSON body using POST body response
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u}) // now go to usersRouter.go put id := bason.ObjectIdHex(path)
}

// user PutOne function
// id bson.ObjectId assign this id to the parsed user record instead of generating a new one

func usersPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {

	// first we need a new user record
	u := new(user.User)

	err := bodyTouser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	//  // -- since this is a new records, we need to assign a new id
	// Assign this id to the parsed user record instead of generating a new one
	u.ID = id
	err = u.Save() // call save method of user object and it will return an error,
	//  2 types of errors here 1. associated with db, 2. assoisabled with failed validation
	if err != nil {
		//  If the error is failed validation we need to notify other about the bad request,
		// otherwise we need to send a 500 internal server error
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		// And in both cases, terminate dysfunction by calling return
		return
	}
	// // if everything went well their code was added, the customer response is to send the location
	// // of the new resource as part of the header and that to a one created code. First let's add the header.
	// //The field is called Location, and the value is relative URL to the created record,
	// //  in our case "/users/"+u.ID.Hex()
	// w.Header().Set("Location", "/users/"+u.ID.Hex())
	// // then just write the code in the header
	// w.WriteHeader(http.StatusCreated)
	// // now add the call to the handler in usersRouter

	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})

}

// - Patch - we need to retrieve existing records

func usersPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {

	u, err := user.One(id)

	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		// anything else means internal server error code
		postError(w, http.StatusInternalServerError)
		return
	}

	err = bodyTouser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}

	// Assign this id to the parsed user record instead of generating a new one
	u.ID = id
	err = u.Save() // call save method of user object and it will return an error,

	if err != nil {

		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		// And in both cases, terminate dysfunction by calling return
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})

}

// DELETE

func usersDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {

	err := user.Delete(id)

	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		// anything else means internal server error code
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

// UsersRouter handles the users route
func UsersRouter(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.Path) // Print the full request path
	// There are going to be 2 types of requests
	// We'll start by trimming the excessive slash at the end,
	//if for some reason we got added to the request,
	// the strings library has the function trim suffix
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			usersPostOne(w, r)
			return
		case http.MethodHead:
			usersPostOne(w, r)
			return
		default:
			// it will take response write and status code
			// API errors should never return the contents of the errors to end user
			// you don't want to expose inner working outside
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}
	path = strings.TrimPrefix(path, "/users")
	if !bson.IsObjectIdHex(path) { // verify id using bson.IsObjectIdHex function, if id is not valid, post 404 not found
		postError(w, http.StatusNotFound)
		return

	}
	// parse id into a sepeate varialbe using bson.ObjectIdHex()
	// and we got resources identified

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		usersGetOne(w, r, id)
		return
	case http.MethodPut:
		usersPutOne(w, r, id)
		return
	case http.MethodPatch:
		usersPatchOne(w, r, id)
		return
	case http.MethodDelete:
		usersDeleteOne(w, r, id)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}

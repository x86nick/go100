package handlers

import "net/http"

// 3. create private rooHandler function
// The handler function needs two arguments:
// the first one is an interface named http.ResponseWriter,
// and the second one is a pointer to HTTP request *http.Request.

// RootHandler handles the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// 4.
	if r.URL.Path != "/" { // if URL is different then "/" then return a different error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)        // send status code
	w.Write([]byte("Running API v1\n")) // output message with API version
}

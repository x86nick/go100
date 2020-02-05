package handlers

import "net/http"

import "encoding/json"

//  It's a string indexed map of empty interfaces.
//  This way we ensure that each element has its own uniquely named property.
// empty interface - Even though it can hold anything, you can just use it as a variable of any type, like in Python or JavaScript.
type jsonResponse map[string]interface{}

// postError is a private function and will take responsWrite and error code as argument
// it will invoke the http.Error function passing as the argument the ResponseWriter, the StatusText of the error code
// which we can obtain by http.StatusText(code)
func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

// postBodyResponse function will take the ResponseWriter code and the content as arguments.
//  As far as the content-type goes, we could just use an empty an interface here, which is Go's equivalent of a container of any type,
// however we want to enforce the naming of the properties,
// and the best way to do this is to actually create our custom type. Let's call it JSON response
func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
	// ensure that content is not nil
	if content != nil {
		js, err := json.Marshal(content) // invoke json.Marshal to convert the content to a JSON byte sequence.
		if err != nil {
			postError(w, http.StatusInternalServerError) // if error occure on the way, handle them as customery 500 internal errors...
			return
		}
		// ...Then we need to set the response header to reflect the type of the response.
		w.Header().Set("Content-Type", "application/json")
		// now write the status code into header
		w.WriteHeader(code)
		// finally pass the json byte sequence to the writer
		w.Write(js)
		return
	}
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
}

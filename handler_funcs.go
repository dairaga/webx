package webx

import "net/http"

var (
	// NotFound page not found (404)
	NotFound = http.NotFound
)

// -----------------------------------------------------------------------------

// BadRequest handler function for Bad Request.
func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "bad request", http.StatusBadRequest)
}

// Forbidden handler function for Forbidden.
func Forbidden(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "forbidden", http.StatusForbidden)
}

// MethodNotAllowed handler function for Method Not Allowed.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

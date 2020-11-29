package webx

import (
	"net/http"
	"reflect"
)

// route ...
type route struct {
	pattern  string
	types    []reflect.Type
	redirect bool
}

var _ http.Handler = &route{}

// -----------------------------------------------------------------------------

func newRoute(pattern string, types []reflect.Type, handler interface{}, redirect bool) *route {

	return &route{
		pattern:  pattern,
		types:    types,
		redirect: redirect,
	}
}

// -----------------------------------------------------------------------------

func (rt *route) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// -----------------------------------------------------------------------------

func validFunc(handler interface{}) ([]reflect.Type, bool) {
	return nil, false
}

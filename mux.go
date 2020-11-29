package webx

import (
	"net/http"
	"strings"
)

// Mux ...
type Mux struct {
	routes           map[string][]*route
	NotFound         http.Handler
	BadRequest       http.Handler
	Forbidden        http.Handler
	MethodNotAllowed http.Handler
}

var _ http.Handler = &Mux{}

// -----------------------------------------------------------------------------

// NewMux returns a router
func NewMux() *Mux {
	return &Mux{
		routes: make(map[string][]*route),
	}
}

// -----------------------------------------------------------------------------

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// -----------------------------------------------------------------------------

// Handle ...
func (m *Mux) Handle(method, pattern string, handler interface{}) {
	method = strings.TrimSpace(method)
	pattern = strings.TrimSpace(pattern)
	if method == "" || pattern == "" || handler == nil {
		return
	}

	types, ok := validFunc(handler)
	if !ok {
		return
	}

}

// -----------------------------------------------------------------------------

// Get ...
func (m *Mux) Get(pattern string, handler interface{}) {
	m.Handle(http.MethodGet, pattern, handler)
}

// -----------------------------------------------------------------------------

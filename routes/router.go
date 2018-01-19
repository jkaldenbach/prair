package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	mw "kaldenbach.design/prayr/middleware"
)

// NewRouter : creates a new url router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = mw.Headers(handler)
		handler = mw.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

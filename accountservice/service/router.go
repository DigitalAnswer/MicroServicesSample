package service

import (
	"github.com/gorilla/mux"
)

// Function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}

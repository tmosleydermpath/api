package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

<<<<<<< HEAD
// NewRouter creates a new routing tree based off of routes
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

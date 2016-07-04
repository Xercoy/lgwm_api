package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes
type Routes []Route

// NewRouter iterates through the static list of routes and adds them to the new Router.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Posts",
		"GET",
		"/posts",
		PostsIndex,
	},
	Route{
		"GET Post",
		"GET",
		"/posts/{id}",
		GETPost,
	},
}

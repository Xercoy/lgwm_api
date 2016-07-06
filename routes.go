package lgwm_api

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
	// Index, displays information.
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	// Get a list of all the Posts.
	Route{
		"Posts",
		"GET",
		"/posts",
		PostHandler,
	},
	// CREATE
	// Create a new Post.
	Route{
		"CREATE Post",
		"POST",
		"/posts",
		PostHandler,
	},
	// READ
	// Retrieve a specific Post.
	Route{
		"GET Post",
		"GET",
		"/posts/{id}",
		PostHandler,
	},
	// UPDATE
	// Update a specific Post.
	Route{
		"UPDATE Post",
		"PUT",
		"/posts/{id}",
		PostHandler,
	},
	// DELETE
	// Delete a particular post.
	Route{
		"DELETE Post",
		"DELETE",
		"/posts/{id}",
		PostHandler,
	},
}

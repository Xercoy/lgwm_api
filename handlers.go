package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `Rest API for http://learngowith.me
POSTS      - GET  - Shows this page.
POSTS      - POST - Creates a new blog post.
POSTS/{id} - GET  - Retrieves a new blog post.
`)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id == 0 { // /posts
		switch r.Method {

		// Show all posts
		case "GET":
			json.NewEncoder(w).Encode(Posts)

		// Create a new post
		case "POST":
			fmt.Fprintf(w, "Created Post")
		}
	}

	if id != 0 { // /posts/{id}
		switch r.Method {

		// Show post with the particular matching ID.
		case "GET":
			for _, p := range Posts {
				if p.ID == id {
					json.NewEncoder(w).Encode(p)
					return
				}
			}
		case "PUT":
			fmt.Fprintf(w, "Updated Post")

		case "DELETE":
			fmt.Fprintf(w, "Deleted Post")
		default:
			fmt.Fprintln(w, "Can't find post. Sorry.\n")
		}

	}
}

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

func PostsIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Posts)
}

func GETPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	for _, p := range Posts {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	fmt.Fprintln(w, "Can't find post. Sorry.\n")
}

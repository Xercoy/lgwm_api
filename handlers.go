package lgwm_api

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
			posts, err := GetAllPosts(GlobalDB)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				json.NewEncoder(w).Encode(posts)
			}
		// Create a new post
		case "POST":
			var p Post

			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				// Received.
				fmt.Fprintf(w, "%v", err)
			}

			err = AddPost(GlobalDB, GlobalDB.Bucket, p)
			if err != nil {
				// Received.
				fmt.Fprintf(w, "%v", err)
			} else {
				json.NewEncoder(w).Encode(p)
			}
		}
	}

	if id != 0 { // /posts/{id}
		switch r.Method {

		// Show post with the particular matching ID.
		case "GET":

		case "PUT":
			fmt.Fprintf(w, "Updated Post")

		case "DELETE":
			fmt.Fprintf(w, "Deleted Post")
		default:
			fmt.Fprintln(w, "Can't find post. Sorry.\n")
		}

	}
}

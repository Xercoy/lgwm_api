package main

import "time"

type Post struct {
	Title  string
	Body   string
	Date   time.Time
	Author string
	ID     int
	tag    string
}

var Posts = []Post{
	Post{
		Title:  "How I met your mother",
		Body:   "At the park, it was nice.",
		Date:   time.Now(),
		Author: "Corey Prak",
		ID:     1,
	},
	Post{
		Title:  "How I survived in the forest",
		Body:   "I hid in a tree",
		Date:   time.Now(),
		Author: "Corey Prak",
		ID:     2,
	},
}

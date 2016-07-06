package lgwm_api

import "time"

type Post struct {
	Title  string
	Body   string
	Date   time.Time
	Author string
	ID     int
	Tag    string
}

func NewPost(t, b, a, ta string, d time.Time, i int) *Post {
	p := Post{
		Title:  t,
		Body:   b,
		Date:   d,
		Author: a,
		ID:     i,
		Tag:    ta,
	}

	return &p
}

package search

import "context"

type Query struct {
	Keywords string
	JobTitle string
	Company  string
	Location string
}

type Profile struct {
	Name string
	URL  string
}

type Searcher interface {
	Search(ctx context.Context, q Query) ([]Profile, error)
}

package main

import (
	"context"
)

// Fetcher interface helps you to fetch the database datasource name fetcher
type Fetcher interface {
	// Fetch should returns the database datasource name / error
	Fetch(ctx context.Context) (dsn string, err error)
}

// FetcherFunc is a single function form of Fetcher
type FetcherFunc func(ctx context.Context) (dsn string, err error)

func (f FetcherFunc) Fetch(ctx context.Context) (dsn string, err error) {
	return f(ctx)
}

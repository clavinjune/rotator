package main

import (
	"database/sql/driver"
)

// Opt is used for registering a new driver
type Opt struct {
	// MaxRetry is used for defining how many times the fetcher will be retried
	MaxRetry int

	// DriverName is used for identifying the custom drivfer
	DriverName string

	// DriverBase is used for opening the database connection
	DriverBase driver.Driver

	// Fetcher is used for fetching the database datasource name
	Fetcher Fetcher
}

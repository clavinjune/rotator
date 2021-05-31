package rotator

import (
	"database/sql/driver"
)

// Opt is used for registering a new driver
type Opt struct {
	// MaxRetry is used for defining how many times the fetcher will be retried
	// if there something wrong when fetching / re-open the connection.
	// If MaxRetry is less than FetcherDefaultMaxRetry, Opt will use FetcherDefaultMaxRetry instead.
	MaxRetry int

	// DriverName is used for identifying the custom driver
	DriverName string

	// DriverBase is used for opening the database connection
	DriverBase driver.Driver

	// Fetcher is used for fetching the database datasource name
	Fetcher Fetcher
}

package main

import (
	"database/sql"
)

const (
	// maxRetry defined the least retry number
	maxRetry int = 2
)

// RegisterRotationDriver registers the custom driver with defined opt
func RegisterRotationDriver(opt Opt) {
	validate(opt)

	sql.Register(opt.DriverName, &rotator{
		base:     opt.DriverBase,
		fetch:    opt.Fetcher.Fetch,
		maxRetry: getMaxRetry(opt),
	})
}

// validate validates the opt
func validate(opt Opt) {
	if opt.DriverName == "" {
		panic("rotator: Driver name is empty")
	}

	if opt.DriverBase == nil {
		panic("rotator: Register driver base is nil")
	}

	if opt.Fetcher == nil {
		panic("rotator: Fetcher is nil")
	}
}

// getMaxRetry gets the maxRetry from the opt
// getMaxRetry will returns 2 at least
func getMaxRetry(opt Opt) int {
	if opt.MaxRetry > maxRetry {
		return opt.MaxRetry
	}

	return maxRetry
}

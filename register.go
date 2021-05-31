package rotator

import (
	"database/sql"
)

// RegisterRotationDriver registers the custom driver with the given opt
// RegisterRotationDriver may trigger panic if the given opt is invalid
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

// getMaxRetry gets the fetcherMaxRetry from the opt
// getMaxRetry will returns 2 at least
func getMaxRetry(opt Opt) int {
	if opt.MaxRetry > fetcherMaxRetry {
		return opt.MaxRetry
	}

	return fetcherMaxRetry
}

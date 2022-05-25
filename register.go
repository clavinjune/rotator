// Copyright 2022 clavinjune/rotator
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rotator

import (
	"database/sql"
	"errors"
)

const (
	// FetcherDefaultMaxRetry defined the least retry number
	FetcherDefaultMaxRetry int = 2
)

var (
	// ErrEmptyDriverName returns when Opt.DriverName is empty
	ErrEmptyDriverName = errors.New("rotator: driver name is empty")

	// ErrEmptyDriverBase returns when Opt.DriverBase is nil
	ErrEmptyDriverBase = errors.New("rotator: driver base is empty")

	// ErrEmptyFetcher returns when Opt.Fetcher is nil
	ErrEmptyFetcher = errors.New("rotator: fetcher is nil")
)

// RegisterRotationDriver registers the custom driver with the given opt.
func RegisterRotationDriver(opt Opt) error {
	if err := validate(opt); err != nil {
		return err
	}

	sql.Register(opt.DriverName, &rotator{
		base:     opt.DriverBase,
		fetch:    opt.Fetcher.Fetch,
		maxRetry: getMaxRetry(opt),
	})

	return nil
}

// MustRegisterRotationDriver calls RegisterRotationDriver
// may trigger panic if the given opt is invalid
func MustRegisterRotationDriver(opt Opt) {
	if err := RegisterRotationDriver(opt); err != nil {
		panic(err)
	}
}

// validate validates the opt
func validate(opt Opt) error {
	if opt.DriverName == "" {
		return ErrEmptyDriverName
	}

	if opt.DriverBase == nil {
		return ErrEmptyDriverBase
	}

	if opt.Fetcher == nil {
		return ErrEmptyFetcher
	}

	return nil
}

// getMaxRetry gets the FetcherDefaultMaxRetry from Opt
// getMaxRetry will return 2 at least
func getMaxRetry(opt Opt) int {
	if opt.MaxRetry > FetcherDefaultMaxRetry {
		return opt.MaxRetry
	}

	return FetcherDefaultMaxRetry
}

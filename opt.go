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
	"database/sql/driver"
)

// Opt is used for registering a new driver
type Opt struct {
	// MaxRetry is used for defining how many times the fetcher will be retried
	// if there is something wrong when fetching / re-open the connection.
	// If MaxRetry is less than FetcherDefaultMaxRetry, Opt will use FetcherDefaultMaxRetry instead.
	MaxRetry int

	// DriverName is used for identifying the custom driver
	DriverName string

	// DriverBase is used for opening the database connection
	DriverBase driver.Driver

	// Fetcher is used for fetching the database datasource name
	Fetcher Fetcher
}

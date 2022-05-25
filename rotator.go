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
	"context"
	"database/sql/driver"
)

// rotator helps base driver to fetch credentials from fetch function with defined FetcherDefaultMaxRetry
type rotator struct {
	base     driver.Driver
	fetch    FetcherFunc
	maxRetry int
}

// Open opens connection from the fetched Datasource name
func (r *rotator) Open(_ string) (driver.Conn, error) {
	var dc driver.Conn
	var err error

	for i := 0; i < r.maxRetry; i++ {
		var dsn string
		dsn, err = r.fetch(context.Background())
		if err != nil {
			continue
		}

		dc, err = r.base.Open(dsn)
		if err == nil {
			break
		}
	}

	return dc, err
}

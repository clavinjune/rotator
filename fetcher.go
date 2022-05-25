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
)

// Fetcher interface helps you to fetch the database datasource name fetcher
type Fetcher interface {
	// Fetch should return the database datasource name / error
	Fetch(ctx context.Context) (dsn string, err error)
}

// FetcherFunc is a single function form of Fetcher
type FetcherFunc func(ctx context.Context) (dsn string, err error)

// Fetch implements FetcherFunc
func (f FetcherFunc) Fetch(ctx context.Context) (dsn string, err error) {
	return f(ctx)
}

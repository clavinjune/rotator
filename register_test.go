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

package rotator_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/ClavinJune/rotator"
)

func TestRegisterRotationDriver(t *testing.T) {
	tt := []struct {
		opt  rotator.Opt
		want error
	}{{
		opt:  rotator.Opt{},
		want: rotator.ErrEmptyDriverName,
	}, {
		opt: rotator.Opt{
			MaxRetry:   1,
			DriverName: "test1",
			DriverBase: nil,
			Fetcher:    nil,
		},
		want: rotator.ErrEmptyDriverBase,
	}, {
		opt: rotator.Opt{
			MaxRetry:   3,
			DriverName: "test2",
			DriverBase: openFunc(func(name string) (driver.Conn, error) {
				return nil, nil
			}),
			Fetcher: nil,
		},
		want: rotator.ErrEmptyFetcher,
	}, {
		opt: rotator.Opt{
			MaxRetry:   2,
			DriverName: "test3",
			DriverBase: openFunc(func(name string) (driver.Conn, error) {
				return nil, nil
			}),
			Fetcher: rotator.FetcherFunc(func(ctx context.Context) (dsn string, err error) {
				return "", nil
			}),
		},
		want: nil,
	}}

	for _, test := range tt {
		innerPanicTest(t, test.opt, test.want)
	}
}

func innerPanicTest(t *testing.T, opt rotator.Opt, want error) {
	fmt.Println(opt.MaxRetry)
	defer func() {
		if got := recover(); got != nil {
			if got != want {
				t.Fatalf(`got "%v", want "%v"`, got, want)
			}
		} else {
			if want != nil {
				t.Fatalf(`got "nil", want "%v"`, want)
			}
		}
	}()

	rotator.MustRegisterRotationDriver(opt)
}

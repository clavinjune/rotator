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
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"

	"github.com/ClavinJune/rotator"
)

const (
	testDSN       string = "testDSN"
	testFailedDSN string = "failedDSN"
)

func TestRotator_OpenMaxRetry3Retry4Times_ExpectError(t *testing.T) {
	counter := 0
	want := errors.New("failed mocking")

	rotator.MustRegisterRotationDriver(rotator.Opt{
		MaxRetry:   3,
		DriverName: "TestRotator_OpenMaxRetry3Retry4Times_ExpectError",
		DriverBase: openFunc(func(name string) (driver.Conn, error) {
			if name != testDSN {
				return nil, want
			}
			return nil, nil
		}),
		Fetcher: rotator.FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			fmt.Println("called")
			counter++
			if counter < 4 {
				return testFailedDSN, nil
			}
			return testDSN, nil
		}),
	})

	db, _ := sql.Open("TestRotator_OpenMaxRetry3Retry4Times_ExpectError", "")

	if got := db.Ping(); got != want {
		t.Fatalf(`got "%v", want "%v"`, got, want)
	}
}

func TestRotator_OpenMaxRetry3Retry3Times_ExpectSuccess(t *testing.T) {
	counter := 0
	want := errors.New("failed mocking")

	rotator.MustRegisterRotationDriver(rotator.Opt{
		MaxRetry:   3,
		DriverName: "TestRotator_OpenMaxRetry3Retry3Times_ExpectSuccess",
		DriverBase: openFunc(func(name string) (driver.Conn, error) {
			if name != testDSN {
				return nil, want
			}
			return nil, nil
		}),
		Fetcher: rotator.FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			fmt.Println("called")
			counter++
			if counter < 3 {
				return testFailedDSN, nil
			}
			return testDSN, nil
		}),
	})

	db, _ := sql.Open("TestRotator_OpenMaxRetry3Retry3Times_ExpectSuccess", "")
	if got := db.Ping(); got != nil {
		t.Fatalf(`got "%v", want "nil"`, got)
	}
}

func TestRotator_FetchRetry_ExpectError(t *testing.T) {
	counter := 0
	want := errors.New("failed fetching")

	rotator.MustRegisterRotationDriver(rotator.Opt{
		MaxRetry:   3,
		DriverName: "TestRotator_FetchRetry_ExpectError",
		DriverBase: openFunc(func(name string) (driver.Conn, error) {
			if name != testDSN {
				return nil, want
			}
			return nil, nil
		}),
		Fetcher: rotator.FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			fmt.Println("called")
			counter++
			if counter < 4 {
				return testDSN, want
			}
			return testDSN, nil
		}),
	})

	db, _ := sql.Open("TestRotator_FetchRetry_ExpectError", "")
	if got := db.Ping(); got != want {
		t.Fatalf(`got "%v", want "%v"`, got, want)
	}
}

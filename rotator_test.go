package rotator

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"testing"
)

func TestRotator_OpenMaxRetry3Retry4Times_ExpectError(t *testing.T) {
	counter := 0
	want := errors.New("failed mocking")

	RegisterRotationDriver(Opt{
		MaxRetry:   3,
		DriverName: "TestRotator_OpenMaxRetry3Retry4Times_ExpectError",
		DriverBase: openFunc(func(name string) (driver.Conn, error) {
			if name != "test" {
				return nil, want
			}
			return nil, nil
		}),
		Fetcher: FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			fmt.Println("called")
			counter++
			if counter < 4 {
				return "failed", nil
			}
			return "test", nil
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

	RegisterRotationDriver(Opt{
		MaxRetry:   3,
		DriverName: "TestRotator_OpenMaxRetry3Retry3Times_ExpectSuccess",
		DriverBase: openFunc(func(name string) (driver.Conn, error) {
			if name != "test" {
				return nil, want
			}
			return nil, nil
		}),
		Fetcher: FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			fmt.Println("called")
			counter++
			if counter < 3 {
				return "failed", nil
			}
			return "test", nil
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

	RegisterRotationDriver(Opt{
		MaxRetry:   3,
		DriverName: "TestRotator_FetchRetry_ExpectError",
		DriverBase: openFunc(func(name string) (driver.Conn, error) {
			if name != "test" {
				return nil, want
			}
			return nil, nil
		}),
		Fetcher: FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			fmt.Println("called")
			counter++
			if counter < 4 {
				return "test", want
			}
			return "test", nil
		}),
	})

	db, _ := sql.Open("TestRotator_FetchRetry_ExpectError", "")
	if got := db.Ping(); got != want {
		t.Fatalf(`got "%v", want "%v"`, got, want)
	}
}

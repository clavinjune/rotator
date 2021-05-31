package rotator

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"
)

func TestRegisterRotationDriver(t *testing.T) {
	tt := []struct {
		opt  Opt
		want string
	}{{
		opt:  Opt{},
		want: "rotator: Driver name is empty",
	}, {
		opt: Opt{
			MaxRetry:   1,
			DriverName: "test1",
			DriverBase: nil,
			Fetcher:    nil,
		},
		want: "rotator: Register driver base is nil",
	}, {
		opt: Opt{
			MaxRetry:   3,
			DriverName: "test2",
			DriverBase: openFunc(func(name string) (driver.Conn, error) {
				return nil, nil
			}),
			Fetcher: nil,
		},
		want: "rotator: Fetcher is nil",
	}, {
		opt: Opt{
			MaxRetry:   2,
			DriverName: "test3",
			DriverBase: openFunc(func(name string) (driver.Conn, error) {
				return nil, nil
			}),
			Fetcher: FetcherFunc(func(ctx context.Context) (dsn string, err error) {
				return "", nil
			}),
		},
		want: "",
	}}

	for _, test := range tt {
		innerPanicTest(t, test.opt, test.want)
	}
}

func innerPanicTest(t *testing.T, opt Opt, want string) {
	fmt.Println(opt.MaxRetry)
	defer func() {
		if got := recover(); got != nil {
			if got != want {
				t.Fatalf(`got "%v", want "%v"`, got, want)
			}
		} else {
			if want != "" {
				t.Fatalf(`got "nil", want "%v"`, want)
			}
		}
	}()

	RegisterRotationDriver(opt)
}

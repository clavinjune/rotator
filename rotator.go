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

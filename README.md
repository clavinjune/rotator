# rotator [![Go Reference](https://pkg.go.dev/badge/github.com/ClavinJune/rotator.svg)](https://pkg.go.dev/github.com/ClavinJune/rotator)
Helper database driver for databases that use credential rotation

## Usage

```bash
go get -u github.com/ClavinJune/rotator@latest
```

## Example

```go
package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ClavinJune/rotator"
	"github.com/lib/pq"
)

func main() {
	rotator.RegisterRotationDriver(rotator.Opt{
		MaxRetry:   3,
		DriverName: "example",
		DriverBase: &pq.Driver{},
		Fetcher: rotator.FetcherFunc(func(ctx context.Context) (dsn string, err error) {
			// fetch and return your datasource name here
			return "postgres://user:password@localhost:5432/dbname", nil
		}),
	})

	db, err := sql.Open("example", "this section will automatically filled by Fetcher func")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(3 * time.Second)

	for i := 0; i < 10; i++ {
		if err := db.Ping(); err != nil {
			panic(err)
		}
		log.Println("connected")
		time.Sleep(time.Second)
	}
}
```

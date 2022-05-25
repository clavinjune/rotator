# rotator [![Go Reference](https://pkg.go.dev/badge/github.com/ClavinJune/rotator.svg)](https://pkg.go.dev/github.com/ClavinJune/rotator)
Helper database driver for databases that use credential rotation (e.g. [Vault](https://learn.hashicorp.com/tutorials/vault/getting-started-dynamic-secrets) / [Heroku](https://devcenter.heroku.com/articles/connecting-to-heroku-postgres-databases-from-outside-of-heroku#credentials)).
By using [rotator](https://pkg.go.dev/github.com/ClavinJune/rotator), your choice of database driver will be wrapped,
so the database datasource name will be dynamically rotated depends on how you want it to be rotated.  

## Usage

```bash
go get -u github.com/clavinjune/rotator@latest
```

## Example

```go
package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/clavinjune/rotator"
	"github.com/lib/pq"
)

func main() {
	// Register your wrapper driver
	rotator.RegisterRotationDriver(rotator.Opt{
		// set Max Retry if there's something wrong when fetching / reopening the connection
		MaxRetry:   3,
		// set your uniq custom DriverName
		DriverName: "example",
		// set your choice of database driver 
		DriverBase: &pq.Driver{},
		// set your way to fetch the DSN
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

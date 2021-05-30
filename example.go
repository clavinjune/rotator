package rotator

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	heroku "github.com/heroku/heroku-go/v5"
	"github.com/lib/pq"
)

func getHerokuSvc(username, password string) *heroku.Service {
	heroku.DefaultTransport.Username = username
	heroku.DefaultTransport.Password = password
	return heroku.NewService(heroku.DefaultClient)
}

func getHerokuFetcher(svc *heroku.Service, appName string) FetcherFunc {
	return func(ctx context.Context) (dsn string, err error) {
		log.Println("heroku fetcher called")
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		m, err := svc.ConfigVarInfoForApp(ctx, appName)
		if err != nil {
			return "", err
		}

		if v, ok := m["DATABASE_URL"]; ok {
			return *v, nil
		}

		return "", errors.New("missing DATABASE_URL")
	}
}

func main() {
	svc := getHerokuSvc(os.Getenv("HEROKU_USERNAME"), os.Getenv("HEROKU_PASSWORD"))
	fetcher := getHerokuFetcher(svc, os.Getenv("HEROKU_APPNAME"))

	RegisterRotationDriver(Opt{
		MaxRetry:   3,
		DriverName: "testing",
		DriverBase: &pq.Driver{},
		Fetcher:    fetcher,
	})

	db, err := sql.Open("testing", "no need to fill this")
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

package rotator

import (
	"database/sql/driver"
)

type openFunc func(name string) (driver.Conn, error)

func (o openFunc) Open(name string) (driver.Conn, error) {
	return o(name)
}

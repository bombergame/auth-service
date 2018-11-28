package mysql

import (
	"fmt"
	"github.com/bombergame/auth-service/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	str string
	db  *sqlx.DB
}

func NewConnection() *Connection {
	return &Connection{
		str: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			config.StorageUser, config.StoragePassword,
			config.StorageHost, config.StoragePort, config.StorageName,
		),
	}
}

func (c *Connection) Open() error {
	var err error

	c.db, err = sqlx.Connect("mysql", c.str)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Close() error {
	return c.db.Close()
}

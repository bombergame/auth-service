package mysql

import (
	"database/sql"
	"fmt"
	"github.com/bombergame/auth-service/config"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	str string
	db  *sql.DB
}

func NewConnection() *Connection {
	return &Connection{
		str: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.StorageUser, config.StoragePassword,
			config.StorageHost, config.StoragePort, config.StorageName,
		),
	}
}

func (c *Connection) Open() error {
	var err error

	c.db, err = sql.Open("mysql", c.str)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Close() error {
	return c.db.Close()
}

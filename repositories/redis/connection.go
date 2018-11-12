package redis

import (
	"github.com/bombergame/auth-service/config"
	"github.com/go-redis/redis"
)

type Connection struct {
	client *redis.Client
}

func NewConnection() *Connection {
	c := redis.NewClient(
		&redis.Options{
			Addr:     config.SessionsStorageAddress,
			Password: config.SessionsStoragePassword,
		},
	)
	return &Connection{
		client: c,
	}
}

func (conn *Connection) Open() error {
	return nil
}

func (conn *Connection) Close() error {
	return conn.client.Close()
}

package profilesgrpc

import (
	"github.com/bombergame/common/logs"
)

type Client struct {
	config *Config
}

type Config struct {
	Logger *logs.Logger
}

func NewClient(c *Config) *Client {
	return nil //TODO
}

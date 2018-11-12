package profilesgrpc

import (
	"context"
	"github.com/bombergame/auth-service/config"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/logs"
	"google.golang.org/grpc"
	"strings"
)

type Client struct {
	config *Config
	conn   *grpc.ClientConn
	client ProfilesServiceClient
}

type Config struct {
	Logger *logs.Logger
}

func NewClient(c *Config) *Client {
	return &Client{
		config: c,
	}
}

func (c *Client) Connect() error {
	var err error
	c.conn, err = grpc.Dial(config.ProfilesServiceGrpcAddress, grpc.WithInsecure())
	if err != nil {
		return wrapError(err)
	}

	c.client = NewProfilesServiceClient(c.conn)

	c.config.Logger.Info("profiles-service grpc connection: " + config.ProfilesServiceGrpcAddress)
	return nil
}

func (c *Client) Disconnect() error {
	c.config.Logger.Info("profiles-service grpc connection shutdown")
	return c.conn.Close()
}

func (c *Client) GetProfileIDByCredentials(cr *Credentials) (*ProfileID, error) {
	p, err := c.client.GetProfileIDByCredentials(context.TODO(), cr)
	if err != nil {
		return nil, wrapError(err)
	}
	return p, nil
}

func wrapError(err error) error {
	msg := err.Error()
	if strings.Contains(msg, "not found") {
		return errs.NewNotFoundError("profile not found")
	}
	return errs.NewServiceError(err)
}

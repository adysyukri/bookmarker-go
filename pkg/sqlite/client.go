package sqlite

import (
	"context"
	"database/sql"
)

type Client struct {
	svc *sql.DB
}

func NewRepo(db *sql.DB) *Client {
	return &Client{db}
}

func (c *Client) Close(ctx context.Context) error {
	return c.svc.Close()
}

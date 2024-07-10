package sqlite

import (
	"context"
	"database/sql"
)

type Client struct {
	svc *sql.DB
}

func NewClient(db *sql.DB) *Client {
	return &Client{db}
}

func (c *Client) Close(ctx context.Context) error {
	return c.svc.Close()
}

func (c *Client) Add(ctx context.Context, query string, data ...any) error {
	_, err := c.svc.ExecContext(ctx, query, data...)
	return err
}

func (c *Client) Get(ctx context.Context, query string, data ...any) (*sql.Rows, error) {
	rows, err := c.svc.QueryContext(ctx, query, data...)
	return rows, err
}

func (c *Client) Delete(ctx context.Context, query string, data ...any) error {
	_, err := c.svc.ExecContext(ctx, query, data...)
	return err
}

func (c *Client) First(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.svc.QueryRowContext(ctx, query, args...)
}

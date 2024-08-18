package cockroachdb

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	db *pgxpool.Pool
}

func NewClient(ctx context.Context, cfg string) (*Client, error) {
	db, err := pgxpool.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Client{db}, nil

}

func (c *Client) Close() {
	c.db.Close()
}

func (c *Client) QueryRow(ctx context.Context, query string, data any, args ...any) error {
	return c.db.QueryRow(ctx, query, args...).Scan(data)
}

func (c *Client) Query(ctx context.Context, query string, data any, args ...any) error {
	return pgxscan.Select(ctx, c.db, data, query, args...)
}

func (c *Client) Exec(ctx context.Context, query string, args ...any) error {
	_, err := c.db.Exec(ctx, query, args...)
	return err
}

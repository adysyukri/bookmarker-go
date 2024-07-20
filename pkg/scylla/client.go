package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

type Client struct {
	session *gocqlx.Session
}

func NewClient(session *gocqlx.Session) *Client {
	return &Client{session: session}
}

func (c *Client) Insert(ctx context.Context, qb qb.Builder, data ...any) error {
	// q := c.session.Query(table.Insert()).Bind(data...).WithContext(ctx)
	err := c.session.Query(qb.ToCql()).WithContext(ctx).Bind(data...).ExecRelease()
	// err := q.ExecRelease()
	return err
}

func (c *Client) Select(ctx context.Context, qb qb.Builder, data ...any) (*gocqlx.Iterx, error) {
	q := c.session.Query(qb.ToCql()).WithContext(ctx)
	iter := q.Iter()
	err := q.ExecRelease()
	return iter, err
}

func (c *Client) Delete(ctx context.Context, qb qb.Builder, data ...any) error {
	err := c.session.Query(qb.ToCql()).WithContext(ctx).Bind(data...).ExecRelease()
	return err
}

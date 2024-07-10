package scylla

import (
	"context"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type Client struct {
	session *gocqlx.Session
}

func NewClient(session *gocqlx.Session) *Client {
	return &Client{session: session}
}

func (c *Client) Add(ctx context.Context, qb qb.Builder, data ...any) error {
	// stmt, names := qb.ToCql()
	q := c.session.Query(qb.ToCql()).Bind(data...).WithContext(ctx)
	err := q.ExecRelease()
	return err
}

func (c *Client) Get(ctx context.Context, qb qb.Builder, data ...any) error {
	// stmt, names := table.SelectAll()
	// q := c.session.ContextQuery(ctx, stmt, names)
	// err := q.SelectRelease(data)
	return nil
}

func (c *Client) Delete(ctx context.Context, table table.Metadata) error {
	return nil
}

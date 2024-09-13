package scylla

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
)

type Client struct {
	session *gocqlx.Session
}

func NewClient(session *gocqlx.Session) *Client {
	return &Client{session}
}

func (c *Client) QueryExec(ctx context.Context, qb qb.Builder, data interface{}) error {
	err := c.session.Query(qb.ToCql()).WithContext(ctx).BindStruct(data).Consistency(gocql.All).ExecRelease()
	return err
}

func (c *Client) QueryRow(ctx context.Context, qb qb.Builder, data ...any) (*gocqlx.Iterx, error) {
	q := c.session.Query(qb.ToCql()).WithContext(ctx)
	iter := q.Iter()
	err := q.ExecRelease()
	return iter, err
}

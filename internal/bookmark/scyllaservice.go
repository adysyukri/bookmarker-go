package bookmark

import (
	"context"

	"github.com/a-h/templ"
	"github.com/adysyukri/bookemarker-go/pkg/scylla"
	"github.com/scylladb/gocqlx/v2"
)

type scyllaservice struct {
	db scylla.Client
}

type ScyllaService interface {
	Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error)
	Get(ctx context.Context, session gocqlx.Session) (templ.Component, error)
	Delete(ctx context.Context, id string) error
}

func NewScyllaService(db scylla.Client) ScyllaService {
	return &scyllaservice{db}
}

func (s *scyllaservice) Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error) {
	//INSERT INTO table (column1, column2, ...) VALUES (value1, value2, ...)
	// builder := qb.Insert(BookmarkTableName).Columns(bookmarkMetadata.Columns...)
	builder := bookmarkTable.InsertBuilder()
	bm := NewBookMark(bp)
	data := []any{
		bm.ID,
		bm.Title,
		bm.Author,
		bm.Total,
		bm.Read,
		bm.CreatedAt,
	}
	err := s.db.Add(ctx, builder, data...)
	if err != nil {
		return nil, err
	}

	return BookmarkCard(bm), nil
}

func (s *scyllaservice) Get(ctx context.Context, session gocqlx.Session) (templ.Component, error) {
	//SELECT column1, column2, ... FROM table

	var bml BookmarkList
	q := session.Query(bookmarkTable.SelectAll()).WithContext(ctx)

	if err := q.SelectRelease(&bml); err != nil {
		return nil, err
	}
	return Home(bml), nil
}

func (s *scyllaservice) Delete(ctx context.Context, id string) error {
	return nil
}

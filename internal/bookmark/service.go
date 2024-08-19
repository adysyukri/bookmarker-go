package bookmark

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
	"github.com/adysyukri/bookemarker-go/pkg/scylla"
	"github.com/scylladb/gocqlx/v3/qb"
	"github.com/scylladb/gocqlx/v3/table"
)

const (
	Keyspace = "book"
)

type service struct {
	db scylla.Client
}

type Service interface {
	Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error)
	Get(ctx context.Context) (templ.Component, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, bp *BookmarkParams) (templ.Component, error)
}

func NewService(db scylla.Client) Service {
	return &service{db}
}

var bookmarkTable = table.New(
	table.Metadata{
		Name:    fmt.Sprintf("%s.%s", Keyspace, BookmarkTableName),
		Columns: []string{"id", "title", "author", "total", "read", "created_at"},
		PartKey: []string{"id"},
		SortKey: []string{"created_at"},
	},
)

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error) {
	//INSERT INTO table (column1, column2, ...) VALUES (value1, value2, ...)
	q := qb.Insert(bookmarkTable.Name()).Columns(bookmarkTable.Metadata().Columns...)

	bm := NewBookMark(bp)
	data := []any{
		bm.ID,
		bm.Title,
		bm.Author,
		bm.Total,
		bm.Read,
		bm.CreatedAt,
	}
	err := s.db.Insert(ctx, q, data...)
	if err != nil {
		return nil, err
	}

	return BookmarkCard(bm), nil
}

func (s *service) Get(ctx context.Context) (templ.Component, error) {
	//SELECT column1, column2, ... FROM table

	var bml BookmarkList

	q := qb.Select(bookmarkTable.Name()).Columns(bookmarkTable.Metadata().Columns...)

	iter, err := s.db.Select(ctx, q)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	for {
		bm := new(Bookmark)

		if !iter.Scan(
			&bm.ID,
			&bm.Title,
			&bm.Author,
			&bm.Total,
			&bm.Read,
			&bm.CreatedAt,
		) {
			break
		}
		bml = append(bml, bm)
	}

	// if err := iter.Select(&bml); err != nil {
	// 	return nil, err
	// }

	return Home(bml), nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	// DELETE FROM book.bookmarks WHERE id = ?;
	q := qb.Delete(bookmarkTable.Name()).Where(qb.Eq("id"))
	return s.db.Delete(ctx, q, id)
}

func (s *service) Update(ctx context.Context, id string, bp *BookmarkParams) (templ.Component, error) {
	// UPDATE FROM book.bookmarks SET column1 = ?, column2 = ? WHERE id = ?;
	q := qb.Update(bookmarkTable.Name()).Set(bookmarkTable.Metadata().Columns...).Where(qb.Eq("id"))
	err := s.db.Update(ctx, q, id)
	return nil, err
}

package bookmark

import (
	"context"

	"github.com/a-h/templ"
	"github.com/adysyukri/bookemarker-go/pkg/scylla"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type scyllaservice struct {
	db scylla.Client
}

type ScyllaService interface {
	Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error)
	Get(ctx context.Context) (templ.Component, error)
	Delete(ctx context.Context, id string) error
}

func NewScyllaService(db scylla.Client) ScyllaService {
	return &scyllaservice{db}
}

var bookmarkTable = table.New(
	table.Metadata{
		Name:    BookmarkTableName,
		Columns: []string{"id", "title", "author", "total", "read", "created_at"},
		PartKey: []string{"id"},
		SortKey: []string{"created_at"},
	},
)

func (s *scyllaservice) Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error) {
	//INSERT INTO table (column1, column2, ...) VALUES (value1, value2, ...)
	q := qb.Insert(BookmarkTableName).Columns(bookmarkTable.Metadata().Columns...)

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

func (s *scyllaservice) Get(ctx context.Context) (templ.Component, error) {
	//SELECT column1, column2, ... FROM table

	var bml BookmarkList

	// stmt := `SELECT id, title, author, total, read, created_at FROM book.bookmarks;`
	q := qb.Select(BookmarkTableName).Columns(bookmarkTable.Metadata().Columns...)

	iter, err := s.db.Select(ctx, q)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	// iter := session.Query(bookmarkTable.SelectAll()).Iter()

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
		} else {
			bml = append(bml, bm)
		}
	}

	if err := iter.Select(&bml); err != nil {
		return nil, err
	}

	return Home(bml), nil
}

func (s *scyllaservice) Delete(ctx context.Context, id string) error {
	// DELETE FROM book.bookmarks WHERE id = ?;
	q := qb.Delete(BookmarkTableName).Where(qb.Eq("id"))
	return s.db.Delete(ctx, q, id)
}

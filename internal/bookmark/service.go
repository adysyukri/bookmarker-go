package bookmark

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
	"github.com/adysyukri/bookemarker-go/pkg/scylla"
	"github.com/scylladb/gocqlx/v3/qb"
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

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error) {
	//INSERT INTO table (column1, column2, ...) VALUES (value1, value2, ...)
	columns := []string{"id", "title", "author", "total", "read", "created_at"}
	q := qb.Insert(BookmarkTableName).Columns(columns...)

	bm := NewBookMark(bp)
	err := s.db.QueryExec(ctx, q, bm)
	if err != nil {
		return nil, err
	}

	stmt, names := q.ToCql()
	fmt.Printf("query: %v\nvalue: %+v\ndata: %v\n", stmt, names, bm)

	return BookmarkCard(bm), nil
}

func (s *service) Get(ctx context.Context) (templ.Component, error) {
	//SELECT column1, column2, ... FROM table
	var bml BookmarkList

	q := qb.Select(BookmarkTableName)

	iter, err := s.db.QueryRow(ctx, q)
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	// for {
	// 	bm := new(Bookmark)

	// 	if !iter.Scan(
	// 		&bm.ID,
	// 		&bm.Title,
	// 		&bm.Author,
	// 		&bm.Total,
	// 		&bm.Read,
	// 		&bm.CreatedAt,
	// 	) {
	// 		break
	// 	}
	// 	bml = append(bml, bm)
	// }

	if err := iter.Select(&bml); err != nil {
		return nil, err
	}
	fmt.Println(q.ToCql())

	return Home(bml), nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	// DELETE FROM book.bookmarks WHERE id = ?;
	q := qb.Delete(BookmarkTableName).Where(qb.Eq("id"))
	fmt.Println(q.ToCql())
	return s.db.QueryExec(ctx, q, Bookmark{ID: id})
}

func (s *service) Update(ctx context.Context, id string, bp *BookmarkParams) (templ.Component, error) {
	// UPDATE FROM book.bookmarks SET column1 = ?, column2 = ? WHERE id = ?;
	q := qb.Update(BookmarkTableName).Set("title", "author", "total", "read", "created_at").Where(qb.Eq("id"))
	bm := UpdateBookMark(bp, id)
	stmt, names := q.ToCql()
	fmt.Printf("query: %v\nvalue: %+v\ndata: %v\n", stmt, names, bm)

	err := s.db.QueryExec(ctx, q, bm)
	if err != nil {
		return nil, err
	}
	return BookmarkCard(bm), err
}

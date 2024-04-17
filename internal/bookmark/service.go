package bookmark

import (
	"context"
	"fmt"

	"github.com/a-h/templ"
	"github.com/adysyukri/bookemarker-go/pkg/sqlite"
)

type service struct {
	db sqlite.Client
}

type Service interface{}

func NewService(db sqlite.Client) Service {
	return &service{db}
}

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (templ.Component, error) {
	q := fmt.Sprintf(
		"INSERT INTO %s (id, title, author, total, read, created_at) VALUES (?, ?, ?, ?);",
		BookmarkTableName,
	)

	bm := NewBookMark(bp)
	data := []any{
		bm.ID,
		bm.Title,
		bm.Author,
		bm.Total,
		bm.Read,
		bm.CreatedAt,
	}

	err := s.db.Add(ctx, q, data...)
	if err != nil {
		return nil, err
	}

	return BookmarkCard(bm), nil
}

func (s *service) Get(ctx context.Context) (templ.Component, error) {
	q := fmt.Sprintf(
		"SELECT id, title, author, total, read, created_at FROM %s;",
		BookmarkTableName,
	)

	rows, err := s.db.Get(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bml BookmarkList

	for rows.Next() {
		bm := new(Bookmark)

		err := rows.Scan(
			&bm.ID,
			&bm.Title,
			&bm.Author,
			&bm.Total,
			&bm.Read,
			&bm.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		bml = append(bml, bm)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Home(bml), nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id = ?;",
		BookmarkTableName,
	)

	return s.db.Delete(ctx, q, id)
}

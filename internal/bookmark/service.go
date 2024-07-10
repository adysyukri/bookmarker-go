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

type Service interface {
	Add(ctx context.Context, bp *BookmarkParams) (templ.Component, templ.Component, error)
	Get(ctx context.Context) (templ.Component, error)
	Delete(ctx context.Context, id string) error
}

func NewService(db sqlite.Client) Service {
	return &service{db}
}

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (templ.Component, templ.Component, error) {
	q := fmt.Sprintf(
		"INSERT INTO %s (id, title, author, total, read, created_at) VALUES (?, ?, ?, ?, ?, ?);",
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
		return nil, nil, err
	}

	count, err := getCount(ctx, s.db)
	if err != nil {
		return nil, nil, err
	}

	return BookmarkCard(bm), BookmarkCounter(count, templ.Attributes{"hx-swap-oob": true}), nil
}

func getCount(ctx context.Context, db sqlite.Client) (int, error) {
	q := fmt.Sprintf("SELECT COUNT(*) FROM %s;", BookmarkTableName)
	row := db.First(ctx, q)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
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

	count, err := getCount(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return Home(bml, count), nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id = ?;",
		BookmarkTableName,
	)

	return s.db.Delete(ctx, q, id)
}

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

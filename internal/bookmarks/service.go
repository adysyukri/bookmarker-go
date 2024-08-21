package bookmarks

import (
	"context"
	"fmt"

	"github.com/adysyukri/bookemarker-go/pkg/cockroachdb"
)

type service struct {
	db cockroachdb.Client
}

type Service interface{}

func NewService(db cockroachdb.Client) Service {
	return &service{db}
}

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (*Bookmark, error) {
	q := fmt.Sprintf(
		"INSERT INTO %s (id, title, author, total, read, created_at) VALUES ($1, $2, $3, $4, $5, $6);",
		BookmarkTableName,
	)

	bm := NewBookMark(bp)

	err := s.db.Exec(ctx, q,
		bm.ID,
		bm.Title,
		bm.Author,
		bm.Total,
		bm.Read,
		bm.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return bm, nil
}

func (s *service) Get(ctx context.Context) (BookmarkList, error) {
	q := fmt.Sprintf(
		"SELECT id, title, author, total, read, created_at FROM %s;",
		BookmarkTableName,
	)

	var bml BookmarkList

	err := s.db.Query(ctx, q, &bml)

	return bml, err
}

func (s *service) Delete(ctx context.Context, id string) error {
	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;",
		BookmarkTableName,
	)

	return s.db.Exec(ctx, q, id)
}

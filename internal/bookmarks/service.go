package bookmarks

import (
	"context"
	"fmt"

	"github.com/adysyukri/bookemarker-go/pkg/cockroachdb"
)

type service struct {
	db cockroachdb.Client
}

type servicer interface {
	Add(ctx context.Context, bp *BookmarkParams) (*Bookmark, int, error)
	Update(ctx context.Context, bp *BookmarkParams) (*Bookmark, error)
	Get(ctx context.Context) (BookmarkList, int, error)
	Delete(ctx context.Context, id string) (int, error)
}

func newService(db cockroachdb.Client) servicer {
	return &service{db}
}

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (*Bookmark, int, error) {
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
		return nil, 0, err
	}

	counter, err := s.counter(ctx)
	if err != nil {
		return nil, 0, err
	}

	return bm, counter, nil
}

func (s *service) counter(ctx context.Context) (int, error) {
	q := fmt.Sprintf(
		"SELECT count(*) from %s;",
		BookmarkTableName,
	)

	var counter int

	if err := s.db.QueryRow(ctx, q, &counter); err != nil {
		return 0, err
	}

	return counter, nil
}

func (s *service) Update(ctx context.Context, bp *BookmarkParams) (*Bookmark, error) {
	q := fmt.Sprintf(
		"UPDATE %s SET title = $1, author = $2, total = $3, read = $4 WHERE id = $5;",
		BookmarkTableName,
	)

	bm := MapBookmark(bp, nil)

	err := s.db.Exec(ctx, q,
		bm.Title,
		bm.Author,
		bm.Total,
		bm.Read,
		bm.ID,
	)
	if err != nil {
		return nil, err
	}

	return bm, nil
}

func (s *service) Get(ctx context.Context) (BookmarkList, int, error) {
	q := fmt.Sprintf(
		"SELECT id, title, author, total, read, created_at FROM %s;",
		BookmarkTableName,
	)

	var bml BookmarkList

	if err := s.db.Query(ctx, q, &bml); err != nil {
		return nil, 0, err
	}

	counter, err := s.counter(ctx)
	if err != nil {
		return nil, 0, err
	}

	return bml, counter, err
}

func (s *service) Delete(ctx context.Context, id string) (int, error) {
	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;",
		BookmarkTableName,
	)

	if err := s.db.Exec(ctx, q, id); err != nil {
		return 0, err
	}

	counter, err := s.counter(ctx)
	if err != nil {
		return 0, err
	}

	return counter, nil
}

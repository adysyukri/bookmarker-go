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
	Add(ctx context.Context, bp *BookmarkParams) (BookmarkList, error)
	Update(ctx context.Context, bp *BookmarkParams) (*Bookmark, error)
	Counter(ctx context.Context) (int, error)
	Get(ctx context.Context) (BookmarkList, int, error)
	Delete(ctx context.Context, bp *BookmarkParams) (BookmarkList, error)
}

func newService(db cockroachdb.Client) servicer {
	return &service{db}
}

var bml BookmarkList

func GetBookmarkList() BookmarkList {
	return bml
}

func (s *service) Add(ctx context.Context, bp *BookmarkParams) (BookmarkList, error) {
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

	bml = append(bml, bm)

	return bml, nil
}

func (s *service) Counter(ctx context.Context) (int, error) {
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
		"SELECT id, title, author, total, read FROM %s;",
		BookmarkTableName,
	)

	// var bml BookmarkList

	err := s.db.Query(ctx, q, &bml)
	if err != nil {
		return nil, 0, err
	}

	counter := len(bml)

	return bml, counter, err
}

func (s *service) Delete(ctx context.Context, bp *BookmarkParams) (BookmarkList, error) {
	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;",
		BookmarkTableName,
	)

	deletedbm := MapBookmark(bp, nil)

	err := s.db.Exec(ctx, q, deletedbm.ID)
	if err != nil {
		return nil, err
	}

	var filteredBookmark BookmarkList

	for _, bm := range bml {
		if bm.ID != deletedbm.ID {
			filteredBookmark = append(filteredBookmark, bm)
			fmt.Printf("filteredBookmark: %#v\n", filteredBookmark)
		}
	}

	bml = filteredBookmark

	return bml, nil
}

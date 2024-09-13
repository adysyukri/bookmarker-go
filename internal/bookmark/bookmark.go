package bookmark

import (
	"time"

	"github.com/segmentio/ksuid"
)

const (
	BookmarkTableName = "bookmarks"
)

// table: bookmarks
type Bookmark struct {
	ID     string `json:"id,omitempty" db:"id,omitempty" cql:"id,omitempty"`
	Title  string `json:"title,omitempty" db:"title,omitempty" cql:"title,omitempty"`
	Author string `json:"author,omitempty" db:"author,omitempty" cql:"author,omitempty"`
	Total  int    `json:"total,omitempty" db:"total,omitempty" cql:"total,omitempty"`
	Read   int    `json:"read,omitempty" db:"read,omitempty" cql:"read,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at,omitempty" cql:"created_at,omitempty"`
}

func NewBookMark(bp *BookmarkParams) *Bookmark {
	ksuid := ksuid.New()

	return &Bookmark{
		ID:        ksuid.String(),
		Title:     bp.Title,
		Author:    bp.Author,
		Total:     bp.Total,
		Read:      bp.Read,
		CreatedAt: time.Now(),
	}
}

func UpdateBookMark(bp *BookmarkParams, id string) *Bookmark {
	return &Bookmark{
		ID:        id,
		Title:     bp.Title,
		Author:    bp.Author,
		Total:     bp.Total,
		Read:      bp.Read,
		CreatedAt: time.Now(),
	}
}

type BookmarkList []*Bookmark

type BookmarkParams struct {
	Title  string `json:"title,omitempty" db:"title,omitempty" cql:"title,omitempty"`
	Author string `json:"author,omitempty" db:"author,omitempty" cql:"author,omitempty"`
	Total  int    `json:"total,omitempty" db:"total,omitempty" cql:"total,omitempty"`
	Read   int    `json:"read,omitempty" db:"read,omitempty" cql:"read,omitempty"`
}

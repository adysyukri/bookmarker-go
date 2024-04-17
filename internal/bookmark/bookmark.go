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
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Total  int    `json:"total,omitempty"`
	Read   int    `json:"read,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
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

type BookmarkList []*Bookmark

type BookmarkParams struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Total  int    `json:"total,omitempty"`
	Read   int    `json:"read,omitempty"`
}

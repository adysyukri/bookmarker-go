package bookmarks

import (
	"time"

	"github.com/segmentio/ksuid"
)

const (
	BookmarkTableName = "bookmarks"
)

// ++ views.templ variables

var inputDefault = `{
	title: "",
	author: "",
	total: null,
	read: null
}`

var resetDefault = `() => {
	modalOpen = false
	title = ""
	author = ""
	total = null
	read = null
}`

// -- views.templ variables

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

	createdAt := time.Now()

	bp.ID = ksuid.String()

	return MapBookmark(bp, &createdAt)
}

func MapBookmark(bp *BookmarkParams, createdAt *time.Time) *Bookmark {
	bm := &Bookmark{
		ID:     bp.ID,
		Title:  bp.Title,
		Author: bp.Author,
		Total:  bp.Total,
		Read:   bp.Read,
	}

	if createdAt != nil {
		bm.CreatedAt = *createdAt
	}

	return bm
}

type BookmarkList []*Bookmark

type BookmarkParams struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Total  int    `json:"total,omitempty"`
	Read   int    `json:"read,omitempty"`
}

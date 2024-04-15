package bookmark

import "time"

// table: bookmarks
type Bookmark struct {
	ID     string
	Title  string
	Author string
	Total  int
	Read   int

	CreatedAt time.Time
}

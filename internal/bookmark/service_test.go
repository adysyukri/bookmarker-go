package bookmark_test

import (
	"context"
	"testing"

	"github.com/adysyukri/bookemarker-go/internal/bookmark"
)

func TestAdd(t *testing.T) {
	_, _, err := svc.Add(context.Background(), &bookmark.BookmarkParams{
		Title:  "Title1",
		Author: "Auhtor1",
		Total:  20,
		Read:   1,
	})

	if err != nil {
		t.Errorf("error add: %s", err)
	}
}

func TestDelete(t *testing.T) {
	err := svc.Delete(context.Background(), "2fDPNDvpVx3cBQ2pvRR3TWaMSmf")

	if err != nil {
		t.Errorf("error add: %s", err)
	}
}

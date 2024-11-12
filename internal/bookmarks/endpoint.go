package bookmarks

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adysyukri/bookemarker-go/pkg/cockroachdb"
)

type endpoint struct {
	svc servicer
}

type API interface {
	// Page endpoints
	Home(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	// API endpoints
	// ...
}

func NewAPI(db cockroachdb.Client) API {
	return &endpoint{
		svc: newService(db),
	}
}

// GET /home (page)
func (e *endpoint) Home(w http.ResponseWriter, r *http.Request) {
	res, err := e.svc.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	Page(res).Render(w)
}

// POST /add (component)
func (e *endpoint) Add(w http.ResponseWriter, r *http.Request) {
	total, err := strconv.Atoi(r.FormValue("total"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	read, err := strconv.Atoi(r.FormValue("read"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	bp := &BookmarkParams{
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
		Total:  total,
		Read:   read,
	}

	res, err := e.svc.Add(r.Context(), bp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	GBookmarkCard(res).Render(w)
}

// POST /edit (component)
func (e *endpoint) Edit(w http.ResponseWriter, r *http.Request) {
	total, err := strconv.Atoi(r.FormValue("total"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	read, err := strconv.Atoi(r.FormValue("read"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	bp := &BookmarkParams{
		ID:     r.PathValue("id"),
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
		Total:  total,
		Read:   read,
	}

	res, err := e.svc.Update(r.Context(), bp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	GBookmarkCard(res).Render(w)
}

// DELETE /delete/{id}
func (e *endpoint) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := e.svc.Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occurs: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

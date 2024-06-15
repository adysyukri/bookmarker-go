package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adysyukri/bookemarker-go/internal/bookmark"
	"github.com/adysyukri/bookemarker-go/internal/order"
	"github.com/adysyukri/bookemarker-go/pkg/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

var db, _ = sql.Open("sqlite3", "./tmp/db.db")

func main() {
	defer db.Close()
	err := InitTable()
	if err != nil {
		log.Fatalln("error init table")
	}
	dbClient := sqlite.NewClient(db)
	svc := bookmark.NewService(*dbClient)
	orderSvc := order.NewService(*dbClient)

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		t, err := svc.Get(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error occurs: %s", err)
			return
		}

		t.Render(r.Context(), w)
	})

	http.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
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

		bp := &bookmark.BookmarkParams{
			Title:  r.FormValue("title"),
			Author: r.FormValue("author"),
			Total:  total,
			Read:   read,
		}

		t, err := svc.Add(r.Context(), bp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error occurs: %s", err)
			return
		}

		t.Render(r.Context(), w)
	})

	http.HandleFunc("DELETE /delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		// w.Write([]byte(id))
		err := svc.Delete(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error occurs: %s", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("GET /order", func(w http.ResponseWriter, r *http.Request) {
		t, err := orderSvc.Page(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error occurs: %s", err)
			return
		}

		t.Render(r.Context(), w)
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/"))))

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func InitTable() error {
	//	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS bookmarks (
		id TEXT PRIMARY KEY,
		title TEXT,
		author TEXT,
		total INTEGER,
		read INTEGER,
		created_at DATETIME
	);`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	fmt.Println("Table created successfully")

	return nil
}

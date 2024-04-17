package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db, _ = sql.Open("sqlite3", "./tmp/db.db")

func main() {
	defer db.Close()
	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		//page.Home().Render(r.Context(), w)
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/"))))

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func InitTable() error {
	defer db.Close()

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

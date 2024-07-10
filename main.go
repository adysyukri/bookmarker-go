package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adysyukri/bookemarker-go/internal/bookmark"
	"github.com/adysyukri/bookemarker-go/pkg/scylla"
	"github.com/gocql/gocql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/scylladb/gocqlx/v2"
)

// var db, _ = sql.Open("sqlite3", "./tmp/db.db")

const keyspace string = "book"

var s, _ = gocqlx.WrapSession(NewSession("localhost"))

func main() {
	// defer db.Close()

	// dbClient := sqlite.NewClient(db)
	// svc := bookmark.NewService(*dbClient)

	defer s.Close()

	s.ExecStmt(fmt.Sprintf(`DROP KEYSPACE IF EXISTS %s;`, keyspace))
	s.ExecStmt(fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'replication_factor' : 1};`, keyspace))

	err := InitTable()
	if err != nil {
		log.Fatalln("error init table: ", err)
	}

	dbClient := scylla.NewClient(&s)
	svc := bookmark.NewScyllaService(*dbClient)

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		t, err := svc.Get(r.Context(), s)
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

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/"))))

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func InitTable() error {
	//	defer db.Close()

	cqlStmt := `
	CREATE TABLE IF NOT EXISTS book.bookmarks (
		id text,
		title text,
		author text,
		total int,
		read int,
		created_at timestamp,
		PRIMARY KEY(id, created_at))
	WITH CLUSTERING ORDER BY (created_at DESC);`

	err := s.ExecStmt(cqlStmt)
	if err != nil {
		return err
	}

	fmt.Println("Table created successfully")

	return nil
}

// func CreateKeyspace(keyspace string) error {
// 	s.ExecStmt(fmt.Sprintf(`DROP KEYSPACE IF EXISTS %#v;`, keyspace))
// 	s.ExecStmt(fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %#v WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'replication_factor' : 1};`, keyspace))
// 	return nil
// }

func NewSession(hosts ...string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln("error create session: ", err)
	}
	return session, nil
}

package bookmark_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/adysyukri/bookemarker-go/internal/bookmark"
	"github.com/adysyukri/bookemarker-go/pkg/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

var dbClient *sqlite.Client
var svc bookmark.Service

func TestMain(m *testing.M) {
	db, _ := sql.Open("sqlite3", "../../tmp/db_test.db")
	defer db.Close()

	dbClient = sqlite.NewClient(db)
	svc = bookmark.NewService(*dbClient)

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
		log.Fatalln("error create table")
	}

	os.Exit(m.Run())
}

package bookmark_test

import (
	"log"
	"os"
	"testing"

	"github.com/adysyukri/bookemarker-go/internal/bookmark"
	"github.com/adysyukri/bookemarker-go/pkg/scylla"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/gocqlxtest"
)

var dbClient *scylla.Client
var svc bookmark.Service
var session gocqlx.Session

func TestMain(m *testing.M) {
	cluster := gocqlxtest.CreateCluster()
	session, _ = gocqlx.WrapSession(cluster.CreateSession())
	defer session.Close()

	dbClient = scylla.NewClient(&session)
	svc = bookmark.NewService(*dbClient)

	cqlStmt := `
	CREATE TABLE IF NOT EXISTS testing.bookmarkstest (
		id text,
		title text,
		author text,
		total int,
		read int,
		created_at timestamp,
		PRIMARY KEY(id, created_at))
	WITH CLUSTERING ORDER BY (created_at DESC);`

	if err := gocqlxtest.CreateKeyspace(cluster, "testing"); err != nil {
		log.Fatalln("error create keyspace: ", err)
	}

	if err := session.ExecStmt(cqlStmt); err != nil {
		log.Fatalln("error create table: ", err)
	}

	os.Exit(m.Run())
}

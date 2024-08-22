package main

import (
	"context"
	"log"

	"github.com/adysyukri/bookemarker-go/pkg/cockroachdb"
)

func main() {
	dbUri := "postgresql://root@localhost:26257/bookmark_go?application_name=bookmark_go&sslmode=disable"
	dbClient, err := cockroachdb.NewClient(context.Background(), dbUri)
	if err != nil {
		log.Fatalln("error init client db", err)
	}

	routes(*dbClient)
}

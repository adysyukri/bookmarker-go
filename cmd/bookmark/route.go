package main

import (
	"fmt"
	"net/http"

	"github.com/adysyukri/bookemarker-go/internal/bookmarks"
	"github.com/adysyukri/bookemarker-go/pkg/cockroachdb"
)

func routes(dbClient cockroachdb.Client) {
	bookmarkEp := bookmarks.NewAPI(dbClient)

	http.HandleFunc("GET /home", bookmarkEp.Home)
	http.HandleFunc("POST /add", bookmarkEp.Add)
	http.HandleFunc("PUT /edit/{id}", bookmarkEp.Edit)
	http.HandleFunc("DELETE /delete/{id}", bookmarkEp.Delete)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/"))))

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

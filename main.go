package main

import (
	"fmt"
	"net/http"

	"github.com/adysyukri/bookemarker-go/page"
)

func main() {
	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		page.Home().Render(r.Context(), w)
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/"))))

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

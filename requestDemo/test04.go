package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		id := query["id"]
		log.Println(id)

		name := query.Get("name")
		log.Println(name)
	})

	http.ListenAndServe("localhost:8080", nil)
}

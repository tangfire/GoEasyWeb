package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.Fragment)
	})

	server.ListenAndServe()
}

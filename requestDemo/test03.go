package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	server.ListenAndServe()
}

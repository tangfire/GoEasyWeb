package main

import (
	"fmt"
	"net/http"
)

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/writeheader", writeHeaderExample)

	server.ListenAndServe()

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {

		// table03.html
		//r.ParseMultipartForm(1024)

		//fmt.Fprintln(w, r.MultipartForm)
		fmt.Fprintln(w, r.FormValue("input1"))
	})

	server.ListenAndServe()
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	// 对应table03.html中的name
	file, _, err := r.FormFile("fileUpload")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/process", process)

	server.ListenAndServe()
}

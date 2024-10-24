package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	post := &Post{
		User:    "tangfire",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()

}

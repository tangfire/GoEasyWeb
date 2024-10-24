package main

import "net/http"

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/redirect", headerExample)

	server.ListenAndServe()

}

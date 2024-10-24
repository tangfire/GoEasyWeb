package main

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := "hello world welcome china"
	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/write", writeExample)

	server.ListenAndServe()

}

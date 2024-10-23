package main

import "net/http"

type helloHandler struct{}

type aboutHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {

	mh := helloHandler{}

	a := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	http.Handle("/hello", &mh)
	http.Handle("/about", &a)
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home!"))
	})

	// way01
	//http.HandleFunc("/welcome", welcome)

	//type HandlerFunc func(ResponseWriter, *Request)
	// ServeHTTP calls f(w, r).
	//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	//	f(w, r)
	//}

	//way02
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
}

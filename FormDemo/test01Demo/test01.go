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

		r.ParseForm()
		// 如果表单和URL里有同样的Key,那么它们会放在一个slice里：表单里的值靠前，URL的值靠后
		// table01.html
		//fmt.Fprintln(w, r.Form)
		// 如果只想要表单的key-value对，不要URL的，可以使用PostForm字段
		// table02.html
		//fmt.Fprintln(w, r.PostForm)

		fmt.Fprintln(w)
	})

	server.ListenAndServe()
}

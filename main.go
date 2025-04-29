package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Selamat Datang di Website saya!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static.", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":5000", nil)
}
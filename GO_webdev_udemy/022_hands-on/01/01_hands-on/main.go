package main

import (
	"io"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "root page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "!!!!eunkyoon kim!!!!")
}

func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", about)

	http.ListenAndServe(":8080", nil)
}

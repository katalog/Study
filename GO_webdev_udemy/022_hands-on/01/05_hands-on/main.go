package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(myName))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "root.gohtml", req.URL.Path)
}

func bar(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", req.URL.Path)
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", req.URL.Path)
}

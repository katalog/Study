package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//var reader io.Reader = req.Body
	//b, _ := ioutil.ReadAll(reader)
	//fmt.Println(string(b))

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(req.Form)

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

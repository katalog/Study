package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counterCookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "counterCookie",
			Value: "some value",
			Path:  "/",
		})

		fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
		fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")

		return
	}

	v, err := strconv.ParseInt(c.Value, 10, 32)
	v++
	c.Value = strconv.FormatInt(v, 10)
	http.SetCookie(w, c)
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "cookie val:", c.Value)
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counterCookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

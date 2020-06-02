package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type urlshortrequest struct {
	Request_url string `json:"request_url"`
}

var URLs map[string]string

func main() {
	router := httprouter.New()
	router.POST("/api/v1/new", createShortUrl)
	router.GET("/api/v1/:url", goToUrl)
	log.Fatal(http.ListenAndServe(":8000", router))
}

var autoinc int

func createShortUrl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var test urlshortrequest
	err := decoder.Decode(&test)
	if err != nil {
		panic(err)
	}

	autoinc = autoinc + 1

	str := "myurl" + strconv.Itoa(autoinc)

	fmt.Println(test.Request_url, " -> ", str)

	_, _ = io.WriteString(w, str)

	URLs = make(map[string]string)
	URLs[str] = test.Request_url
}

func goToUrl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	encodedurl := params.ByName("url")

	if gotourl, ok := URLs[encodedurl]; ok {

		http.Redirect(w, r, gotourl, 301)
	}
}

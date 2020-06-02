package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type httpCode struct {
	Descrip string `json:"Descrip`
	Code    int    `json:"Code`
}

type httpCodes []httpCode

func main() {
	rcvd := `[
		{
		  "Descrip": "StatusOK",
		  "Code": 200
		},
		{
		  "Descrip": "StatusMovedPermanently",
		  "Code": 301
		},
		{
		  "Descrip": "StatusFound",
		  "Code": 302
		},
		{
		  "Descrip": "StatusSeeOther",
		  "Code": 303
		},
		{
		  "Descrip": "StatusTemporaryRedirect ",
		  "Code": 307
		},
		{
		  "Descrip": "StatusBadRequest",
		  "Code": 400
		},
		{
		  "Descrip": "StatusUnauthorized",
		  "Code": 401
		},
		{
		  "Descrip": "StatusPaymentRequired",
		  "Code": 402
		},
		{
		  "Descrip": "StatusForbidden",
		  "Code": 403
		},
		{
		  "Descrip": "StatusNotFound",
		  "Code": 404
		},
		{
		  "Descrip": "StatusMethodNotAllowed",
		  "Code": 405
		},
		{
		  "Descrip": "StatusTeapot",
		  "Code": 418
		},
		{
		  "Descrip": "StatusInternalServerError",
		  "Code": 500
		}]`

	var data httpCodes
	bs := []byte(rcvd)
	err := json.Unmarshal(bs, &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Printf("%v\n", v)
	}
}

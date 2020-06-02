package main

import (
	"log"
	"os"
	"text/template"
)

type Regi struct {
	Southern bool
	Central  bool
	Northern bool
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  Regi
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := make([]Hotel, 1)
	hotels[0].Name = "cali"
	hotels[0].Address = "lmlm"
	hotels[0].City = "asdqq"
	hotels[0].Zip = 10101
	hotels[0].Region.Southern = true

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}

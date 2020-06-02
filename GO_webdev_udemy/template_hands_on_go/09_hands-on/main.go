package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

/// Date,Open,High,Low,Close,Volume,Adj Close

type csvdata struct {
	// Date     string
	// Open     float64
	// High     float64
	// Low      float64
	// Close    float64
	// Volume   int64
	// AdjClose float64
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	// read csv to struct
	csvfile, _ := os.Open("table.csv")
	r := csv.NewReader(csvfile)
	var dat []csvdata

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		dat = append(dat, csvdata{
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
			line[5],
			line[6],
		})
	}
	// struct to html
	err := tpl.Execute(os.Stdout, dat)
	if err != nil {
		log.Fatalln(err)
	}
}

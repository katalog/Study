package main

import (
	"log"
	"os"
	"text/template"
)

type Menu struct {
	Breakfast string
	Lunch     string
	Dinner    string
}

type Menus []Menu

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menus := Menus{
		Menu{
			Breakfast: "미국식",
			Lunch:     "중국식",
			Dinner:    "일식",
		},
		Menu{
			Breakfast: "한식",
			Lunch:     "한식",
			Dinner:    "한식",
		},
	}

	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}

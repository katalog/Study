package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	encd()
	decd()

	//mshl()
}

func encd() {
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(os.Stdout).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func decd() {
	const jsonStream = `
		{"Fname": "Ed", "Lname": "Bond", "Items": ["Suit", "Gun", "Wry sense of humor"]}
	`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m person
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}
}

func mshl() {
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(j)

	p2 := person{}

	json.Unmarshal(j, &p2)

	fmt.Println(p2)
}

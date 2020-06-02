package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Veniam nisi sunt cillum magna ex et sint exercitation laboris excepteur qui"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("I'm giving her all she's got Captain!", err)
	}

	n1 := len(s)
	n2 := len(s64)

	fmt.Println("orgsize : ", n1, "encsize : ", n2)

	fmt.Println(s)

	fmt.Println(s64)

	fmt.Println(string(bs))
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GoesToEleven/golang-web-dev/042_mongodb/07_solution/controllers"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/07_solution/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}

func loadSessionFromFile(uc *UserController) {
	var file, err = os.Open("session.dat")
	if err != nil {
		return
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	for {
		var u models.User
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
		uc.session[u.Id] = u
	}
}

func saveSessionToFile(uc *UserController) {
	var file, err = os.Create("session.dat")
	if err != nil {
		return
	}
	defer file.Close()

	dec := json.NewEncoder(file)

	for _, v := range uc.session {
		dec.Encode(&v)
	}
}

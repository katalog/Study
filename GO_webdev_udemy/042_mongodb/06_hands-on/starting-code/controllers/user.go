package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoesToEleven/golang-web-dev/042_mongodb/04_controllers/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type UserModel struct {
	DB map[string]User
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (um *UserModel) GetDB() *map[string]User {
	return &um.DB
}

type UserController struct {
	DB *UserModel
}

func NewUserController(um *UserModel) *UserController {
	return &UserController{um}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	oid := p.ByName("id")

	// composite literal
	//u := models.User{}

	// Fetch user

	u, err := uc.DB.GetDB()[oid]
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = string(uuid.Must(uuid.NewV4()))

	// store the user in mongodb
	//uc.session.DB("go-web-dev-db").C("users").Insert(u)
	uc.DB.GetDB()[u.ID] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	oid := p.ByName("id")

	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(404)
	// 	return
	// }

	// oid := bson.ObjectIdHex(id)

	// Delete user
	if _, ok := uc.DB.GetDB()[oid]; !ok {
		w.WriteHeader(404)
		return
	}
	delete(uc.DB.GetDB(), oid)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}

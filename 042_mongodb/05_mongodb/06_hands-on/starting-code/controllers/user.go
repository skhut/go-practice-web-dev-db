package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skhut/go-practice-web-dev-db/042_mongodb/05_mongodb/06_hands-on/starting-code/models"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	sessionMap *map[string]models.User
}

func NewUserController(s *map[string]models.User) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oid := bson.ObjectIdHex(id)

	// composite literal
	u := models.User{}

	u, ok := uc.sessionMap[string(oid)]

	if !ok {
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
	u.Id = bson.NewObjectId()

	// store the user in mongodb
	uc.sessionMap[string(u.Id)] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// composite literal
	u := models.User{}

	u, ok := uc.sessionMap[string(oid)]

	if !ok {
		w.WriteHeader(404)
		return
	}
	delete(uc.sessionMap, oid)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skhut/go-practice-web-dev-db/042_mongodb/05_mongodb/06_hands-on/starting-code/controllers"
	"github.com/skhut/go-practice-web-dev-db/042_mongodb/05_mongodb/06_hands-on/starting-code/models"
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

func getSession() *map[string]models.User {
	return &make(map[string]models.User)
}

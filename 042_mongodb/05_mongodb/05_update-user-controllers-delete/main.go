package main

import (
	"net/http"

	//"github.com/GoesToEleven/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"
	//"github.com/GoesToEleven/golang-web-dev/042_mongodb/07_solution/controllers"
	//"github.com/GoesToEleven/golang-web-dev/042_mongodb/07_solution/controllers"
	"github.com/GoesToEleven/golang-web-dev/042_mongodb/07_solution/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

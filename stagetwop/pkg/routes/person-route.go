package routes

import (
	"github.com/Zekeriyyah/hngx/stagetwop/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/api", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/{userId}", controllers.DeleteUser).Methods("DELETE")

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zekeriyyah/hngx/stagetwop/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running: ")
	log.Fatal(http.ListenAndServe("localhost:9840", r))
}

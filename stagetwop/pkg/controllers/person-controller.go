package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zekeriyyah/hngx/stagetwop/pkg/models"
	"github.com/Zekeriyyah/hngx/stagetwop/pkg/utils"
	"github.com/gorilla/mux"
)

var NewUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	newPersons := models.GetAllUsers()
	res, err := json.Marshal(newPersons)
	if err != nil {
		fmt.Println("Error while Marshaling!")
	}

	w.Header().Set("Content-Type:", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	userDetails, _ := models.GetUserById(ID)

	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type:", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	b := createUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	ID, _ := strconv.ParseInt(userId, 0, 0)

	deletedUser := models.DeleteUser(ID)

	res, _ := json.Marshal(deletedUser)
	w.Header().Set("Content-Type:", "application/json")
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	params := mux.Vars(r)
	userId := params["userId"]

	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error: Error while parsing")
	}

	userDetails, db := models.GetUserById(ID)
	if userDetails.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if userDetails.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if userDetails.Country != "" {
		userDetails.Country = updateUser.Country
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

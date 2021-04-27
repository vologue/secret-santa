package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	mysqlExtention "secretSanta/extentions/mysql"
	"secretSanta/utils"
	"time"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	userData, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var user mysqlExtention.Users
	err := json.Unmarshal(userData, &user)
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Bad request", nil)
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	mysqlResponse := mysqlExtention.CreateUser(&user)
	if mysqlResponse.Error != nil {
		log.Println(mysqlResponse.Error)
		response := "Create user failed for  " + user.Email
		utils.RespondWithError(w, http.StatusInternalServerError, response, nil)
	} else {
		log.Println("User created successfully!")
		utils.RespondWithError(w, http.StatusOK, "User", user)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userData, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var user mysqlExtention.Users
	err := json.Unmarshal(userData, &user)
	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Bad request", nil)
	}
	user.UpdatedAt = time.Now()
	sqlResponse := mysqlExtention.UpdateUser(user)
	if sqlResponse.Error != nil {
		log.Println(sqlResponse.Error)
		response := "update user failed on user " + fmt.Sprint(user.UserId)
		utils.RespondWithError(w, http.StatusInternalServerError, response, nil)
	} else {
		log.Println("Update user successfull!")
		utils.RespondWithError(w, http.StatusOK, "update successfull", nil)
	}
}

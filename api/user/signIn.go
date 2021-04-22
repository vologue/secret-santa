package users

import (
	"net/http"
	mysqlExtention "secretSanta/extentions/mysql"
	"secretSanta/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := mysqlExtention.Login{
		Username: "vologue",
		Password: "password",
	}
	mysqlExtention.LoginInsert(user)
	utils.RespondWithError(w, http.StatusOK, "test")
}

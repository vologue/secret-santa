package main

import (
	"fmt"
	"log"
	"net/http"
	"secretSanta/api/healthCheck"
	users "secretSanta/api/users"
	mysqlExtention "secretSanta/extentions/mysql"

	"github.com/gorilla/mux"
)

func startUp() {
	db := mysqlExtention.ConnectDB()
	db.AutoMigrate(&mysqlExtention.Users{})
	log.Println("Table created")
}

func main() {
	startUp()
	Router := mux.NewRouter()
	Router.HandleFunc("/api/v1/health-check", healthCheck.GetHealth).Methods("GET")
	Router.HandleFunc("/api/v1/create-user", users.NewUser).Methods("POST")
	Router.HandleFunc("/api/v1/update-user", users.UpdateUser).Methods("PUT")
	err := http.ListenAndServe(":8000", Router)
	if err != nil {
		fmt.Print(err)
	}

}

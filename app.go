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
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health-check", healthCheck.GetHealth).Methods("GET")
	router.HandleFunc("/api/v1/create-user", users.NewUser).Methods("POST")
	router.HandleFunc("/api/v1/update-user", users.UpdateUser).Methods("PUT")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}

}

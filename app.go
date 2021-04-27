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
	router.HandleFunc("/alive", healthCheck.GetHealth).Methods("GET")
	router.HandleFunc("/create", users.NewUser).Methods("POST")
	router.HandleFunc("/create", users.UpdateUser).Methods("PUT")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}

}

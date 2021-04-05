package main

import (
	"fmt"
	"net/http"
	"secretSanta/api/healthCheck"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/alive", healthCheck.GetHealth).Methods("GET")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}

}

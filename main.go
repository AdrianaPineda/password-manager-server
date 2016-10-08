package main

import (
	"fmt"
	database "github.com/AdrianaPineda/password-manager-server/database"
	"log"
	"net/http"
)

func main() {

	database, err := database.InitDB("user=adrianaPineda dbname = adrianaPineda sslmode=disable")

	router := NewRouter()

	fmt.Printf("API initialized\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}

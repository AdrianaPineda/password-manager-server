package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	fmt.Printf("API initialized\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}

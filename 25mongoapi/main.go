package main

import (
	"fmt"
	"github.com/Swaty-G/GolangTutorial/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", router.Router()))
	fmt.Println("Listening on port 4000...")
}

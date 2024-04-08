package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// https://go.dev/ref/mod -- Go Modules reference documentation link to understand the Go modules in detail and how to use them in Go
// https://pkg.go.dev/github.com/gorilla/mux -- Gorilla Mux package documentation link to understand the package and its functions in detail

// run -->  go get -u github.com/gorilla/mux  (to install the mux package in the Go project)
func main() { // main function to start the execution of the program in Go and run the web server to handle the incoming HTTP requests and route them to the appropriate handler functions based on the request URL and method using the mux package
	fmt.Println("Mod in Go")
	greeter()                                   // calling the greeter function to greet the users of the mod package
	r := mux.NewRouter()                        // creating a new router using the NewRouter function from the mux package to handle the incoming HTTP requests and route them to the appropriate handler functions based on the request URL and method
	r.HandleFunc("/", serveHome).Methods("GET") // registering a handler function serveHome for the home page URL "/" with the GET method using the HandleFunc method of the router r from the mux package

	log.Fatal(http.ListenAndServe(":8000", r)) // starting the HTTP server on port 8000 and passing the router r to the ListenAndServe function to handle the incoming HTTP requests and route them to the appropriate handler functions; log.Fatal is used to log any error returned by the ListenAndServe function and exit the program if there is an error; the server runs on the local machine or localhost and listens on port 8000; ListenAndServe function is from the net/http package in Go and it starts an HTTP server with the specified address and handler to handle the incoming HTTP requests and route them to the appropriate handler functions based on the request URL and method using the router r in this case from the mux package

}

func greeter() { // greeter function to greet the users of the mod package
	fmt.Println("Hey there mod users!")
}

func serveHome(w http.ResponseWriter, r *http.Request) { // serveHome function to serve the home page of the web application to the users; it takes the response writer and the request as input parameters and returns nothing as output;
	w.Write([]byte("<h1>Welcome to golang series on youtube</h1>")) // writes the welcome message to the response writer to display on the web page when the user visits the home page of the web application using the web browser and the URL http://localhost:8000/ in this case as the server is running on port 8000 on the local machine or localhost
}

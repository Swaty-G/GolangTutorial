package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com" // URL of the website whose content needs to be fetched

func main() {
	fmt.Println("Web requests in Go")

	response, err := http.Get(url) // makes a GET request to the URL and returns the response and any error in making the request
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body) // reads the content of the website using the io.ReadAll function from the io package and stores it in the databytes slice of bytes and any error in reading the content is stored in the err variable of type error in Go
	if err != nil {
		panic(err)
	}
	content := string(databytes)                               // converting the slice of bytes to a string
	fmt.Println("Content of the website is: ", content)        // prints the content of the website
	fmt.Println("Content of the website is: ", content[:1000]) // prints the first 1000 characters of the content of the website

}

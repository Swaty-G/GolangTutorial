package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://www.google.com:8080/search?q1=golang&q2=reactjs" // URL to be parsed using the url.Parse function from the net/url package

func main() {
	fmt.Println("handling URLs in Go")
	fmt.Println("URL is: ", myURL)

	result, err := url.Parse(myURL) // parses the URL and returns the result and any error in parsing the URL; the result is of type *url.URL; the result contains the Scheme, Host, Path, RawQuery, and other details of the URL; parsing the URL is useful to extract the different parts of the URL
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme: ", result.Scheme)     // prints the scheme of the URL which is https in this case
	fmt.Println("Host: ", result.Host)         // prints the host of the URL which is www.google.com:8080 in this case
	fmt.Println("Path: ", result.Path)         // prints the path of the URL which is /search in this case
	fmt.Println("Port: ", result.Port())       // prints the port of the URL which is 8080 in this case
	fmt.Println("RawQuery: ", result.RawQuery) // prints the raw query of the URL which is q=golang in this case

	qparams := result.Query()                                    // returns the query parameters of the URL as a map; the query parameters are q=golang in this case
	fmt.Printf("The type of query parameters is: %T\n", qparams) // prints url.Values as the type of the query parameters; url.Values is a map[string][]string type in Go where the key is a string and the value is a slice of strings as the query parameters can have multiple values for the same key in the URL query string
	fmt.Println(qparams["q2"])                                   // prints the value of the query parameter with key q which is [golang] in this case

	// iterating over the query parameters; order is not guaranteed
	for _, val := range qparams { // iterating over the query parameters map
		fmt.Println("Param is: ", val) // prints the query parameters of the URL which is [golang reactjs] in this case as the query parameters can have multiple values for the same key in the URL query string and the values are stored as a slice of strings in the url.Values map in Go
	}

	// creating a URL from parts of the URL using the url.URL struct from the net/url package
	partsOfUrl := &url.URL{ //pass reference to the URL struct to store the parts of the URL because the URL struct is a pointer type in Go
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/search",
		RawQuery: "q1=golang&q2=reactjs",
	}
	anotherURL := partsOfUrl.String() // returns the URL as a string using the String method of the url.URL struct because the URL struct is a pointer type in Go
	fmt.Println("Another URL is: ", anotherURL)
}

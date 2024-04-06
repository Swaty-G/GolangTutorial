package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs in Go")
	url := "http://localhost:8000/get" //"https://httpbin.org/get"
	PerformGetRequest(url)             // makes a GET request to the specified URL and prints the status code, content length, and content of the response body
	PerformPostJsonRequest()           // makes a POST request to the specified URL with a JSON payload in the request body and prints the content of the response body
	PerformPostFormRequest()           // makes a POST request to the specified URL with a form data in the request body and prints the content of the response body
}

// PerformGetRequest function makes a GET request to the specified URL and prints the status code, content length, and content of the response body
func PerformGetRequest(url string) {
	fmt.Println("Performing GET request to: ", url)

	response, err := http.Get(url) // makes a GET request to the URL and returns the response and any error in making the request
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // caller's responsibility to close the connection

	fmt.Println("Status code: ", response.StatusCode) // prints the status code of the response which is 200 for a successful request

	// Read the content of the response body using io.ReadAll function -- option 1
	/*	content, err := io.ReadAll(response.Body) // reads the content of the website using the io.ReadAll function from the io package and stores it in the databytes slice of bytes and any error in reading the content is stored in the err variable of type error in Go
		if err != nil {
			panic(err)
		}
		fmt.Println("Content1 length: ", response.ContentLength) // prints the content length of the response body which is -1 if the length is unknown or -1 if the body is nil
		fmt.Println("Content1: ", string(content))               // prints the content of the response body as a string
	*/

	// Read the content of the response body using strings.Builder -- option 2 (preferred) as it is more efficient for large content size and also provides more control over the content read operation
	var responseString strings.Builder         // strings.Builder is a struct in Go that is used to efficiently build strings using Write method  -- it is more efficient than using string concatenation using + operator
	content1, err := io.ReadAll(response.Body) // reads the content of the website using the io.ReadAll function from the io package and stores it in the databytes slice of bytes and any error in reading the content is stored in the err variable of type error in Go
	if err != nil {
		panic(err)
	}
	byteCount, _ := responseString.Write(content1)                        // Write method of strings.Builder struct writes the content of the response body to the strings.Builder struct and returns the number of bytes written and any error in writing the content
	fmt.Println("Content2 length of the Get request: ", byteCount)        // prints the number of bytes written to the strings.Builder struct
	fmt.Println("Content2 of the Get request: ", responseString.String()) // prints the content of the response body as a string using the String method of the strings.Builder struct

}

// PerformPostJsonRequest function makes a POST request to the specified URL with a JSON payload in the request body and prints the content of the response body
func PerformPostJsonRequest() {
	const myURL string = "http://localhost:8000/post" //"https://httpbin.org/post" // URL to be used for the POST request

	// fake json payload to be sent in the POST request body
	requestBody := strings.NewReader(`{"name":"swaty", "job":"developer"}`) // creating a fake JSON payload to be sent in the POST request body using the strings.NewReader function from the strings package

	response, err := http.Post(myURL, "application/json", requestBody) // makes a POST request to the URL with the JSON payload in the request body and the content type as application/json
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()               // caller's responsibility to close the connection
	content, err := io.ReadAll(response.Body) // reads the content of the response body to avoid connection leaks
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of the POST request is: ", string(content)) // prints the content of the response body of the POST request
}

// PerformPostFormRequest function makes a POST request to the specified URL with a form data in the request body and prints the content of the response body
func PerformPostFormRequest() {
	const myURL string = "http://localhost:8000/postform" //URL to be used for the POST form request

	//form data to be sent in the POST request body
	data := url.Values{} // creating a url.Values map to store the form data to be sent in the POST request body
	data.Add("name", "swaty")
	data.Add("job", "developer")
	data.Add("location", "USA")

	response, err := http.PostForm(myURL, data) //makes a POST request to the URL with the form data in the request body
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()                                           //caller's responsibility to close the connection
	content, err := io.ReadAll(response.Body)                             //reads the content of the response body to avoid connection leaks
	fmt.Println("Content of the POST form request is: ", string(content)) //prints the content of the response body of the POST form request
}

/* difference between post and postform is that postform is used to send form data in the request body, while post is used to send any type of data in the request body like JSON, XML, etc. and the content type needs to be specified in the post request while it is not required in the postform request as it is set to application/x-www-form-urlencoded by default in the postform request in Go  */

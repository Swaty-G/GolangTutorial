package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup // sync package is used to synchronize goroutines; WaitGroup is a struct that waits for a collection of goroutines to finish; it is used to wait for the completion of all goroutines launched from the main function

func main() {
	fmt.Println("Goroutines in Go")

	// invoking greeter function using goroutines using time.Sleep function to pause the execution of the current goroutine for 2 milliseconds
	/*	go greeter("Hello") // go keyword is used to create a goroutine; a goroutine is a lightweight thread managed by the Go runtime; it is a function that runs concurrently with other functions; means it won't wait for the function to finish before moving on to the next line of code (non-blocking)
		greeter("World")*/

	// invoke the getStatusCode function for each website in the list
	websiteList := []string{"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com", "https://www.linkedin.com", "https://badwebsite.com"}
	for _, website := range websiteList {
		go getStatusCode(website)
		waitGroup.Add(1) // Add method increments the WaitGroup counter by one; it is used to add the number of goroutines to wait for
	}
	waitGroup.Wait() //blocks main from finishing until all goroutines are done executing (blocking) // Wait method blocks until the WaitGroup counter is zero; it is used to wait for the completion of all goroutines launched from the main function
}

// greeter function prints the string 5 times with a 2 millisecond delay between each print statement using time.Sleep function
func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond) // time.Sleep is used to pause the execution of the current goroutine for a specified duration; time.Millisecond is a duration of 1 millisecond
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer waitGroup.Done() // Done method decrements the WaitGroup counter by one; it is used to signal the completion of a goroutine
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("OOPS! Error occurred for %s with error: %s", endpoint, err)
	} else {
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint) // %d is a placeholder for an integer; %s is a placeholder for a string
	}
}

/* This code snippet demonstrates the use of goroutines in Go. A goroutine is a lightweight thread managed by the Go runtime. It is a function that runs concurrently with other functions. The go keyword is used to create a goroutine. It is a non-blocking operation, meaning it won't wait for the function to finish before moving on to the next line of code.

waitGroup is a struct that waits for a collection of goroutines to finish. It is used to wait for the completion of all goroutines launched from the main function. The Add method increments the WaitGroup counter by one, and the Wait method blocks until the WaitGroup counter is zero, waiting for the completion of all goroutines. The Done method decrements the WaitGroup counter by one, signaling the completion of a goroutine.

*/

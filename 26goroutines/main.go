package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup // pointer to a WaitGroup struct; WaitGroup is a struct that waits for a collection of goroutines to finish; it is used to wait for the completion of all goroutines launched from the main function
// sync package is used to synchronize goroutines;

var signals = []string{"test"} // test is a placeholder for a string; signals is a slice of strings

var mutex sync.Mutex // pointer to a Mutex struct; Mutex is a mutual exclusion lock; sync.Mutex is a struct that provides locking mechanism to synchronize access to shared resources

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
	fmt.Println("Signals:", signals)
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
		fmt.Printf("OOPS! Error occurred for %s with error: %s\n", endpoint, err)
	} else {
		mutex.Lock() // Lock method locks the mutex; it is used to synchronize access to shared resources; it is used to prevent
		signals = append(signals, endpoint)
		mutex.Unlock() // Unlock method unlocks the mutex; it is used to unlock the mutex

		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint) // %d is a placeholder for an integer; %s is a placeholder for a string
		fmt.Println("Signals:", signals)
	}
}

/* This code snippet demonstrates the use of goroutines in Go. A goroutine is a lightweight thread managed by the Go runtime. It is a function that runs concurrently with other functions. The go keyword is used to create a goroutine. It is a non-blocking operation, meaning it won't wait for the function to finish before moving on to the next line of code.

waitGroup is a struct that waits for a collection of goroutines to finish. It is used to wait for the completion of all goroutines launched from the main function. The Add method increments the WaitGroup counter by one, and the Wait method blocks until the WaitGroup counter is zero, waiting for the completion of all goroutines. The Done method decrements the WaitGroup counter by one, signaling the completion of a goroutine.

 Here's how the Go program you provided executes:

1. **Start of the main function:** The program begins execution with the `main` function, where it first prints "Goroutines in Go" to the console.

2. **Setting up goroutines for status checks:** The program then creates a slice called `websiteList` containing various website URLs. It iterates over this slice, and for each website, it launches a goroutine to execute the `getStatusCode` function.

   - **Concurrency with goroutines:** The `go` keyword before `getStatusCode(website)` starts a new goroutine, which allows the `getStatusCode` function to run concurrently with the main function and any other goroutines.

   - **Tracking goroutines with WaitGroup:** Before launching each goroutine, the program calls `waitGroup.Add(1)`, signaling the `WaitGroup` that there's one more goroutine to wait for. This ensures the `main` function waits for all goroutines to finish before exiting.

3. **Wait for goroutines to finish:** After launching all the goroutines, the program calls `waitGroup.Wait()`. This call blocks the main function until all goroutines have signaled completion by each calling `waitGroup.Done()`.

4. **Execution of getStatusCode:** Inside each `getStatusCode` goroutine, the program:

   - Makes an HTTP GET request to the given website endpoint.
   - If there's an error, it prints an error message to the console.
   - If the GET request is successful, it locks the mutex to ensure exclusive access to the `signals` slice, appends the website endpoint to the `signals` slice, and then unlocks the mutex. This prevents race conditions where multiple goroutines might try to write to the `signals` slice simultaneously.
   - Prints the HTTP status code for the website to the console.
   - Calls `defer waitGroup.Done()` at the start of the function, which will be executed as the last step before the function returns, indicating to the `WaitGroup` that the goroutine's work is done.

5. **Main function ends:** After all goroutines have finished and `waitGroup.Wait()` has returned, the program prints the contents of the `signals` slice to the console.

6. **Program exits:** With no further code to execute, the program ends.

This program demonstrates how goroutines can be used to perform concurrent tasks, such as making HTTP requests to multiple websites simultaneously. The use of `WaitGroup` ensures that the main function waits for all goroutines to complete before exiting. The `mutex` is used to synchronize access to shared resources

 Here’s a detailed walkthrough of the program’s flow, highlighting the key points where execution passes through different parts of the code:

1. **Start of the main function:**
   - The program starts and prints "Goroutines in Go" to the console.

2. **Initialization of `signals`:**
   - A slice called `signals` is initialized with one element, "test".

3. **Setup of `websiteList` and launching of goroutines :**
   - A slice called `websiteList` is initialized with a list of website URLs.
   - A for loop iterates over each website URL.
     - Inside the loop, for each website:
       - A goroutine is launched with `go getStatusCode(website)` to fetch the status code for that website. Execution in this goroutine begins immediately.
       - The `waitGroup` counter is incremented by 1 with `waitGroup.Add(1)` to track this newly started goroutine.

4. **`waitGroup.Wait()` call:**
   - The `main` function blocks at this line until all the goroutines that have had `waitGroup.Add(1)` called for them have also called `waitGroup.Done()`.

5. **Concurrent execution of `getStatusCode`:**
   - Each goroutine executes `getStatusCode` for its assigned website.
   - The function attempts to make an HTTP GET request to the endpoint.
     - If an error occurs, it's printed to the console.
     - If the GET request is successful, the following steps are taken within the goroutine:
       - The mutex is locked.
       - The website is appended to the `signals` slice.
       - The mutex is unlocked.
       - The HTTP status code and endpoint are printed to the console.
   - Regardless of success or failure, `waitGroup.Done()` is deferred to be called at the end of the function, decrementing the `waitGroup` counter.

6. **End of goroutines:**
   - As each goroutine completes, it calls `waitGroup.Done()`, potentially unblocking the `waitGroup.Wait()` call in the `main` function if it's the last outstanding goroutine.

7. **Printing final `signals` content:**
   - Once all goroutines have finished (i.e., the `waitGroup` counter is 0), the `main` function resumes execution after the `waitGroup.Wait()` call.
   - The contents of the `signals` slice, which now includes the endpoints that have been processed, are printed to the console.

8. **End of the main function:**
   - The program has no more statements to execute in the `main` function and thus terminates.

Please note that the specific sequence of execution for the goroutines is nondeterministic; the Go runtime schedules them, and they may run in any order. Also, the exact lines I've referred to match the structure of the code you posted, assuming it is structured in a typical way without additional code or comments in between.
*/

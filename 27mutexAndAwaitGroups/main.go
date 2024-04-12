package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions in Go")

	waitGroup := &sync.WaitGroup{} // pointer to a WaitGroup struct; WaitGroup is a struct that waits for a collection of goroutines to finish; it is used to wait for the completion of all goroutines launched from the main function; & is used to get the memory address of the WaitGroup struct variable

	mutex := &sync.RWMutex{} // pointer to a Mutex struct; Mutex is a mutual exclusion lock; sync.Mutex is a struct that provides locking mechanism to synchronize access to shared resources; & is used to get the memory address of the Mutex struct variable;whenever a goroutine wants to access a shared resource, it must lock the mutex; when the goroutine is done with the shared resource, it must unlock the mutex

	var score = []int{0} // score is a slice of integers; 0 is a placeholder for an integer value

	waitGroup.Add(3)                                      // 3 is the number of goroutines to wait for; Add method increments the WaitGroup counter by one; it is used to add the number of goroutines to wait for
	go func(waitGroup *sync.WaitGroup, m *sync.RWMutex) { // go keyword is used to create a goroutine; a goroutine is a lightweight thread managed by the Go runtime; it is a function that runs concurrently with other functions; means it won't wait for the function to finish before moving on to the next line of code (non-blocking)
		fmt.Println("One Run")
		mutex.Lock() // Lock method locks the mutex; it is used to synchronize access to shared resources; it is used to prevent
		score = append(score, 1)
		mutex.Unlock()   // Unlock method unlocks the mutex; it is used to unlock the mutex
		waitGroup.Done() // Done method decrements the WaitGroup counter by one; it is used to signal the completion of a goroutine
	}(waitGroup, mutex)
	go func(waitGroup *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two Runs")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		waitGroup.Done()
	}(waitGroup, mutex)
	// waitGroup.Add(1) // add after each go routine or add 3 at once in the beginning
	go func(waitGroup *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three Runs")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()

		mutex.RLock() // RLock method locks the mutex for reading; it is used to synchronize access to shared resources; it is used to prevent
		fmt.Println("Score:", score)
		mutex.RUnlock() // RUnlock method unlocks the mutex for reading; it is used to unlock the mutex
		waitGroup.Done()
	}(waitGroup, mutex)

	waitGroup.Wait() //blocks main from finishing until all goroutines are done executing (blocking) // Wait method blocks until the WaitGroup counter is zero; it is used to wait for the completion of all goroutines launched from the main function
	fmt.Println("Score:", score)
}

/* Order of execution:
1. main function is invoked
2. WaitGroup variable waitGroup is declared and initialized with the memory address of a WaitGroup struct
3. score variable is declared and initialized with a slice of integers containing 0
4. Add method is called on waitGroup to increment the WaitGroup counter by 3
5. Three goroutines are created using the go keyword
6. Each goroutine appends an integer to the score slice

Output: Order is not guaranteed due to the concurrent nature of goroutines; the order of execution may vary each time the program is run due to the concurrent nature of goroutines
One Run
Two Runs
Three Runs
Score: [0 1 2 3]
*/

/*
go run --race .
This will run the program and check for race conditions in the code.
returns exit status 66 if there are race conditions

if u see race condition in the output, you can use Mutex to fix

Mutex is a mutual exclusion lock; sync.Mutex is a struct that provides locking mechanism to synchronize access to shared resources

Mutex.RLock() is used to lock the mutex for reading and Mutex.RUnlock() is used to unlock the mutex for reading
you should use Mutex.RLock() when you want to read the shared resource and Mutex.Lock() when you want to write to the shared resource to prevent other goroutines from reading or writing to the shared resource at the same time
*/

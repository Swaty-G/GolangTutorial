package main

import (
	"fmt"
	"sync"
)

/*Channels are a way in which multiple Go routines talk to each other. They still don't know what's happening in another go routine or how long it takes to run another go routine*/
func main() {
	fmt.Println("Channels in Go")

	// channel is a pipe that connects two goroutines; it is used to send and receive data between goroutines; it is a typed conduit through which you can send and receive values with the channel operator, <-; channels are used to synchronize goroutines; channels are used to communicate between goroutines

	myChannel := make(chan int, 1) // make is a built-in function that creates a channel; chan is a keyword used to create a channel; int is the type of data that the channel will hold; myChannel is a channel that holds integers; 1 is the buffer size of the channel; buffer size is the number of elements that the channel can hold; if the buffer size is 0, the channel is unbuffered; if the buffer size is greater than 0, the channel is buffered; a buffered channel can hold up to the buffer size elements without blocking; if the buffer is full, the sender will block until the receiver receives a value from the channel

	/*myChannel <- 42          // send 42 to the channel myChannel; <- is the channel operator; it is used to send and receive data from a channel; it is used to send data to a channel
	fmt.Println(<-myChannel) // receive data from the channel myChannel and print it to the console; <- is the channel operator
	*/

	waitGroup := &sync.WaitGroup{} // to resolve deadlock
	waitGroup.Add(2)

	//send ONLY
	go func(myChannel chan<- int, waitGroup *sync.WaitGroup) {
		fmt.Println("Sending 42")
		//close(myChannel)
		myChannel <- 42
		myChannel <- 43

		//close(myChannel) // close the channel after sending data; close is a built-in function that closes the channel; it is used to close the channel after sending data; it is used to signal that no more data will be sent to the channel; it is used to signal that the channel is done sending data; no one can write to the channel after it is closed
		waitGroup.Done()
	}(myChannel, waitGroup)

	//receive ONLY
	go func(myChannel <-chan int, waitGroup *sync.WaitGroup) {
		fmt.Println("Receiving data")

		//val, isChannelOpen := <-myChannel // receive data from the channel myChannel; <- is the channel operator; isChannelOpen is a boolean variable that checks if the channel is open; if the channel is open, isChannelOpen is true; if the channel is closed, isChannelOpen is false
		//if isChannelOpen {
		//	fmt.Println("Channel is open")
		//	fmt.Println(val)
		//} else {
		//	fmt.Println("Channel is closed")
		//}
		//waitGroup.Done()

		fmt.Println(<-myChannel) //listen to the channel for 1st value 42 or use make(chan int, 1) to buffer the channel
		fmt.Println(<-myChannel) //listen to the channel for 2nd value 43
		waitGroup.Done()
	}(myChannel, waitGroup)

	waitGroup.Wait()
}

/* Output:
Channels in Go
fatal error: all goroutines are asleep - deadlock!

this means that the main goroutine is asleep and there are no other goroutines to wake it up; this is a deadlock; a deadlock occurs
 deadlock happened because there is no one listening to the channel, so the main goroutine is waiting for someone to send data to the channel, but there is no one to send data to the channel

solution: add a channel to listen to the channel

if u have 2 values to send, then 2 values should be received

add a buffer to the channel to prevent the deadlock; a buffered channel can hold up to the buffer size elements without blocking; if the buffer is full, the sender will block until the receiver receives a value from the channel

*/

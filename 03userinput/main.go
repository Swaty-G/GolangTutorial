package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome) // prints --> Welcome to user input

	reader := bufio.NewReader(os.Stdin) // bufio is a package that provides buffered I/O.; os is a package that provides a platform-independent interface to operating system functionality and it has a Stdin variable that is a file object that represents the standard input stream; bufio.NewReader() is a function that takes an io.Reader object and returns a new bufio.Reader object that wraps the io.Reader object and provides buffering and some help for textual I/O. in Go and it is used to read input from the user in the console or terminal and it returns a pointer to the bufio.
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok syntax || error handling
	input, _ := reader.ReadString('\n')       // ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter. _ is used to ignore the error returned by the function as we are not handling it here, input is the variable that stores the value entered by the user
	fmt.Println("Thanks for rating, ", input) // prints --> Thanks for rating,  5
	fmt.Printf("Type of rating is %T", input) // prints --> Type of rating is string
}

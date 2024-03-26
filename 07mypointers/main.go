package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in Go")

	/*	var ptr *int // declaring a pointer variable, ptr is the variable name, *int is the type of the variable, *int is a pointer to an integer, it can store the memory address of an integer variable

		fmt.Println("Value of pointer is: ", ptr) // prints --> Value of ptr is:  <nil>, nil is a special value in Go that represents a zero value for pointers, interfaces, maps, slices, channels, and functions, it is the zero value for pointers, it is the default value for pointers when they are not initialized
	*/

	myNumber := 23

	var ptr2 = &myNumber                                        // ptr2 is a pointer to an integer, it stores the memory address of the myNumber variable
	fmt.Println("memory address of actual pointer  is: ", ptr2) // prints --> Value of actual pointer is:  0xc0000140b0, it is the memory address of the myNumber variable
	fmt.Println("Value of actual pointer is: ", *ptr2)          // prints --> Value of actual pointer is:  23, *ptr2 is the value stored at the memory address stored in the ptr2 variable, it is called de-referencing the pointer, it returns the value stored at the memory address stored in the pointer variable

	// so ptr2 is the reference to the memory address of the variable myNumber and *ptr2 is the value stored at the memory address stored in the ptr variable

	*ptr2 = *ptr2 + 10
	fmt.Println("new value of ptr2 is: ", *ptr2) // prints --> new value of ptr2 is:  33, it adds 10 to the value stored at the memory address stored in the ptr2 variable, which means operation is performed on the actual value stored at the memory address stored in the ptr2 variable and not on any copies of the value
}

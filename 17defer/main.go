package main

import "fmt"

// defer is used to delay the execution of a statement until the end of the function; printed in reverse order of defer statements like LIFO (Last In First Out)
func main() {
	fmt.Println("Defer in Go")
	defer fmt.Println("World") // defer is used to delay the execution of a statement until the end of the function
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello ") // prints --> Hello

	myDefer() // calling myDefer function that has defer statements inside it and will print in reverse order of defer statements like LIFO (Last In First Out) --> 4 3 2 1 0
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// defer statements are executed in LIFO (Last In First Out) order -->
// World One Two
// 0 1 2 3 4

/* Output:
Defer in Go
Hello
4
3
2
1
0
Two
One
World
*/

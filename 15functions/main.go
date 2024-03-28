package main

import "fmt"

func main() { // no calling of the function is required in Go, the main function is automatically called when the program is run
	fmt.Println("Functions in Go")
	greeter() // calling the greeter function

	/*func greeterTwo(){ // function cannot be declared inside another function in Go and will throw an error if done so
		fmt.Println("Namaste from Go")
	}*/

	result := adder(3, 5) // calling the adder function and storing the result in the result variable
	fmt.Println("Result is: ", result)

	proResult := proAdder(1, 2, 3, 4, 5) // calling the proAdder function and storing the result in the proResult variable
	fmt.Println("Pro result is: ", proResult)
}

func adder(valOne int, valTwo int) int { // the return type of the function is declared after the parameter list; the return type is int in this case
	return valOne + valTwo
}

func proAdder(values ...int) int { // the ... before the type name of the last parameter indicates that it takes zero or more arguments of that type; values is a slice of int in this case
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {
	fmt.Println("Namaste from Go")
}

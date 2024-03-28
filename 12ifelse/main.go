package main

import "fmt"

func main() {
	fmt.Println("if else in Go")

	loginCount := 23
	var result string

	// eg1: using if else
	if loginCount < 10 { // if the condition is true, the code inside the if block will be executed
		result = "Regular User"
	} else if loginCount > 10 { // if the condition in the if block is false, the code inside the else if block will be executed
		result = "Watch out"
	} else {
		result = "Exactly 10 logins"
	}
	fmt.Println(result)

	// eg2: direct number check without any variable
	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	// eg3: using if else with initialization; using ; to separate the initialization and the condition
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

}

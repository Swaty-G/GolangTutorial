package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days of the week: ", days)

	// for loop can be used to iterate over a slice
	//eg1
	for d := 0; d < len(days); d++ { //d is the index of the element in the days slice, len(days) returns the length of the days slice
		fmt.Printf("The day is %v\n", days[d])
	}

	//eg2
	for i := range days { // range is a keyword in Go that is used to iterate over elements in a variety of data structures, here we are using it to iterate over the elements of the days slice; it returns the index of the element in the slice
		fmt.Println(days[i])
	}

	//eg3
	for index, day := range days { // this is similar to the above for loop, but here we are using two variables index and day to get the index and the value of the element in the days slice
		fmt.Printf("At index %v, the day is %v\n", index, day) //prints --> At index 0, the day is Sunday, At index 1, the day is Monday, ..., At index 6, the day is Saturday ; %v is used to print the value of the variable in the placeholder position
	}

	//eg4
	for _, day := range days { // this is similar to the above for loop, but here we are using _ to ignore the index and only get the value of the element in the days slice
		fmt.Printf("the day is %v\n", day) //prints --> the day is Sunday, the day is Monday, ..., the day is Saturday
	}

	// for loop with condition
	//eg5
	rougueValue := 1
	for rougueValue < 10 { // for loop with condition
		if rougueValue == 2 {
			goto lco // goto is a keyword in Go that is used to jump to a specific label in the code, here we are using it to jump to the lco label
		}
		if rougueValue == 5 {
			break // break is a keyword in Go that is used to exit the loop, here we are using it to exit the loop when the value of rougueValue is 5
		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Print("jumping at LCO ")

}

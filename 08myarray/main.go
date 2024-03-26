package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList) // prints --> Fruit list is:  [Apple Tomato  Peach], it prints the values of the array, there is an empty value at index 2 as it is not assigned any value and its default value is an empty string for string type arrays

	fmt.Println("Fruit list length is: ", len(fruitList)) // prints --> Fruit list length is:  4, even though there are only 3 values in the array, the length of the array is 4 as it is the size of the array

	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg list is: ", len(vegList))

	var fruitList2 [3]int
	fruitList2[1] = 2
	fmt.Println("Fruit list 2 is: ", fruitList2) // prints --> Fruit list 2 is:  [0 2 0], it prints the values of the array, there are empty values at index 0 and index 2 as they are not assigned any value and their default value is 0 for integer type arrays
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go")

	var fruitList = []string{"Apple", "Tomato", "Peach", "Orange"} // []string is a slice of strings, it is a collection of strings, it is a reference to an underlying array, it is a dynamic array, it can grow or shrink in size, it is a reference type in Go, it is a reference to an underlying array, it is a dynamic array, it can grow or shrink in size, it is a reference type in Go
	fmt.Printf("Type of fruitList is: %T\n", fruitList)
	fmt.Println("Fruit list is: ", fruitList) // prints --> Fruit list is:  [Apple Tomato Peach Orange], it prints the values of the slice

	fruitList = append(fruitList, "Banana", "Mango")               // append() is a function that adds new elements to the slice
	fmt.Println("Fruit list after adding new fruits: ", fruitList) // prints --> Fruit list after adding a new fruit is:  [Apple Tomato Peach Orange Banana Mango], it prints the values of the slice after adding a new fruit

	fruitList = append(fruitList[1:])                                 // it starts from index 1 and removes index 1
	fmt.Println("Fruit list after removing first fruit: ", fruitList) // prints --> Fruit list after removing first fruit is:  [Tomato Peach Orange Banana Mango], it prints the values of the slice after removing the first fruit

	fruitList = append(fruitList[1:4])                                               // it starts from index 1 and removes the elements up to index 3 but not including index 3
	fmt.Println("Fruit list after removing elements from index 1 to 3: ", fruitList) // prints --> [Peach Orange Banana] it prints the values of the slice after removing elements from index 1 to 3

	fruitList = append(fruitList[:2])                                                      // it starts from the default index 0 and removes the elements up to index 2 but not including index 2
	fmt.Println("Fruit list after removing elements from default index to 1: ", fruitList) // prints --> [Peach Orange] it prints the values of the slice after removing elements from the default index to 1

	// slices with make() function
	highScores := make([]int, 4) // make() is a built-in function that creates a slice of a specified type, length, and capacity, it creates a slice of integers with a length and capacity of 4
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777                          // it will throw an error as the index is out of range, the length of the slice is 4 and the index is 4, so it is out of range
	highScores = append(highScores, 777)         // it will add the value 777 to the slice
	fmt.Println("High scores are: ", highScores) // prints --> High scores are:  [234 945 465 867], it prints the values of the slice

	// sorting slices
	fmt.Println(sort.IntsAreSorted(highScores))         // prints --> false, it prints false as the slice is not sorted in increasing order
	sort.Ints(highScores)                               // sort.Ints() is a function that sorts the slice of integers in increasing order
	fmt.Println("Sorted high scores are: ", highScores) // prints --> Sorted high scores are:  [234 465 777 867], it prints the values of the slice after sorting
	fmt.Println(sort.IntsAreSorted(highScores))         // prints --> true, it prints true as the slice is sorted in increasing order

	// removing elements from slices
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby", "golang"}
	fmt.Println("Courses: ", courses) // prints --> Courses:  [reactjs javascript swift python ruby golang], it prints the values of the slice
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)  // it removes the element at index 2
	fmt.Println("Courses after removing index 2: ", courses) // prints --> Courses after removing index 2:  [reactjs javascript python ruby golang], it prints the values of the slice after removing the element at index 2; ... is used to unpack the slice and pass it as an argument to the append() function; it is called variadic arguments; courses[:index] returns the elements up to index 2 but not including index 2; courses[index+1:] returns the elements from index 3 to the end of the slice

}

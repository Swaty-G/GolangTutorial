package main

import "fmt"

func main() {
	fmt.Println("Welcome to my maps in Go")

	// creating a map in Go
	languages := make(map[string]string) // make() is a built-in function in Go that is used to create a map, slice, or channel, it returns an initialized map, slice, or channel, it takes the type of the map, slice, or channel as an argument, here we are creating a map with string keys and string values
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages) // prints --> List of all languages:  map[JS:Javascript RB:Ruby PY:Python], it prints the key-value pairs of the map
	fmt.Println("JS shorts for: ", languages["JS"])   // prints --> JS shorts for:  Javascript, it prints the value of the key JS; languages["JS"] is used to access the value of the key JS

	// deleting an entry from the map
	delete(languages, "RB")                                               // delete() is a built-in function in Go that is used to delete an entry from the map, it takes the map and the key as arguments, here we are deleting the key RB from the languages map
	fmt.Println("List of all languages after deleting Ruby: ", languages) // prints --> List of all languages after deleting Ruby:  map[JS:Javascript PY:Python], it prints the key-value pairs of the map after deleting the key RB

	// iterating over a map
	for key, value := range languages { // range is a keyword in Go that is used to iterate over elements in a variety of data structures, here we are using it to iterate over the key-value pairs of the languages map
		fmt.Printf("For key %v, value is %v\n", key, value) // prints --> For key JS, value is Javascript, For key PY, value is Python, it prints the key-value pairs of the map
	}

	// option to ignore the key
	for _, value := range languages { // _ is used to ignore the key and only get the value
		fmt.Printf("For key v, value is %v\n", value) // prints --> For key v, value is Javascript, For key v, value is Python, it prints the values of the map
	}
}

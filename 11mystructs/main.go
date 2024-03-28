package main

import "fmt"

func main() {
	fmt.Println("Welcome to my structs in Go")
	//no inheritance in go; no super or parent concepts in go

	swaty := User{"Swaty", "swaty@go.dev", true, 25}                    //creating a user object using the User struct
	fmt.Println(swaty)                                                  //prints the struct with field values only; prints --> {Swaty swaty@go.dev true 25}
	fmt.Printf("Swaty details are: %+v\n", swaty)                       //prints the struct with field names and values; prints --> Swaty details are: {Name:Swaty Email:swaty@go.dev Status:true Age:25}
	fmt.Printf("Name is %v and email is %v\n", swaty.Name, swaty.Email) //prints specific fields of the struct; prints --> Name is Swaty and email is swaty@go.dev
}

type User struct { //defining a struct in Go	//struct is a composite data type in Go, it is used to group together different types of data, it is similar to a class in object-oriented programming languages; it is a collection of fields; is written in capital case
	Name   string
	Email  string
	Status bool
	Age    int
} // all the fields in the struct are called fields, they are like variables in the struct and are in capital case, they are exported and can be used outside the package

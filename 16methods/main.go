package main

import "fmt"

// Methods in Go
// Methods are functions that belong to a type and have a receiver argument that is the type of the method and is placed between the func keyword and the method name in the method definition syntax in Go as shown below:
// func (receiver receiverType) methodName(parameters) returnType { // method definition syntax in Go

// difference between a function and a method in Go is that a method is associated with a type and a function is not associated with any type
func main() {
	fmt.Print("methods in Go")

	swaty := User{"Swaty", "swaty@go.dev", true, 25}                    //creating a user object using the User struct
	fmt.Println(swaty)                                                  //prints the struct with field values only; prints --> {Swaty swaty@go.dev true 25}
	fmt.Printf("Swaty details are: %+v\n", swaty)                       //prints the struct with field names and values; prints --> Swaty details are: {Name:Swaty Email:swaty@go.dev Status:true Age:25}
	fmt.Printf("Name is %v and email is %v\n", swaty.Name, swaty.Email) //prints specific fields of the struct; prints --> Name is Swaty and email is swaty@go.dev
	swaty.GetStatus()                                                   //calling the GetStatus method on the swaty object
	swaty.NewEmail()                                                    //calling the NewEmail method on the swaty object
	fmt.Print("Swaty details are: ", swaty)                             //actual email of the user is not changed, prints --> Swaty details are: {Swaty swaty@go.dev true 25} as copy of the struct is passed to the method and not the actual struct itself so the changes are not reflected in the actual struct object swaty itself; if we want to change the actual struct object we need to pass the pointer to the struct object to the method

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
} // all the fields in the struct are called fields, they are like variables in the struct and are in capital case, they are exported and can be used outside the package

func (u User) GetStatus() { // defining a method in Go
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() { // defining a method in Go
	u.Email = "test@go.dev"
	fmt.Println("Email of this user : ", u.Email)
}

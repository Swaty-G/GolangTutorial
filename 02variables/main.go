package main

import "fmt"

const LoginToken string = "asjdhfjhsdfjhsdf" //constant variable, can't be changed later on in the code,
// capital L in the beginning of the variable name means it is exported and can be used outside the package like public access modifier in other languages

func main() {
	fmt.Println("variables") // prints --> variables

	var username string = "swaty"
	fmt.Println(username)                              // prints --> swaty
	fmt.Printf("Variable is of type: %T \n", username) // prints --> Variable is of type: string

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn) // prints --> Variable is of type: bool

	/*uint8       the set of all unsigned  8-bit integers (0 to 255)
	uint16      the set of all unsigned 16-bit integers (0 to 65535)
	uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	url: https://go.dev/ref/spec#Numeric_types
	*/
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal) // prints --> Variable is of type: uint8

	var smallFloat float64 = 255.455555545455 //float64 is the default type for floating point numbers, float32 is less precise and gives upto 5 decimal points
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat) // prints --> Variable is of type: float64

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)                              //do not print any garbage value but prints 0
	fmt.Printf("Variable is of type: %T \n", anotherVariable) // prints --> Variable is of type: int

	//implicit type
	var website = "https://swaty.com" // type is inferred from the value, so lexer does the job of inferring the type as it is declared, but you can't assign another type to it later eg: website = 5 will throw an error as it is inferred as string type, and you can't assign int to it later on in the code
	fmt.Println(website)              // prints --> https://swaty.com

	// no var style
	numberOfUser := 300000 // no need to use var keyword, but you can't use this outside a function, this is a shorthand way of declaring and initializing a variable, and you can change the value later
	fmt.Println(numberOfUser)

	numberOfUser = 300001     // as you can change the value later
	fmt.Println(numberOfUser) // prints --> 300001

	//print constant
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken) // prints --> Variable is of type: string
}

/*In summary, := is used for declaring and initializing a variable, while = is used for assigning a value to an already declared variable.
The var keyword is used for declaring a variable with a specific type, while the type can be inferred from the value when using :=. Constants are declared using the const keyword and cannot be changed later in the code.*/

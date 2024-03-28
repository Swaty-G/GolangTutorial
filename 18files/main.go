package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files in Go")

	content := "This needs to be written to the file\n"
	file, err := os.Create("./myFile.txt") // creates a file in the current directory with the name myFile.txt; if the file already exists, it will be truncated to zero length
	if err != nil {                        // if there is an error in creating the file, the error is printed
		panic(err) // panic is used to stop the program if there is an error in creating the file
	}

	length, err := io.WriteString(file, content) // writes the content to the file created above using the io.WriteString function from the io package
	/*if err != nil {                              // if there is an error in writing the content to the file, the error is printed
		panic(err) // panic is used to stop the program if there is an error in writing the content to the file
	}*/
	CheckNilErr(err)
	fmt.Println("Length of the content written to the file is: ", length) // prints the length of the content written to the file
	defer file.Close()                                                    // closes the file after writing the content to it using defer so that the file is closed at the end of the function execution

	readFile("./myFile.txt") // calling the readFile function to read the content of the file created above

}

func readFile(filename string) { // function to read the content of the file
	databyte, err := os.ReadFile(filename)                        // reads the content of the file using the os.ReadFile function from the os package and stores it in the databyte slice of bytes and any error in reading the file is stored in the err variable of type error in Go
	CheckNilErr(err)                                              // if there is an error in reading the content of the file, the error is printed
	fmt.Println("Data read from the file is: ", string(databyte)) // prints the content of the file as a string by converting the slice of bytes to a string using the string function in Go that converts a slice of bytes to a string in Go otherwise it will print the slice of bytes as it is in the output

}

// CheckNilErr is a function to check if the error is nil or not, if it is not nil, it panics and stops the program execution
func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

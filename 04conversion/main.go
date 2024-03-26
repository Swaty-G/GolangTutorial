package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our pizza: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // 64 is the bit size of the floating point number, 64 is the default size, strconv is the package that provides functions for converting strings to other types, and vice versa; strings is the package that provides functions for manipulating strings in Go; strings.TrimSpace removes leading and trailing white spaces from the string otherwise it will throw an error while converting the string to a number as it can't convert a string with leading or trailing white spaces to a number eg: " 5" or "5 " will throw an error while converting to a number as it has leading or trailing white spaces which are not allowed in a number conversion from a string to a number in Go;
	// strconv.ParseFloat converts a string to a floating point number and returns the number and an error if any occurred during the conversion
	if err != nil { // if there is an error while converting the string to a number, it will be handled here
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}

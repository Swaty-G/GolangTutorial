package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of go lang")

	presentTime := time.Now() // time.Now() is a function that returns the current date and time
	fmt.Println("The present time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 01-02-2006 15:04:05 Monday is the format of the date and time in Go, it is a reference date and time in Go, you can use this format to print the date and time in this format, it is a reference date and time in Go, you can use this format to print the date and time in this format

	createdDate := time.Date(2020, time.December, 11, 23, 0, 0, 0, time.UTC) // time.Date() is a function that creates a date and time with the given year, month, day, hour, minute, second, nanosecond, and location
	fmt.Println("Created date is: ", createdDate)                            // prints --> Created date is:  2020-12-10 23:00:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006 Monday"))                     // prints --> 12-10-2020 Thursday
}

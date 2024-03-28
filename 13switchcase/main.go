package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in Go")

	rand.Seed(time.Now().UnixNano()) // seeding the random number generator with the current time in nanoseconds, this is done to get a different random number every time the program is run
	diceNumber := rand.Intn(6) + 1   // generates a random number between 0 and 5, adding 1 to get a random number between 1 and 6
	fmt.Println("Dice number is: ", diceNumber)

	switch diceNumber {
	case 1: // if diceNumber is 1, this block will be executed
		fmt.Println("Dice value is 1 and you can open the game")
	case 2: // if diceNumber is 2, this block will be executed
		fmt.Println("Dice value is 2 and you can move 2 steps of the game")
	case 3: // if diceNumber is 3, this block will be executed
		fmt.Println("Dice value is 3 and you can move 3 steps of the game")
		fallthrough //fallthrough is used to execute the next case block even if the condition is not met for the next case block
	case 4: // if diceNumber is 4, this block will be executed
		fmt.Println("Dice value is 4 and you can move 4 steps of the game")
	case 5: // if diceNumber is 5, this block will be executed
		fmt.Println("Dice value is 5 and you can move 5 steps of the game")
	case 6: // if diceNumber is 6, this block will be executed
		fmt.Println("Dice value is 6 and you can move 6 steps of the game and roll the dice again")
	default: // if diceNumber is not between 1 and 6, this block will be executed
		fmt.Println("Invalid dice number")

	}
}

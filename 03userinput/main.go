package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Welcome message
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	// Create reader to take input from keyboard
	reader := bufio.NewReader(os.Stdin)

	// Ask user for rating
	fmt.Println("Enter the rating for our pizza:")

	// Read input until Enter key (\n)
	input, _ := reader.ReadString('\n')

	// Show result
	fmt.Println("Thanks for rating:", input)

	fmt.Printf("Type of this rating input is %T\n", input)
}
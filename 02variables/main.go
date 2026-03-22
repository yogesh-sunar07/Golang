package main

import "fmt"

const LoginToken string ="hello" //public cuz L is capital

func main() {
	var username string = "yogesh"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	var smallValue uint8  = 225
	fmt.Println(smallValue)
	fmt.Printf("Variable is of Type: %T \n", smallValue)

	var smallFloat float64  = 225.663636636383883838
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type: %T \n", smallFloat)

	//default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of Type: %T \n", anotherVariable)

	//implicit type (lexer)
	var website = "yogesh"
	fmt.Println(website)

	// no var style
	numberOfUser :=30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of Type: %T \n", LoginToken)
}
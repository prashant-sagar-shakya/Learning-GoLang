package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("Person is an Adult")
	} else {
		fmt.Println("Person is a Minor")
	}

	// if-else if-else ladder
	age = 25
	if age >= 18 && age <= 60 {
		fmt.Println("Person is an Adult")
	} else if age > 60 {
		fmt.Println("Person is an Senior Citizen")
	} else {
		fmt.Println("Person is a Minor")
	}

	// Logical operators in golang
	// && --> AND
	// || --> OR
	// ! --> NOT

	// if with short statement
	if age := 20; age >= 18 {
		fmt.Println("Person is an Adult")
	}

	// if with short statement and else
	if age := 20; age >= 18 {
		fmt.Println("Person is an Adult")
	} else {
		fmt.Println("Person is a Minor")
	}

	// if with short statement and else if-else ladder
	if age := 20; age >= 18 && age <= 60 {
		fmt.Println("Person is an Adult")
	} else if age > 60 {
		fmt.Println("Person is an Senior Citizen")
	} else {
		fmt.Println("Person is a Minor")
	}

	//  multi prints separated by comma
	fmt.Println("Hello", "Golang", "World")

	// Go does not have ternary operator
	// but we can use if-else to achieve the same

}

package main

import "fmt"

func main() {
	var name string = "Prashant"
	fmt.Println(name)

	// infer
	var name2 = "Golang"
	fmt.Println(name2)

	var age int = 30
	fmt.Println(age)

	var isStudent bool = true
	fmt.Println(isStudent)

	var height float64 = 5.9
	fmt.Println(height)

	// short variable declaration, shorthand syntax
	name3 := "Prashant"
	fmt.Println(name3)

	age2 := 30
	fmt.Println(age2)

	isStudent2 := true
	fmt.Println(isStudent2)

	height2 := 5.9
	fmt.Println(height2)

	// if we want to assign the value later then we must have to write the variable naming as var name type
	var name4 string
	name4 = "Prashant"
	fmt.Println(name4)

	var age3 int
	age3 = 30
	fmt.Println(age3)

	var isStudent3 bool
	isStudent3 = true
	fmt.Println(isStudent3)

	var height3 float64
	height3 = 5.9
	fmt.Println(height3)

	// we can also use var keyword to declare multiple variables
	var name5, age4, isStudent4, height4 = "Prashant", 30, true, 5.9
	fmt.Println(name5, age4, isStudent4, height4)

	// we can also use := to declare multiple variables
	name6, age5, isStudent5, height5 := "Prashant", 30, true, 5.9
	fmt.Println(name6, age5, isStudent5, height5)
}
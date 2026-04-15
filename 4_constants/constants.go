package main

import "fmt"

const age = 30 // --> we can declare constants in global scope too
// but using the shorthand := we can not declare constants in global scope
func main() {
	const name string = "golang"
	// name = "javascript" ---> can not assign to constant
	fmt.Println(name)
	fmt.Println(age)

	// we can also declare multiple constants in one go
	const (
		name1 = "golang"
		age1  = 30
	)
	fmt.Println(name1, age1)
}

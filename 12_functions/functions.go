package main

import "fmt"

// function syntax is func functionName(parameters) returnType {
// 	return value
// }
func add(a int, b int) int {
	return a + b
}

// same type parameters can be written as
func add2(a, b int) int {
	return a + b
}

// multiple return values
func add3(a, b int) (int, int) {
	return a + b, a - b
}

// assigning function a function
// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	result := add(1, 2)
	fmt.Println(result)

	result2 := add2(1, 2)
	fmt.Println(result2)

	result3, result4 := add3(1, 2)
	fmt.Println(result3, result4)

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn := processIt()
	fmt.Println(fn(6))
}

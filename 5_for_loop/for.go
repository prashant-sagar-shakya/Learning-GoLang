package main

import "fmt"

func main() {
	// There's only one loop in golang which is for loop
	// for initialization; condition; post-statement {}

	// While loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop (break added so code below remains reachable)
	// for {
	// 	println("1")
	// }

	// we also have continue and break keywords
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// Classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// range based loops exclude the last value and starts from 0
	// for range loop
	for i := range 10 {
		fmt.Println(i)
	}

	// for range loop with break
	for i := range 10 {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// for range loop with continue
	for i := range 10 {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

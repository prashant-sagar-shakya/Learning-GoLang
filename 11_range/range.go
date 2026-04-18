package main

import "fmt"

// Iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
	sum := 0
	// returning only value
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// printing index along with data using range
	for i, num := range nums {
		fmt.Println(i, num)
	}

	m := map[string]string{"fname": "john", "lname": "doe"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// returning only key
	for k := range m {
		fmt.Println(k)
	}
	// iterating over string characters, it will print index along with character of string in ASCII value/unicode code point (rune)
	for i, c := range "hello" {
		fmt.Println(i, c) // string(c) to convert rune to string
	}
}

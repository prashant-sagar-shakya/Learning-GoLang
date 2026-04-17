package main

import "fmt"

// numbered sequence of specific length

func main() {
	var nums [4]int // by default it will be initialized with 0
	fmt.Println(len(nums))
	fmt.Println(nums[0])

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	fmt.Println(nums)

	var vals [3]bool // by default it will be initialized with false
	vals[2] = true
	fmt.Println(vals)

	var names [3]string // by default it will be initialized with empty string
	names[0] = "Prashant"
	names[1] = "John"
	names[2] = "Doe"
	fmt.Println(names)

	// we can also use := to declare an array
	nums2 := [4]int{1, 2, 3, 4}
	fmt.Println(nums2)

	// we can also use ... to declare an array, it will automatically calculate the length of the array
	nums3 := [...]int{1, 2, 3, 4}
	fmt.Println(nums3)

	// we can also use range to iterate over an array
	for i, v := range nums3 {
		fmt.Println(i, v)
	}

	// 2-D array
	var matrix [3][3]int
	fmt.Println(matrix)

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9
	fmt.Println(matrix)

	// we can also use := to declare a 2-D array
	matrix2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix2)

	// Arrays is used when fixed size, that is predictable
	// Memory optimization
	// Cache friendly
	// Constant time access
}

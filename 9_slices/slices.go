package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slices are dynamic arrays, more flexible than arrays, more memory efficient than arrays
	// Slices are more cache friendly, more performant than arrays

	// make([]type, length, capacity) // capacity is optional, by default capacity is equal to length
	slice := make([]int, 5, 10) // it will create a slice of length 5 and capacity 10
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// append (if capacity is not enough, it will create a new array with double the capacity, like vectors in C++)
	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// copy (copy is used to copy elements from one slice to another)
	slice2 := make([]int, 5)
	copy(slice2, slice)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// delete (delete is not a built-in function, we have to use append to delete elements)
	slice = append(slice[:2], slice[3:]...) // this is how we delete elements from a slice, this line means take elements from index 0 to 2 and append elements from index 3 to the end, known as slice operator (:) where first index is inclusive and second index is exclusive, by default first index is 0 and second index is len(slice)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// slice package
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2)) // What it is doing is it is checking if the two slices are equal or not, it is not checking if the two slices are same or not, it is checking if the two slices have same elements in same order

	// slice from array (slice is a view of an array)
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[1:3]
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// slice from slice (slice is a view of an array)
	slice4 := slice3[1:2]
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	// nil slice (uninitialized slice is  known as nil slice, by default it will have length 0 and capacity 0 [])
	var slice5 []int
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	// empty slice (initialized slice, by default it will have length 0 and capacity 0 [])
	slice6 := []int{}
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// append to empty slice (if capacity is not enough, it will create a new array with double the capacity, like vectors in C++)
	slice6 = append(slice6, 1)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// append multiple values
	slice6 = append(slice6, 1, 2, 3, 4, 5)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// append slice to slice
	slice7 := []int{6, 7, 8, 9, 10}
	slice6 = append(slice6, slice7...)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// 2-D slice (2-D slice is a slice of slices)
	slice8 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(slice8)
	fmt.Println(len(slice8))
	fmt.Println(cap(slice8))
}

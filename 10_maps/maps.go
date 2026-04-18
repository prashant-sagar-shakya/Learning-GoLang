package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, dict
// map syntax: make(map[keyType]valueType) or map[keyType]valueType{elements}
func main() {
	// ceating map
	m := make(map[string]string)
	// setting an element
	m["name"] = "Prashant"
	m["age"] = "25"
	m["city"] = "Mumbai"
	fmt.Println(m)

	// get an element
	fmt.Println(m["name"], m["age"], m["city"])
	// Imp: if key does not exist it will return the zero value of the value type
	fmt.Println(m["country"])

	// delete an element
	delete(m, "age")
	fmt.Println(m)

	// clear a map
	clear(m)
	fmt.Println(m)

	// iterate over map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// map literal without make function
	m2 := map[string]string{
		"name": "Prashant",
		"age":  "25",
		"city": "Mumbai",
	}
	fmt.Println(m2, len(m2))

	// check if key exists
	_, ok := m["name"]
	if ok {
		fmt.Println("All OK", ok)
	} else {
		fmt.Println("Not OK")
	}

	// comparing maps
	fmt.Println(maps.Equal(m, m2))
}

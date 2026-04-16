package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Tuesday":
		fmt.Println("Tuesday")
	case "Wednesday":
		fmt.Println("Wednesday")
	case "Thursday":
		fmt.Println("Thursday")
	case "Friday":
		fmt.Println("Friday")
	case "Saturday":
		fmt.Println("Saturday")
	case "Sunday":
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Multicondition switch
	day = "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// type switch, we can assign a function output to another or we can also store a function in a variable
	whoAmI := func(i interface{}) {
		// switch i.type()
		switch t := i.(type) {
		case int:
			fmt.Println("I am an integer")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a boolean")
		default:
			fmt.Println("I am something else", t)
		}
	}

	whoAmI(10)
	whoAmI("Prashant")
	whoAmI(true)
	whoAmI(5.9)
}

package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write", i, "as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//Multiple expression in one statement. and use of default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend.")
	default:
		fmt.Println("It is a weekday")
	}

	//switch like if else logic

	var t = time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its a before noon.")
	default:
		fmt.Println("Its a after noon.")
	}

	whatAmI := func(i interface{}) {

		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool.")
		case int:
			fmt.Println("I'm a int.")
		default:
			fmt.Println("Do know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}

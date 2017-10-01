package main

import "fmt"

func main() {

	//The most basic type , with a single condition, a while like for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Classic For Loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//infinte for

	for {
		fmt.Println("Infinite")
		break
	}

	for n := 0; n <= 5; n++ {

		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

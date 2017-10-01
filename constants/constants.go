package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	//A numeric constant until it is given one
	const n float64 = 500000000
	fmt.Printf("Type: %T Value: %v\n", n, n)

	//const expression can perform arithmetic with arbitrary precision
	const d = 3e20 / n

	fmt.Println(int64(d))

	//A number can be given type by using it in a context that requires one.
	fmt.Println(math.Sin(n))

	fmt.Printf("Type: %T Value: %v\n", n, n)
}

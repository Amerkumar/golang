package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func three() (int, int, int) {
	return 1, 2, 3
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	d, e, f := three()
	fmt.Println(d, e, f)

}

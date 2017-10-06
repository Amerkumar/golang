//varidic function are called with any number of arguments
//like fmt.Println

package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3, 4)

	//if u have already multiple arguments then apply them using variadic function using func(slice...)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum(nums...)

}

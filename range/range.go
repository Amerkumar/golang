package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	//here first is index and second is value
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum :", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys: ", k)
	}

	//rune is just an integer value. It just like ASCII. Mapped to unicode codepoint.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

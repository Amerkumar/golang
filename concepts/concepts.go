//embedded types ; inner-type promotion
//composition
// methods
// interfaces

package main

import (
	"fmt"
	"reflect"
)

type person struct {
	fname string
	lname string
}

type secret struct {
	person
	licenseToKill bool
}

//type along with method
//It is related with type
//It ls like a object which knows about something and does something.

func (anything person) speak() {
	fmt.Println("My name is ", anything.fname, anything.lname)
}

func (sa secret) speak() {
	fmt.Println(sa.fname, sa.lname)
}

type saySomething interface {
	speak()
}

func woa(h saySomething) {

	h.speak()
}

func main() {

	p1 := person{
		"Charlie",
		"Chaplin",
	}

	woa(p1)

	sa1 := secret{
		person{"James",
			"Bond"},
		true,
	}
	woa(sa1)
	fmt.Println(reflect.TypeOf(sa1).String())
}

package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Println(a)

	//Print a's address
	fmt.Println(&a)

	//print type of &a
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(*b)  //give the value store at an address
	fmt.Println(*&a) //same line 17

	*b = 43
	println(a) //result: 43

	//address of b
	fmt.Println(&b)
}

package main

import "fmt"

func main() {

	//for normal
	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i, " ")
	// }

	//while loop
	// index := 0
	// for index < 10 {
	// 	fmt.Println(index)
	// 	index++
	// }

	//for infinity
	// for {
	// 	index++
	// 	fmt.Println(index)
	// }

	//for each range loop
	strs := []string{"one", "two", "three"}
	for i, str := range strs {
		fmt.Println(i, str)
	}
}

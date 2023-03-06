package main

import (
	"factory-method/factory"
	"fmt"
)

func main() {
	fmt.Print("Hello")
	mercedes := factory.GetCar("BMW")
}

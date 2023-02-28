package main

import "fmt"

func main() {

	//array
	array := [5]int{1, 2, 3, 4, 5}
	for i, e := range array {
		fmt.Println(i, e)
	}
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	//slice
	cards := []string{"Ace of Diamond", newCard()}
	cards = append(cards, "Six of Diamond")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

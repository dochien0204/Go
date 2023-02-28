package main

import "fmt"

// var book string

const pi = 3.14

func main() {
	//full declaration variable
	// var card string = "Chien ne"

	//short declare
	// card := "Card 1"
	// card = "Chien day nua ne"
	// book = "Hello"
	//Return function
	// card2 := newCard()
	// returnInt := returnInt()
	// returnBool := returnBoolean()
	// returnFloat64 := returnFloat64()
	// fmt.Println(card)
	// fmt.Println(card2)
	// fmt.Println(returnInt)
	// fmt.Println(returnBool)
	// fmt.Println(returnFloat64)
	// voidFunction()
	// var intValue, stringValue, boolValue = returnMultitype()
	// fmt.Println(intValue)
	// fmt.Println(stringValue)
	// fmt.Println(boolValue)

	// fmt.Println(book)

	//Deal Cards
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("\nStarted cards")
	// cards.print()

	// fmt.Println("\nDeal cards")
	// hand.print()

	// fmt.Println("\nRemaining cards")
	// remainingCards.print()

	//byteSlices
	cards := newDeck()
	// fmt.Println(cards.toString())

	//Write to file
	cards.saveToFile("my_cards")

	//Read from File
	fmt.Println("\nDeck from file")
	cardsFromFile := deckFromFile("my_cards")
	cardsFromFile.print()

	//Shuffle cards from file
	fmt.Println("\nShuffle cards from file")
	cardsFromFile.shuffle()
	cardsFromFile.print()

}

// func newCard() string {
// 	return "New Card is Generated from newCard functions"
// }

func returnInt() int {
	return 0
}

func returnFloat64() float64 {
	return 3.14
}

func returnBoolean() bool {
	return true
}

func voidFunction() {
	fmt.Println("Void function is called")
}

func returnMultitype() (int, string, bool) {
	return 6, "Hello", false
}

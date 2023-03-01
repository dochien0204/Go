package main

import "fmt"

type bot interface {
	GetGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.GetGreeting())
}

func (eb englishBot) GetGreeting() string {
	return "Hi there"
}

func (sb spanishBot) GetGreeting() string {
	return "Hola"
}

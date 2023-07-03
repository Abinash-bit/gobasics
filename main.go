package main

import "fmt"

func main() {
	cards := []string{"Ace of Dimonds", newCard()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

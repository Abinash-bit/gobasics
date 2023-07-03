package main

func main() {
	cards := deck{"Ace of Dimonds", newCard()}
	cards = append(cards, "Six of Spades")

}

func newCard() string {
	return "Five of Diamonds"
}

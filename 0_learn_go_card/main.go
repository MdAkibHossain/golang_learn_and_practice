package main

func main() {
	cards := newDeck()
	//cards.print()
	// hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_card")
	//cards = newDecFromFile("my_car")
	cards.suffle()
	cards.print()

}

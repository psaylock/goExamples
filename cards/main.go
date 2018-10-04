package main

/*
func main() {
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	card := "Ace of Spades"
	fmt.Println(card)
}
*/

func main() {
	/*	cards := newDeck()
		hand, remainingCards := deal(cards, 5)
		fmt.Println("cards:")
		cards.print()
		fmt.Println("hand:")
		hand.print()
		fmt.Println("Remaining:")
		remainingCards.print()
	*/
	/*
		cards := newDeck()
		//	fmt.Println(cards.toString())
		cards.saveToFile("my_cards.txt")
	*/
	/*
		cards := newDeckFromFile("my_cards.txt")
		fmt.Println(cards.toString())
	*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

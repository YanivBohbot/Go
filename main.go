package main

func main() {

	//cards := newCard()

	//hand, remainingCards := deal(cards,5)

	cards := newDeck()
	// fmt.Println(cards.toString())
	// greeting := "oijweoijewf"
	// fmt.Println([]byte(greeting))
	cards.saveToFile("my_cards")
}

package main

func main() {
	//var card string = "ace of space" which can also be written as the code below
	card := NewCard()
	fmt.println(card)

}

func NewCard() string {
	return "five of diamonds"
}

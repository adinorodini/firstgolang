import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = apppend(cards, "Six of Spades")
	fmt.Println(cards)

}

func newCard() string {
	return "five of diamonds"
} 
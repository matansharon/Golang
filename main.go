package main

import (
	"fmt"
)

func main() {
	
	list_of_cards := deck {1,2,3,4,5,6,7,8,9,10}

	fmt.Println(list_of_cards)
	list_of_cards.print()
	// for i := range list_of_cards {
	// 	fmt.Println(i)
	// }
}
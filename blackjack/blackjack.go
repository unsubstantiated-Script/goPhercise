package blackjack

import (
	"fmt"
	deck "goPhercise/go-deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func BlackJack() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		//Drawing off the first card and leaving the rest
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}

	var h Hand = cards[0:3]

	fmt.Println(h)
}

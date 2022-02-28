package main

import (
	"blackjack/deck"
	"fmt"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i, val := range h {
		strs[i] = val.String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + " , ***Hidden***"
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string

	for input != "s" {
		fmt.Println("player: ", player)
		fmt.Println("dealer: ", dealer.DealerString())
		fmt.Println("Whar will you do? (h)it , (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	fmt.Println("Final Hand")
	fmt.Println("player: ", player)
	fmt.Println("dealer: ", dealer)
}

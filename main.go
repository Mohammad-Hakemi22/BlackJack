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

func (h Hand) Score() int {
	minscore := h.MinScore()
	if minscore > 11 {
		return minscore
	}
	for _, card := range h {
		if card.Rank == deck.Ace {
			return minscore + 10 //for change value of Ace from 1 to 11; already Ace = 1 => 1+ 10 = 11
		}
	}
	return minscore
}

func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
		fmt.Println("player:", player)
		fmt.Println("dealer: ", dealer.DealerString())
		fmt.Println("Whar will you do? (h)it , (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("Final Hand")
	fmt.Println("player: ", player, "\tScore: ", pScore)
	fmt.Println("dealer: ", dealer, "\tScore: ", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("You Win")
	case pScore > dScore:
		fmt.Println("You Win")
	case pScore < dScore:
		fmt.Println("You Lose")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}

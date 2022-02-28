//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String() //String() -> generate by stringer in suit_string.go
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	minRank = Ace
	maxRank = King
)

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return absoluteRank(cards[i]) < absoluteRank(cards[j])
	})
	return cards
}

func absoluteRank(card Card) int {
	return (int(card.Suit) * int(maxRank)) + int(card.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(card []Card) []Card {
	res := make([]Card, len(card))
	perm := shuffleRand.Perm(len(card))
	for i, val := range perm {
		res[i] = card[val]
	}
	return res
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Rank: Rank(i), Suit: Joker})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var res []Card
		for _, c := range cards {
			if !f(c) {
				res = append(res, c)
			}
		}
		return res
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var res []Card
		for i := 0; i < n; i++ {
			res = append(res, cards...)
		}
		return res
	}
}

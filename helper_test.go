package card_test

import (
	"fmt"

	"github.com/makabe/card"
)

type cardsName string

const (
	nilCards         cardsName = "nil"
	empty            cardsName = "empty"
	standardDeck     cardsName = "standard deck"
	piquetPack       cardsName = "piquet pack"
	fourAces         cardsName = "4 aces"
	singleSpadeAce   cardsName = "single spade ace"
	doubleSpadeAce   cardsName = "double spade ace"
	singleInvalidAce cardsName = "single invalid ace"
)

func cards(name cardsName) card.Cards {
	switch name {
	case nilCards:
		return nil

	case empty:
		return card.Cards{}

	case standardDeck:
		return card.NewStandardDeck()

	case piquetPack:
		return card.NewPiquetPack()

	case fourAces:
		return card.Cards{card.HA, card.CA, card.DA, card.SA}

	case singleSpadeAce:
		return card.Cards{card.SA}

	case doubleSpadeAce:
		return card.Cards{card.SA, card.SA}

	case singleInvalidAce:
		return card.Cards{card.Card(52)} // !(4)A
	}

	panic(fmt.Sprint("unknown cards ", name))
}

type predicateName string

const (
	isValid      predicateName = "is valid"
	isAce        predicateName = "is ace"
	is6Pips      predicateName = "is 6 pips"
	isSpade      predicateName = "is spade"
	isHeart      predicateName = "is heart"
	isRed        predicateName = "is red"
	isPiquetCard predicateName = "is piquet card"
)

func predicate(name predicateName) card.Predicate {
	switch name {
	case isValid:
		return func(c card.Card) bool {
			return c.IsValid()
		}

	case isAce:
		return func(c card.Card) bool {
			return c.Rank() == card.Ace
		}

	case is6Pips:
		return func(c card.Card) bool {
			return c.Rank() == card.Six
		}

	case isSpade:
		return func(c card.Card) bool {
			return c.Suit() == card.Spade
		}

	case isHeart:
		return func(c card.Card) bool {
			return c.Suit() == card.Heart
		}

	case isRed:
		return func(c card.Card) bool {
			return c.Suit() == card.Heart || c.Suit() == card.Diamond
		}

	case isPiquetCard:
		return card.IsPiquetCard
	}

	panic(fmt.Sprint("unknown predicate ", name))
}

type comparatorName string

const (
	suitAsc        comparatorName = "suit asc"
	suitDesc       comparatorName = "suit desc"
	rankAsc        comparatorName = "rank asc"
	rankDesc       comparatorName = "rank desc"
	suitAscRankAsc comparatorName = "suit asc rank asc"
	rankAscSuitAsc comparatorName = "rank asc suit asc"
)

func comparator(name comparatorName) card.Comparator {
	switch name {
	case suitAsc:
		return func(c1, c2 card.Card) int {
			return int(c1.Suit()) - int(c2.Suit())
		}

	case suitDesc:
		return func(c1, c2 card.Card) int {
			return int(c1.Suit()) - int(c2.Suit())
		}

	case rankAsc:
		return func(c1, c2 card.Card) int {
			return int(c1.Rank()) - int(c2.Rank())
		}

	case rankDesc:
		return func(c1, c2 card.Card) int {
			return int(c2.Rank()) - int(c1.Rank())
		}

	case suitAscRankAsc:
		return func(c1, c2 card.Card) int {
			o := int(c1.Suit()) - int(c2.Suit())
			if o != 0 {
				return o
			}
			return int(c1.Rank()) - int(c2.Rank())
		}

	case rankAscSuitAsc:
		return func(c1, c2 card.Card) int {
			o := int(c1.Rank()) - int(c2.Rank())
			if o != 0 {
				return o
			}
			return int(c1.Suit()) - int(c2.Suit())
		}
	}

	panic(fmt.Sprint("unknown comparator ", name))
}

// shareArray reports whether or not x and y share the underlying array.
func shareArray(x, y card.Cards) bool {
	cx, cy := cap(x), cap(y)
	return cx > 0 && cy > 0 && &x[0:cx][cx-1] == &y[0:cy][cy-1]
}

// equal returns true if x and y both are nil or pointed values by x and y are the same,
// false otherwise.
func equal(x, y *card.Card) bool {
	return (x == nil && y == nil) || (x != nil && y != nil && *x == *y)
}

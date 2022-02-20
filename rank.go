package card

import "fmt"

const (
	numRanks = 13
)

// Rank represents a card rank.
type Rank uint8

// Valid ranks.
const (
	Ace Rank = iota
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

// Ranks returns an array containing all the valid ranks.
func Ranks() [13]Rank {
	return [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

// IsValid reports whether the rank is valid.
func (r Rank) IsValid() bool {
	return r <= King
}

// Name returns the full written name of the rank.
func (r Rank) Name() string { //nolint:cyclop
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	}
	return fmt.Sprintf("Invalid Rank(%d)", r)
}

// String returns the shorthand notation of the rank.
func (r Rank) String() string { //nolint:cyclop
	switch r {
	case Ace:
		return "A"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "T"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	}
	return fmt.Sprintf("!(%d)", r)
}

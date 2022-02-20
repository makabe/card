package card

import "fmt"

const (
	numSuits = 4
)

// SuitSymbols is an array containing the shorthand notations of each suit,
// used as the return value of the "Suit.String()".
// Each symbol corresponds to a club, diamond, heart, and spade from front to back.
var SuitSymbols = [4]string{ // nolint:gochecknoglobals
	"C", // symbol for Club
	"D", // symbol for Diamond
	"H", // symbol for Heart
	"S", // symbol for Spade
}

// Suit represents a card suit.
type Suit uint8

// Valid suits.
const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

// Suits returns an array containing all the valid suits.
func Suits() [4]Suit {
	return [4]Suit{Club, Diamond, Heart, Spade}
}

// IsValid reports whether the suit is valid.
func (s Suit) IsValid() bool {
	return s <= Spade
}

// Name returns the full written name of the suit.
func (s Suit) Name() string {
	switch s {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"
	}
	return fmt.Sprintf("Invalid Suit(%d)", s)
}

// String returns the shorthand notation of the suit.
// The notation can be changed from the definition of SuitSymbols.
func (s Suit) String() string {
	if !s.IsValid() {
		return fmt.Sprintf("!(%d)", s)
	}
	return SuitSymbols[s]
}

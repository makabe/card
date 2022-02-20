package card

import "fmt"

// Card represents a single playing card.
type Card uint8

// Valid cards.
const (
	CA Card = iota // Ace of Clubs
	C2             // Two of Clubs
	C3             // Three of Clubs
	C4             // Four of Clubs
	C5             // Five of Clubs
	C6             // Six of Clubs
	C7             // Seven of Clubs
	C8             // Eight of Clubs
	C9             // Nine of Clubs
	CT             // Ten of Clubs
	CJ             // Jack of Clubs
	CQ             // Queen of Clubs
	CK             // King of Clubs

	DA // Ace of Diamonds
	D2 // Two of Diamonds
	D3 // Three of Diamonds
	D4 // Four of Diamonds
	D5 // Five of Diamonds
	D6 // Six of Diamonds
	D7 // Seven of Diamonds
	D8 // Eight of Diamonds
	D9 // Nine of Diamonds
	DT // Ten of Diamonds
	DJ // Jack of Diamonds
	DQ // Queen of Diamonds
	DK // King of Diamonds

	HA // Ace of Hearts
	H2 // Two of Hearts
	H3 // Three of Hearts
	H4 // Four of Hearts
	H5 // Five of Hearts
	H6 // Six of Hearts
	H7 // Seven of Hearts
	H8 // Eight of Hearts
	H9 // Nine of Hearts
	HT // Ten of Hearts
	HJ // Jack of Hearts
	HQ // Queen of Hearts
	HK // King of Hearts

	SA // Ace of Spades
	S2 // Two of Spades
	S3 // Three of Spades
	S4 // Four of Spades
	S5 // Five of Spades
	S6 // Six of Spades
	S7 // Seven of Spades
	S8 // Eight of Spades
	S9 // Nine of Spades
	ST // Ten of Spades
	SJ // Jack of Spades
	SQ // Queen of Spades
	SK // King of Spades
)

// With returns a card with the suit and the rank.
func With(s Suit, r Rank) Card {
	if s.IsValid() && r.IsValid() {
		return Card(uint(s*numRanks) + uint(r))
	}
	return Card(numSuits * numRanks)
}

// IsValid reports whether the card is valid.
func (c Card) IsValid() bool {
	return c.Suit().IsValid() // c.Rank().IsValid() always returns true
}

// Name returns the full written name (e.g., Ace of Spades) of the card.
func (c Card) Name() string {
	if c.IsValid() {
		return fmt.Sprintf("%s of %ss", c.Rank().Name(), c.Suit().Name())
	}
	return fmt.Sprintf("Invalid Card(%d)", c)
}

// Rank returns the rank of the card.
// The return rank is always valid.
func (c Card) Rank() Rank {
	return Rank(c % numRanks)
}

// String returns the shorthand notation of the card (e.g., SA).
// The suit notation can be changed from the definition of SuitSymbols.
func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Suit(), c.Rank())
}

// Suit returns the suit of the card.
// A valid card always returns a valid suit, and an invalid card always returns an invalid suit.
func (c Card) Suit() Suit {
	return Suit(c / numRanks)
}

package card

const (
	standardDeckSize = 52
	piquetPackSize   = 32
)

// IsPiquetCard reports whether the card is part of a piquet pack or not.
func IsPiquetCard(c Card) bool {
	if !c.IsValid() {
		return false
	}
	r := c.Rank()
	if Two <= r && r <= Six {
		return false
	}
	return true
}

// IsPiquetPack reports whether or not the cards are a piquet pack.
func IsPiquetPack(cs Cards) bool {
	if len(cs) != piquetPackSize {
		return false
	}
	if cs.haveDuplicates() {
		return false
	}
	return cs.Every(IsPiquetCard)
}

// IsStandardDeck reports whether or not the cards are a standard 52-card deck.
func IsStandardDeck(cs Cards) bool {
	if len(cs) != standardDeckSize {
		return false
	}
	if cs.haveDuplicates() {
		return false
	}
	isValid := func(c Card) bool {
		return c.IsValid()
	}
	return cs.Every(isValid)
}

// NewPiquetPack returns cards representing a piquet pack,
// consisting of 32 cards, not including ranks 2 through 6.
func NewPiquetPack() Cards {
	pack := [piquetPackSize]Card{
		HA, H7, H8, H9, HT, HJ, HQ, HK,
		CA, C7, C8, C9, CT, CJ, CQ, CK,
		DK, DQ, DJ, DT, D9, D8, D7, DA,
		SK, SQ, SJ, ST, S9, S8, S7, SA,
	}
	return pack[:]
}

// NewStandardDeck returns cards representing a standard 52-card deck
// consisting of 4 suits with 13 ranks each.
func NewStandardDeck() Cards {
	// https://www.quora.com/When-you-buy-a-deck-of-cards-does-it-come-mixed-or-in-order/answer/Joshua-Burch-11
	pack := [standardDeckSize]Card{
		HA, H2, H3, H4, H5, H6, H7, H8, H9, HT, HJ, HQ, HK,
		CA, C2, C3, C4, C5, C6, C7, C8, C9, CT, CJ, CQ, CK,
		DK, DQ, DJ, DT, D9, D8, D7, D6, D5, D4, D3, D2, DA,
		SK, SQ, SJ, ST, S9, S8, S7, S6, S5, S4, S3, S2, SA,
	}
	return pack[:]
}

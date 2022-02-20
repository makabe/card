package card

// Comparator represents a function that compares the order of two cards.
// All comparators are considered to return a negative value if c1 is less than c2,
// a positive value if c1 is greater than c2, or 0 if both are equal,
// based on their intended order.
type Comparator func(c1, c2 Card) int

// Reversed returns a comparator that compares two cards in the reverse order of the comparator.
func (c Comparator) Reversed() Comparator {
	return func(c1, c2 Card) int {
		return c(c1, c2) * -1
	}
}

// Then returns a comparator with another comparator.
// If the comparator considers the two cards equal, the other determines the order.
func (c Comparator) Then(other Comparator) Comparator {
	return func(c1, c2 Card) int {
		o := c(c1, c2)
		if o != 0 {
			return o
		}
		return other(c1, c2)
	}
}

// less returns the sort.Interface.Less representation of the comparator.
func (c Comparator) less() func(Card, Card) bool {
	return func(c1, c2 Card) bool {
		return c(c1, c2) < 0
	}
}

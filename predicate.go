package card

// Predicate represents a predicate that takes a card as its only argument.
type Predicate func(Card) bool

// And returns a composed predicate that represents the logical conjunction (AND) of the predicate and another.
func (p Predicate) And(other Predicate) Predicate {
	return func(c Card) bool {
		return p(c) && other(c)
	}
}

// Not returns a predicate that represents the logical negation of the predicate.
func (p Predicate) Not() Predicate {
	return func(c Card) bool {
		return !p(c)
	}
}

// Or returns a composed predicate that represents the logical disjunction (OR) of the predicate and another.
func (p Predicate) Or(other Predicate) Predicate {
	return func(c Card) bool {
		return p(c) || other(c)
	}
}

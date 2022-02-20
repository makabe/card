package card

import (
	"math/rand"
	"sort"
)

// Cards represents zero or more playing cards.
//
// Cards consists of a slice of Card, and for-range form can be used to iterate over them.
//
//   hand := card.Cards{card.SA, card.SJ}
//   for _, c := range hand {
//     fmt.Println(c)
//   }
//
// Cards provides several non-destructive methods for common operations on playing cards.
// The non-destructive methods return newly allocated cards and leave the receiver cards unchanged.
//
//   deck := card.NewPiquetPack()
//
//   // shuffle
//   deck = deck.Shuffle(nil)
//
//   // deal 5 cards in batches of 3 and 2
//   hand1, talon := deck.Take(3)
//   hand2, talon := talon.Take(3)
//   hand1, talon = talon.Move(2, hand1)
//   hand2, talon = talon.Move(2, hand2)
//
//   // sort
//   hand1 = hand1.Sort(suitAscRankAsc)
//   hand2 = hand2.Sort(suitAscRankAsc)
//
//   fmt.Println(deck) // undealt
type Cards []Card

// Add returns new cards with other cards added on the bottom of the cards.
func (cs Cards) Add(other ...Card) Cards {
	if len(other) < 1 {
		return cs.Clone()
	}
	res := make(Cards, 0, len(cs)+len(other))
	res = append(res, cs...)
	res = append(res, other...)
	return res
}

// Any returns whether the predicate holds for at least one of the cards.
// If called on nil or empty cards, returns false.
func (cs Cards) Any(p Predicate) bool {
	for _, c := range cs {
		if p(c) {
			return true
		}
	}
	return false
}

// Bottom returns two values. One is new cards, and the other is a pointer to a new card.
// The new card is a copy of the bottom card of the cards,
// and the new cards are a copy of the others.
// If called on nil, the return values will also be nil;
// on empty, the return values will be empty and nil.
func (cs Cards) Bottom() (remaining Cards, bottom *Card) {
	if cs == nil {
		return nil, nil
	}
	n := len(cs)
	if n < 1 {
		return Cards{}, nil
	}
	rem := make(Cards, n-1)
	for i := range rem {
		rem[i] = cs[i]
	}
	c := cs[n-1]
	return rem, &c
}

// Clone returns new cards consisting of the same cards in the same order as the cards.
// If called on nil, the return value will also be nil.
func (cs Cards) Clone() Cards {
	if cs == nil {
		return nil
	}
	res := make(Cards, len(cs))
	copy(res, cs)
	return res
}

// Empty reports whether the number of cards is zero or not.
// If called on nil, returns true.
func (cs Cards) Empty() bool {
	return len(cs) == 0
}

// Every returns whether the predicate holds for all cards of the cards.
// If called on nil or empty cards, returns true.
func (cs Cards) Every(p Predicate) bool {
	for _, c := range cs {
		if !p(c) {
			return false
		}
	}
	return true
}

// Filter returns new cards consisting of all cards in the cards that satisfy the predicate.
// The order of the cards is preserved.
// If called on nil, the return value will also be nil.
func (cs Cards) Filter(p Predicate) Cards {
	if cs == nil {
		return nil
	}
	res := make(Cards, 0)
	for _, c := range cs {
		if p(c) {
			res = append(res, c)
		}
	}
	return res
}

// Include returns whether the cards include all the cards in the other.
// Duplicate cards are also considered as one card.
func (cs Cards) Include(other ...Card) bool {
	return Cards(other).Every(func(o Card) bool {
		return cs.Any(func(c Card) bool { return c == o })
	})
}

// Move returns a pair of new cards;
// one consists of all cards in the other cards and the first n cards of the cards,
// and the other consists of the remaining cards.
// If n is greater than the number of the cards,
// Move will not panic and replace n with the number and continue processing.
// If called on nil, the return values will be a copy of the other cards and nil.
func (cs Cards) Move(n uint, to Cards) (dst, src Cards) {
	taken, remaining := cs.Take(n)
	return to.Add(taken...), remaining
}

// Partition returns a pair of new cards,
// one is consisting of all cards in the cards that satisfy the predicate and
// the other is consisting of all cards in the cards that do not satisfy the predicate.
// The order of the cards is preserved.
// If called on nil, the return values will also be nil.
func (cs Cards) Partition(p Predicate) (satisfied, unsatisfied Cards) {
	if cs == nil {
		return nil, nil
	}
	satisfied = make(Cards, 0)
	unsatisfied = make(Cards, 0)
	for _, c := range cs {
		if p(c) {
			satisfied = append(satisfied, c)
		} else {
			unsatisfied = append(unsatisfied, c)
		}
	}
	return
}

// Remove returns new cards that do not contain the cards specified by the arguments.
// The order of the cards is preserved.
// Duplicate arguments are considered as one card.
// If called on nil, the return value will also be nil.
func (cs Cards) Remove(other ...Card) Cards {
	return cs.Filter(func(c Card) bool {
		return Cards(other).Every(func(o Card) bool { return o != c })
	})
}

// Reverse returns new cards with cards in reversed order.
// If called on nil, the return value will also be nil.
func (cs Cards) Reverse() Cards {
	if cs == nil {
		return nil
	}
	n := len(cs)
	res := make(Cards, n)
	for i := 0; i < n; i++ {
		res[i] = cs[n-1-i]
	}
	return res
}

// Shuffle returns new cards consisting of the same cards in a different order
// by shuffling the cards with the source of random numbers.
// If called on nil, the return value will also be nil.
func (cs Cards) Shuffle(r *rand.Rand) Cards {
	res := cs.Clone()
	swap := func(i, j int) { res[i], res[j] = res[j], res[i] }
	if r == nil {
		rand.Shuffle(len(res), swap)
	} else {
		r.Shuffle(len(res), swap)
	}
	return res
}

// Size returns the number of cards in the cards.
func (cs Cards) Size() uint {
	return uint(len(cs))
}

// Sort returns new cards with cards sorted by the comparator.
// If called on nil, the return value will also be nil.
func (cs Cards) Sort(c Comparator) Cards {
	if cs == nil {
		return nil
	}
	res := cs.Clone()
	less := c.less()
	sort.Slice(res, func(i, j int) bool {
		return less(res[i], res[j])
	})
	return res
}

// Take returns a pair of new cards, one consisting of the first n cards
// and the other consisting of the remaining cards.
// If n is greater than the number of the cards,
// Take will not panic and replace n with the number and continue processing.
// When called on nil, the return values will also be nil.
func (cs Cards) Take(n uint) (taken, remaining Cards) {
	return cs.splitAt(int(n))
}

// Top returns two values. One is a pointer to a copy of the top card of the cards,
// and the other is a copy of the other cards.
// If called on nil, the return values will also be nil;
// on empty, the return values will be empty and nil.
func (cs Cards) Top() (top *Card, remaining Cards) {
	if cs == nil {
		return nil, nil
	}
	n := len(cs)
	if n < 1 {
		return nil, Cards{}
	}
	rem := make(Cards, n-1)
	for i := range rem {
		rem[i] = cs[i+1]
	}
	c := cs[0]
	return &c, rem
}

// haveDuplicates reports whether the cards contain at least one pair of identical cards.
func (cs Cards) haveDuplicates() bool {
	m := make(map[Card]struct{})
	for _, c := range cs {
		_, ok := m[c]
		if ok {
			return true
		}
		m[c] = struct{}{}
	}
	return false
}

// splitAt returns a pair of new cards, one consisting of the first n cards
// and the other consisting of the remaining cards.
// If n is greater than the Size of the cards,
// splitAt will not panic but will replace n with Size and continue processing.
// If n is less than 0, -n indicates the number of lower cards.
// When called on nil, the return values will also be nil.
func (cs Cards) splitAt(n int) (upper, lower Cards) {
	if cs == nil {
		return nil, nil
	}
	csn := len(cs)
	if n < 0 {
		n = csn + n
	}
	if n < 0 {
		n = 0
	}
	if n > csn {
		n = csn
	}
	left := make(Cards, n)
	copy(left, cs)
	right := make(Cards, csn-n)
	copy(right, cs[n:])
	return left, right
}

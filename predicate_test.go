package card_test

import (
	"fmt"
	"testing"

	"github.com/makabe/card"
)

func TestPredicate_And(t *testing.T) {
	tests := []struct {
		p1Name predicateName
		p2Name predicateName
		card   card.Card
		want   bool
	}{
		{isSpade, isAce, card.SA, true},
		{isSpade, isAce, card.S2, false},
		{isSpade, isAce, card.HA, false},
		{isSpade, isAce, card.CA, false},
		{isSpade, isAce, card.DA, false},

		{isSpade, isRed, card.SA, false},
		{isSpade, isRed, card.S2, false},
		{isSpade, isRed, card.HA, false},
		{isSpade, isRed, card.CA, false},
		{isSpade, isRed, card.DA, false},

		{isAce, isRed, card.SA, false},
		{isAce, isRed, card.S2, false},
		{isAce, isRed, card.HA, true},
		{isAce, isRed, card.CA, false},
		{isAce, isRed, card.DA, true},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s %s", tt.p1Name, tt.p2Name)
		p1 := predicate(tt.p1Name)
		p2 := predicate(tt.p2Name)
		t.Run(name, func(t *testing.T) {
			if got := p1.And(p2)(tt.card); got != tt.want {
				t.Errorf("%#v.And(%#v)(%v) = %v, want %v", p1, p2, tt.card, got, tt.want)
			}
		})
	}
}

func TestPredicate_Not(t *testing.T) {
	tests := []struct {
		name predicateName
		card card.Card
		want bool
	}{
		{isSpade, card.SA, false},
		{isSpade, card.S2, false},
		{isSpade, card.HA, true},
		{isSpade, card.CA, true},
		{isSpade, card.DA, true},

		{isAce, card.SA, false},
		{isAce, card.S2, true},
		{isAce, card.HA, false},
		{isAce, card.CA, false},
		{isAce, card.DA, false},

		{isRed, card.SA, true},
		{isRed, card.S2, true},
		{isRed, card.HA, false},
		{isRed, card.CA, true},
		{isRed, card.DA, false},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s %s", tt.name, tt.card)
		pred := predicate(tt.name)
		t.Run(name, func(t *testing.T) {
			if got := pred.Not()(tt.card); got != tt.want {
				t.Errorf("%#v.Not()(%v) = %v, want %v", pred, tt.card, got, tt.want)
			}
		})
	}
}

func TestPredicate_Or(t *testing.T) {
	tests := []struct {
		p1Name predicateName
		p2Name predicateName
		card   card.Card
		want   bool
	}{
		{isSpade, isAce, card.SA, true},
		{isSpade, isAce, card.S2, true},
		{isSpade, isAce, card.HA, true},
		{isSpade, isAce, card.CA, true},
		{isSpade, isAce, card.DA, true},

		{isSpade, isRed, card.SA, true},
		{isSpade, isRed, card.S2, true},
		{isSpade, isRed, card.HA, true},
		{isSpade, isRed, card.CA, false},
		{isSpade, isRed, card.DA, true},

		{isAce, isRed, card.SA, true},
		{isAce, isRed, card.S2, false},
		{isAce, isRed, card.HA, true},
		{isAce, isRed, card.CA, true},
		{isAce, isRed, card.DA, true},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s %s", tt.p1Name, tt.p2Name)
		p1 := predicate(tt.p1Name)
		p2 := predicate(tt.p2Name)
		t.Run(name, func(t *testing.T) {
			if got := p1.Or(p2)(tt.card); got != tt.want {
				t.Errorf("%#v.Or(%#v)(%v) = %v, want %v", p1, p2, tt.card, got, tt.want)
			}
		})
	}
}

func ExamplePredicate_And() {
	isFaceCard := card.Predicate(func(c card.Card) bool {
		r := c.Rank()
		return r == card.Jack || r == card.Queen || r == card.King
	})

	isSpade := card.Predicate(func(c card.Card) bool {
		return c.Suit() == card.Spade
	})

	hand := card.Cards{card.SA, card.HA, card.DA, card.SK, card.HK}

	fmt.Println(hand.Filter(isFaceCard))
	fmt.Println(hand.Filter(isSpade))
	fmt.Println(hand.Filter(isFaceCard.And(isSpade)))
	// Output:
	// [SK HK]
	// [SA SK]
	// [SK]
}

func ExamplePredicate_Not() {
	isFaceCard := card.Predicate(func(c card.Card) bool {
		r := c.Rank()
		return r == card.Jack || r == card.Queen || r == card.King
	})

	hand := card.Cards{card.SA, card.SK, card.SQ, card.SJ, card.ST}

	fmt.Println(hand.Filter(isFaceCard))
	fmt.Println(hand.Filter(isFaceCard.Not()))
	// Output:
	// [SK SQ SJ]
	// [SA ST]
}

func ExamplePredicate_Or() {
	isFaceCard := card.Predicate(func(c card.Card) bool {
		r := c.Rank()
		return r == card.Jack || r == card.Queen || r == card.King
	})

	isSpade := card.Predicate(func(c card.Card) bool {
		return c.Suit() == card.Spade
	})

	hand := card.Cards{card.SA, card.HA, card.DA, card.SK, card.HK}

	fmt.Println(hand.Filter(isFaceCard))
	fmt.Println(hand.Filter(isSpade))
	fmt.Println(hand.Filter(isFaceCard.Or(isSpade)))
	// Output:
	// [SK HK]
	// [SA SK]
	// [SA SK HK]
}

func BenchmarkPredicate(b *testing.B) {
	p := func(c card.Card) bool {
		return c.Suit() == card.Spade
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SA)
	}
}

func BenchmarkPredicate_And(b *testing.B) {
	p1 := func(c card.Card) bool {
		return c.Suit() == card.Spade
	}
	p2 := func(c card.Card) bool {
		return c.Rank() == card.Ace
	}
	p := card.Predicate(p1).And(p2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SA)
	}
}

func BenchmarkPredicate_Handwritten_And(b *testing.B) {
	p := func(c card.Card) bool {
		return c.Suit() == card.Spade && c.Rank() == card.Ace
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SA)
	}
}

func BenchmarkPredicate_Not(b *testing.B) {
	p := func(c card.Card) bool {
		return c.Suit() == card.Spade
	}
	p = card.Predicate(p).Not()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SA)
	}
}

func BenchmarkPredicate_Handwritten_Not(b *testing.B) {
	p := func(c card.Card) bool {
		return c.Suit() != card.Spade
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SA)
	}
}

func BenchmarkPredicate_Or(b *testing.B) {
	p1 := func(c card.Card) bool {
		return c.Suit() == card.Spade
	}
	p2 := func(c card.Card) bool {
		return c.Rank() == card.Ace
	}
	p := card.Predicate(p1).Or(p2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.HA)
	}
}

func BenchmarkPredicate_Handwritten_Or(b *testing.B) {
	p := func(c card.Card) bool {
		return c.Suit() == card.Spade || c.Rank() == card.Ace
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.HA)
	}
}

func BenchmarkPredicate_And_Or(b *testing.B) {
	p1 := func(c card.Card) bool {
		return c.Suit() == card.Spade
	}
	p2 := func(c card.Card) bool {
		return c.Rank() == card.Ace
	}
	p3 := func(c card.Card) bool {
		return c.Rank() == card.King
	}
	p := card.Predicate(p1).And(p2).Or(p3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SK)
	}
}

func BenchmarkPredicate_Handwritten_And_Or(b *testing.B) {
	p := func(c card.Card) bool {
		return c.Suit() == card.Spade && c.Rank() == card.Ace || c.Rank() == card.King
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SK)
	}
}

func BenchmarkPredicate_And_Or_Not(b *testing.B) {
	p1 := func(c card.Card) bool {
		return c.Suit() == card.Spade
	}
	p2 := func(c card.Card) bool {
		return c.Rank() == card.Ace
	}
	p3 := func(c card.Card) bool {
		return c.Rank() == card.King
	}
	p := card.Predicate(p1).And(p2).Or(p3).Not()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SK)
	}
}

func BenchmarkPredicate_Handwritten_And_Or_Not(b *testing.B) {
	p := func(c card.Card) bool {
		return !(c.Suit() == card.Spade && c.Rank() == card.Ace || c.Rank() == card.King)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p(card.SK)
	}
}

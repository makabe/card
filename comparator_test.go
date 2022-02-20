package card_test

import (
	"fmt"
	"testing"

	"github.com/makabe/card"
)

func TestComparator_Reversed(t *testing.T) {
	tests := []struct {
		name comparatorName
		c1   card.Card
		c2   card.Card
		want int
	}{
		{suitAscRankAsc, card.D2, card.D2, 0},
		{suitAscRankAsc, card.D2, card.C2, -1},
		{suitAscRankAsc, card.D2, card.H2, 1},
		{suitAscRankAsc, card.D2, card.DA, -1},
		{suitAscRankAsc, card.D2, card.D3, 1},
		{suitAscRankAsc, card.D2, card.C3, -1},
		{suitAscRankAsc, card.D2, card.HA, 1},

		{rankAscSuitAsc, card.D2, card.D2, 0},
		{rankAscSuitAsc, card.D2, card.C2, -1},
		{rankAscSuitAsc, card.D2, card.H2, 1},
		{rankAscSuitAsc, card.D2, card.DA, -1},
		{rankAscSuitAsc, card.D2, card.D3, 1},
		{rankAscSuitAsc, card.D2, card.C3, 1},
		{rankAscSuitAsc, card.D2, card.HA, -1},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s %s %s", tt.name, tt.c1, tt.c2)
		comp := comparator(tt.name)
		t.Run(name, func(t *testing.T) {
			if got := comp.Reversed()(tt.c1, tt.c2); got != tt.want {
				t.Errorf("%#v.Reversed()(%v, %v) = %v, want %v", comp, tt.c1, tt.c2, got, tt.want)
			}
		})
	}
}

func TestComparator_Then(t *testing.T) {
	tests := []struct {
		name      comparatorName
		otherName comparatorName
		c1        card.Card
		c2        card.Card
		want      int
	}{
		{suitAsc, rankAsc, card.D2, card.D2, 0},
		{suitAsc, rankAsc, card.D2, card.DA, 1},
		{suitAsc, rankAsc, card.D2, card.D3, -1},
		{suitAsc, rankAsc, card.D2, card.HA, -1},
		{suitAsc, rankAsc, card.D2, card.C3, 1},

		{suitAsc, rankDesc, card.D2, card.D2, 0},
		{suitAsc, rankDesc, card.D2, card.DA, -1},
		{suitAsc, rankDesc, card.D2, card.D3, 1},
		{suitAsc, rankDesc, card.D2, card.HA, -1},
		{suitAsc, rankDesc, card.D2, card.C3, 1},

		{rankAsc, suitAsc, card.D2, card.D2, 0},
		{rankAsc, suitAsc, card.D2, card.DA, 1},
		{rankAsc, suitAsc, card.D2, card.D3, -1},
		{rankAsc, suitAsc, card.D2, card.HA, 1},
		{rankAsc, suitAsc, card.D2, card.C3, -1},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v %v %s %s", tt.name, tt.otherName, tt.c1, tt.c2)
		comp := comparator(tt.name)
		other := comparator(tt.otherName)
		composed := comp.Then(other)
		t.Run(name, func(t *testing.T) {
			if got := composed(tt.c1, tt.c2); got != tt.want {
				t.Errorf("%#v.Then(%#v)(%v, %v) = %v, want %v", comp, other, tt.c1, tt.c2, got, tt.want)
			}
		})
	}
}

func TestComparator_less(t *testing.T) {
	tests := []struct {
		name comparatorName
		c1   card.Card
		c2   card.Card
		want bool
	}{
		{suitAscRankAsc, card.D2, card.D2, false},
		{suitAscRankAsc, card.D2, card.C2, false},
		{suitAscRankAsc, card.D2, card.H2, true},
		{suitAscRankAsc, card.D2, card.DA, false},
		{suitAscRankAsc, card.D2, card.D3, true},
		{suitAscRankAsc, card.D2, card.C3, false},
		{suitAscRankAsc, card.D2, card.HA, true},

		{rankAscSuitAsc, card.D2, card.D2, false},
		{rankAscSuitAsc, card.D2, card.C2, false},
		{rankAscSuitAsc, card.D2, card.H2, true},
		{rankAscSuitAsc, card.D2, card.DA, false},
		{rankAscSuitAsc, card.D2, card.D3, true},
		{rankAscSuitAsc, card.D2, card.C3, true},
		{rankAscSuitAsc, card.D2, card.HA, false},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s %s %s", tt.name, tt.c1, tt.c2)
		comp := comparator(tt.name)
		t.Run(name, func(t *testing.T) {
			less := card.ExportComparatorLess(comp) // comp.less()
			if got := less(tt.c1, tt.c2); got != tt.want {
				t.Errorf("%#v.less()(%v, %v) = %v, want %v", comp, tt.c1, tt.c2, got, tt.want)
			}
		})
	}
}

func ExampleComparator() {
	suitDescRankDesc := func(c1, c2 card.Card) int {
		return int(c2) - int(c1)
	}

	hand := card.Cards{card.SA, card.HA, card.CA, card.S4, card.C2}

	fmt.Println(hand.Sort(suitDescRankDesc))
	// Output:
	// [S4 SA HA C2 CA]
}

func ExampleComparator_Reversed() {
	suitAscRankAsc := func(c1, c2 card.Card) int {
		return int(c1) - int(c2)
	}

	suitDescRankDesc := card.Comparator(suitAscRankAsc).Reversed()

	hand := card.Cards{card.SA, card.HA, card.CA, card.S4, card.C2}

	fmt.Println(hand.Sort(suitAscRankAsc))
	fmt.Println(hand.Sort(suitDescRankDesc))
	// Output:
	// [CA C2 HA SA S4]
	// [S4 SA HA C2 CA]
}

func ExampleComparator_Then() {
	suitAsc := func(c1, c2 card.Card) int {
		return int(c1.Suit()) - int(c2.Suit())
	}

	rankAsc := func(c1, c2 card.Card) int {
		return int(c1.Rank()) - int(c2.Rank())
	}

	suitAscRankAsc := card.Comparator(suitAsc).Then(rankAsc)
	rankAscSuitAsc := card.Comparator(rankAsc).Then(suitAsc)

	hand := card.Cards{card.SA, card.HA, card.CA, card.S4, card.C2}

	fmt.Println(hand.Sort(suitAscRankAsc))
	fmt.Println(hand.Sort(rankAscSuitAsc))
	// Output:
	// [CA C2 HA SA S4]
	// [CA HA SA C2 S4]
}

func BenchmarkComparator(b *testing.B) {
	comp := func(c1, c2 card.Card) int {
		return int(c1) - int(c2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comp(card.SA, card.S2)
	}
}

func BenchmarkComparator_Reversed(b *testing.B) {
	comp := func(c1, c2 card.Card) int {
		return int(c1) - int(c2)
	}
	comp = card.Comparator(comp).Reversed()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comp(card.SA, card.S2)
	}
}

func BenchmarkComparator_Handwritten_Reversed(b *testing.B) {
	comp := func(c1, c2 card.Card) int {
		return int(c2) - int(c1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comp(card.SA, card.S2)
	}
}

func BenchmarkComparator_Then(b *testing.B) {
	suitAsc := func(c1, c2 card.Card) int {
		return int(c1.Suit()) - int(c2.Suit())
	}
	rankAsc := func(c1, c2 card.Card) int {
		return int(c1.Rank()) - int(c2.Rank())
	}
	comp := card.Comparator(suitAsc).Then(rankAsc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comp(card.SA, card.S2)
	}
}

func BenchmarkComparator_Handwritten_Then(b *testing.B) {
	comp := func(c1, c2 card.Card) int {
		res := int(c1.Suit()) - int(c2.Suit())
		if res != 0 {
			return res
		}
		return int(c1.Rank()) - int(c2.Rank())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		comp(card.SA, card.S2)
	}
}

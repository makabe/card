package card_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/makabe/card"
)

func TestIsPiquetCard(t *testing.T) {
	tests := []struct {
		name string
		rank card.Rank
		want bool
	}{
		{"Ace", card.Ace, true},
		{"Two", card.Two, false},
		{"Three", card.Three, false},
		{"Four", card.Four, false},
		{"Five", card.Five, false},
		{"Six", card.Six, false},
		{"Seven", card.Seven, true},
		{"Eight", card.Eight, true},
		{"Nine", card.Nine, true},
		{"Ten", card.Ten, true},
		{"Jack", card.Jack, true},
		{"Queen", card.Queen, true},
		{"King", card.King, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range card.Suits() {
				c := card.Card(uint(s)*13 + uint(tt.rank))
				if got := card.IsPiquetCard(c); got != tt.want {
					t.Errorf("IsPiquetCard() = %v, want %v", got, tt.want)
				}
			}
		})
	}
	for i := 52; i < 256; i++ {
		c := card.Card(i)
		t.Run(fmt.Sprintf("invalid card %s", c), func(t *testing.T) {
			if card.IsPiquetCard(c) {
				t.Errorf("IsPiquetCard() = true, want false")
			}
		})
	}
}

func TestIsPiquetPack(t *testing.T) {
	piquetPack := card.NewPiquetPack()
	shuffled := piquetPack.Shuffle(nil)
	long := append(piquetPack.Clone(), card.C7) // cannot exclude duplicates
	short := piquetPack[:(len(piquetPack) - 1)]
	duplicate := append(short.Clone(), card.S7)
	invalid := append(short.Clone(), card.Card(52))
	nonPiquet := append(short.Clone(), card.C6)
	tests := []struct {
		name  string
		cards card.Cards
		want  bool
	}{
		{"nil", nil, false},
		{"long", long, false},
		{"short", short, false},
		{"has duplicate card", duplicate, false},
		{"has invalid card", invalid, false},
		{"has non piquet card", nonPiquet, false},
		{"32 valid unique card", piquetPack, true},
		{"32 valid unique card (shuffled)", shuffled, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := card.IsPiquetPack(tt.cards); got != tt.want {
				t.Errorf("IsPiquetPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStandardDeck(t *testing.T) {
	stdDeck := card.NewStandardDeck()
	shuffled := stdDeck.Shuffle(nil)
	long := append(stdDeck.Clone(), card.C2) // cannot exclude duplicates
	short := stdDeck[:51]
	duplicate := append(short.Clone(), card.C3)
	invalid := append(short.Clone(), card.Card(52))
	tests := []struct {
		name  string
		cards card.Cards
		want  bool
	}{
		{"nil", nil, false},
		{"long", long, false},
		{"short", short, false},
		{"has duplicate card", duplicate, false},
		{"has invalid card", invalid, false},
		{"52 valid unique card", stdDeck, true},
		{"52 valid unique card (shuffled)", shuffled, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := card.IsStandardDeck(tt.cards); got != tt.want {
				t.Errorf("IsStandardDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPiquetPack(t *testing.T) {
	want := make(card.Cards, 0, 32)
	for _, s := range []card.Suit{card.Heart, card.Club} {
		for _, r := range []card.Rank{
			card.Ace, card.Seven, card.Eight, card.Nine,
			card.Ten, card.Jack, card.Queen, card.King,
		} {
			c := card.Card(uint(s)*13 + uint(r))
			want = append(want, c)
		}
	}
	for _, s := range []card.Suit{card.Diamond, card.Spade} {
		for _, r := range []card.Rank{
			card.King, card.Queen, card.Jack, card.Ten,
			card.Nine, card.Eight, card.Seven, card.Ace,
		} {
			c := card.Card(uint(s)*13 + uint(r))
			want = append(want, c)
		}
	}
	got := card.NewPiquetPack()
	t.Run("items", func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("NewPiquetPack() = %v, want %v", got, want)
		}
	})
	t.Run("not sharing underlying array", func(t *testing.T) {
		if shareArray(got, card.NewPiquetPack()) {
			t.Error("NewPiquetPack() returns slice sharing underlying array")
		}
	})
}

func TestNewStandardDeck(t *testing.T) {
	want := make(card.Cards, 0, 52)
	for _, s := range []card.Suit{card.Heart, card.Club} {
		for r := card.Ace; r <= card.King; r++ {
			c := card.Card(uint(s)*13 + uint(r))
			want = append(want, c)
		}
	}
	for _, s := range []card.Suit{card.Diamond, card.Spade} {
		for r := int(card.King); r >= int(card.Ace); r-- {
			c := card.Card(uint(s)*13 + uint(r))
			want = append(want, c)
		}
	}
	got := card.NewStandardDeck()
	t.Run("items", func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("NewStandardDeck() = %v, want %v", got, want)
		}
	})
	t.Run("not sharing underlying array", func(t *testing.T) {
		if shareArray(got, card.NewStandardDeck()) {
			t.Error("NewStandardDeck() returns slice sharing underlying array")
		}
	})
}

func ExampleIsPiquetCard() {
	for c := card.HA; c <= 52; c++ {
		fmt.Println(c, card.IsPiquetCard(c))
	}
	// Output:
	// HA true
	// H2 false
	// H3 false
	// H4 false
	// H5 false
	// H6 false
	// H7 true
	// H8 true
	// H9 true
	// HT true
	// HJ true
	// HQ true
	// HK true
	// SA true
	// S2 false
	// S3 false
	// S4 false
	// S5 false
	// S6 false
	// S7 true
	// S8 true
	// S9 true
	// ST true
	// SJ true
	// SQ true
	// SK true
	// !(4)A false
}

func ExampleIsPiquetPack() {
	p := func(name string, cs card.Cards) {
		fmt.Println(name, card.IsPiquetPack(cs))
	}

	piquetPack := card.NewPiquetPack()

	p("piquetPack", piquetPack)            // [HA H7 ... S7 SA]
	p("shuffled", piquetPack.Shuffle(nil)) // [C7 DQ ... CA C9]
	p("long", append(piquetPack, card.HA)) // [HA H7 ... S7 SA HA]

	short := piquetPack[1:]

	p("short", short)                          // [H7 H8 ... SA]
	p("invalid", append(short, card.Card(52))) // [H7 H8 ... SA !A]
	p("include2", append(short, card.S2))      // [H7 H8 ... SA S2]
	p("duplicated", append(short, card.SA))    // [H7 H8 ... SA SA]
	p("rotated", append(short, card.HA))       // [H7 H8 ... SA HA]
	// Output:
	// piquetPack true
	// shuffled true
	// long false
	// short false
	// invalid false
	// include2 false
	// duplicated false
	// rotated true
}

func ExampleIsStandardDeck() {
	p := func(name string, cs card.Cards) {
		fmt.Println(name, card.IsStandardDeck(cs))
	}

	stdDeck := card.NewStandardDeck()

	p("stdDeck", stdDeck)               // [HA H2 ... S2 SA]
	p("shuffled", stdDeck.Shuffle(nil)) // [S5 DJ ... D6 C7]
	p("long", append(stdDeck, card.HA)) // [HA H2 ... S2 SA HA]

	short := stdDeck[1:]

	p("short", short)                          // [H2 H3 ... SA]
	p("invalid", append(short, card.Card(52))) // [H2 H3 ... SA !A]
	p("duplicated", append(short, card.SA))    // [H2 H3 ... SA SA]
	p("rotated", append(short, card.HA))       // [H2 H3 ... SA HA]
	// Output:
	// stdDeck true
	// shuffled true
	// long false
	// short false
	// invalid false
	// duplicated false
	// rotated true
}

func ExampleNewPiquetPack() {
	deck := card.NewPiquetPack()

	fmt.Println(deck)
	fmt.Println(deck.Size())
	fmt.Println(deck.Every(card.IsPiquetCard))
	// Output:
	// [HA H7 H8 H9 HT HJ HQ HK CA C7 C8 C9 CT CJ CQ CK DK DQ DJ DT D9 D8 D7 DA SK SQ SJ ST S9 S8 S7 SA]
	// 32
	// true
}

//nolint:lll
func ExampleNewStandardDeck() {
	isValid := func(c card.Card) bool { return c.IsValid() }

	deck := card.NewStandardDeck()

	fmt.Println(deck)
	fmt.Println(deck.Size())
	fmt.Println(deck.Every(isValid))
	// Output:
	// [HA H2 H3 H4 H5 H6 H7 H8 H9 HT HJ HQ HK CA C2 C3 C4 C5 C6 C7 C8 C9 CT CJ CQ CK DK DQ DJ DT D9 D8 D7 D6 D5 D4 D3 D2 DA SK SQ SJ ST S9 S8 S7 S6 S5 S4 S3 S2 SA]
	// 52
	// true
}

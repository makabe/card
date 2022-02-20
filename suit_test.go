package card_test

import (
	"fmt"
	"testing"

	"github.com/makabe/card"
)

func TestSuits(t *testing.T) {
	want := [4]card.Suit{card.Club, card.Diamond, card.Heart, card.Spade}
	if got := card.Suits(); got != want {
		t.Errorf("Suits() = %v, want %v", got, want)
	}
}

func TestSuit_LessThan(t *testing.T) {
	tests := []struct {
		less    card.Suit
		greater card.Suit
	}{
		{card.Club, card.Diamond},
		{card.Diamond, card.Heart},
		{card.Heart, card.Spade},
	}
	for _, tt := range tests {
		name := fmt.Sprint(tt.less, "<", tt.greater)
		t.Run(name, func(t *testing.T) {
			if got := tt.less < tt.greater; got != true {
				t.Errorf("%s < %s = %v, want true", tt.less, tt.greater, got)
			}
		})
	}
}

func TestSuit_IsValid(t *testing.T) {
	tests := []struct {
		name string
		s    card.Suit
		want bool
	}{
		{"Club", card.Club, true},
		{"Diamond", card.Diamond, true},
		{"Heart", card.Heart, true},
		{"Spade", card.Spade, true},
		{"greater than Spade", (card.Spade + 1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValid(); got != tt.want {
				t.Errorf("Suit.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuit_Name(t *testing.T) {
	tests := []struct {
		name string
		s    card.Suit
		want string
	}{
		{"Club", card.Club, "Club"},
		{"Diamond", card.Diamond, "Diamond"},
		{"Heart", card.Heart, "Heart"},
		{"Spade", card.Spade, "Spade"},
		{"greater than Spade", (card.Spade + 1), "Invalid Suit(4)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Name(); got != tt.want {
				t.Errorf("Suit.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuit_String(t *testing.T) {
	tests := []struct {
		name string
		s    card.Suit
		want string
	}{
		{"Club", card.Club, "C"},
		{"Diamond", card.Diamond, "D"},
		{"Heart", card.Heart, "H"},
		{"Spade", card.Spade, "S"},
		{"greater than Spade", (card.Spade + 1), "!(4)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Suit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSuit_String() {
	fmt.Println(card.SuitSymbols)
	fmt.Println(card.Spade)

	card.SuitSymbols = [4]string{
		"\u2663", // ♣
		"\u2662", // ♢
		"\u2661", // ♡
		"\u2660", // ♠
	}
	fmt.Println(card.Spade)
	// Output:
	// [C D H S]
	// S
	// ♠
	card.SuitSymbols = [4]string{"C", "D", "H", "S"}
}

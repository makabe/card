package card_test

import (
	"fmt"
	"testing"

	"github.com/makabe/card"
)

func TestRanks(t *testing.T) {
	want := [13]card.Rank{
		card.Ace, card.Two, card.Three, card.Four, card.Five, card.Six, card.Seven,
		card.Eight, card.Nine, card.Ten, card.Jack, card.Queen, card.King,
	}
	if got := card.Ranks(); got != want {
		t.Errorf("Ranks() = %v, want %v", got, want)
	}
}

func TestRank_LessThan(t *testing.T) {
	tests := []struct {
		less    card.Rank
		greater card.Rank
	}{
		{card.Ace, card.Two},
		{card.Two, card.Three},
		{card.Three, card.Four},
		{card.Four, card.Five},
		{card.Five, card.Six},
		{card.Six, card.Seven},
		{card.Seven, card.Eight},
		{card.Eight, card.Nine},
		{card.Nine, card.Ten},
		{card.Ten, card.Jack},
		{card.Jack, card.Queen},
		{card.Queen, card.King},
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

func TestRank_IsValid(t *testing.T) {
	tests := []struct {
		name string
		r    card.Rank
		want bool
	}{
		{"Ace", card.Ace, true},
		{"Two", card.Two, true},
		{"Three", card.Three, true},
		{"Four", card.Four, true},
		{"Five", card.Five, true},
		{"Six", card.Six, true},
		{"Seven", card.Seven, true},
		{"Eight", card.Eight, true},
		{"Nine", card.Nine, true},
		{"Ten", card.Ten, true},
		{"Jack", card.Jack, true},
		{"Queen", card.Queen, true},
		{"King", card.King, true},
		{"higher than King", (card.King + 1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsValid(); got != tt.want {
				t.Errorf("Rank.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_Name(t *testing.T) {
	tests := []struct {
		name string
		r    card.Rank
		want string
	}{
		{"Ace", card.Ace, "Ace"},
		{"Two", card.Two, "Two"},
		{"Three", card.Three, "Three"},
		{"Four", card.Four, "Four"},
		{"Five", card.Five, "Five"},
		{"Six", card.Six, "Six"},
		{"Seven", card.Seven, "Seven"},
		{"Eight", card.Eight, "Eight"},
		{"Nine", card.Nine, "Nine"},
		{"Ten", card.Ten, "Ten"},
		{"Jack", card.Jack, "Jack"},
		{"Queen", card.Queen, "Queen"},
		{"King", card.King, "King"},
		{"higher than King", (card.King + 1), "Invalid Rank(13)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Name(); got != tt.want {
				t.Errorf("Rank.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_String(t *testing.T) {
	tests := []struct {
		name string
		r    card.Rank
		want string
	}{
		{"Ace", card.Ace, "A"},
		{"Two", card.Two, "2"},
		{"Three", card.Three, "3"},
		{"Four", card.Four, "4"},
		{"Five", card.Five, "5"},
		{"Six", card.Six, "6"},
		{"Seven", card.Seven, "7"},
		{"Eight", card.Eight, "8"},
		{"Nine", card.Nine, "9"},
		{"Ten", card.Ten, "T"},
		{"Jack", card.Jack, "J"},
		{"Queen", card.Queen, "Q"},
		{"King", card.King, "K"},
		{"higher than King", (card.King + 1), "!(13)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Rank.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

package card_test

import (
	"fmt"
	"testing"

	"github.com/makabe/card"
)

func TestWith(t *testing.T) {
	tests := []struct {
		suit card.Suit
		rank card.Rank
		want card.Card
	}{
		// Clubs
		{card.Club, card.Ace, card.CA},
		{card.Club, card.Two, card.C2},
		{card.Club, card.Three, card.C3},
		{card.Club, card.Four, card.C4},
		{card.Club, card.Five, card.C5},
		{card.Club, card.Six, card.C6},
		{card.Club, card.Seven, card.C7},
		{card.Club, card.Eight, card.C8},
		{card.Club, card.Nine, card.C9},
		{card.Club, card.Ten, card.CT},
		{card.Club, card.Jack, card.CJ},
		{card.Club, card.Queen, card.CQ},
		{card.Club, card.King, card.CK},
		// Diamonds
		{card.Diamond, card.Ace, card.DA},
		{card.Diamond, card.Two, card.D2},
		{card.Diamond, card.Three, card.D3},
		{card.Diamond, card.Four, card.D4},
		{card.Diamond, card.Five, card.D5},
		{card.Diamond, card.Six, card.D6},
		{card.Diamond, card.Seven, card.D7},
		{card.Diamond, card.Eight, card.D8},
		{card.Diamond, card.Nine, card.D9},
		{card.Diamond, card.Ten, card.DT},
		{card.Diamond, card.Jack, card.DJ},
		{card.Diamond, card.Queen, card.DQ},
		{card.Diamond, card.King, card.DK},
		// Hearts
		{card.Heart, card.Ace, card.HA},
		{card.Heart, card.Two, card.H2},
		{card.Heart, card.Three, card.H3},
		{card.Heart, card.Four, card.H4},
		{card.Heart, card.Five, card.H5},
		{card.Heart, card.Six, card.H6},
		{card.Heart, card.Seven, card.H7},
		{card.Heart, card.Eight, card.H8},
		{card.Heart, card.Nine, card.H9},
		{card.Heart, card.Ten, card.HT},
		{card.Heart, card.Jack, card.HJ},
		{card.Heart, card.Queen, card.HQ},
		{card.Heart, card.King, card.HK},
		// Spades
		{card.Spade, card.Ace, card.SA},
		{card.Spade, card.Two, card.S2},
		{card.Spade, card.Three, card.S3},
		{card.Spade, card.Four, card.S4},
		{card.Spade, card.Five, card.S5},
		{card.Spade, card.Six, card.S6},
		{card.Spade, card.Seven, card.S7},
		{card.Spade, card.Eight, card.S8},
		{card.Spade, card.Nine, card.S9},
		{card.Spade, card.Ten, card.ST},
		{card.Spade, card.Jack, card.SJ},
		{card.Spade, card.Queen, card.SQ},
		{card.Spade, card.King, card.SK},
		// invalid
		{card.Suit(4), card.Ace, card.Card(52)},
		{card.Suit(4), card.Two, card.Card(52)},
		{card.Club, card.Rank(13), card.Card(52)},    // not DA
		{card.Diamond, card.Rank(13), card.Card(52)}, // not HA
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.suit, tt.rank), func(t *testing.T) {
			if got := card.With(tt.suit, tt.rank); got != tt.want {
				t.Errorf("With(%v, %v) = %v, want %v", tt.suit, tt.rank, got, tt.want)
			}
		})
	}
}

func TestCard_IsValid(t *testing.T) {
	tests := []struct {
		name string
		card card.Card
		want bool
	}{
		// Clubs
		{"CA", card.CA, true},
		{"C2", card.C2, true},
		{"C3", card.C3, true},
		{"C4", card.C4, true},
		{"C5", card.C5, true},
		{"C6", card.C6, true},
		{"C7", card.C7, true},
		{"C8", card.C8, true},
		{"C9", card.C9, true},
		{"CT", card.CT, true},
		{"CJ", card.CJ, true},
		{"CQ", card.CQ, true},
		{"CK", card.CK, true},
		// Diamonds
		{"DA", card.DA, true},
		{"D2", card.D2, true},
		{"D3", card.D3, true},
		{"D4", card.D4, true},
		{"D5", card.D5, true},
		{"D6", card.D6, true},
		{"D7", card.D7, true},
		{"D8", card.D8, true},
		{"D9", card.D9, true},
		{"DT", card.DT, true},
		{"DJ", card.DJ, true},
		{"DQ", card.DQ, true},
		{"DK", card.DK, true},
		// Hearts
		{"HA", card.HA, true},
		{"H2", card.H2, true},
		{"H3", card.H3, true},
		{"H4", card.H4, true},
		{"H5", card.H5, true},
		{"H6", card.H6, true},
		{"H7", card.H7, true},
		{"H8", card.H8, true},
		{"H9", card.H9, true},
		{"HT", card.HT, true},
		{"HJ", card.HJ, true},
		{"HQ", card.HQ, true},
		{"HK", card.HK, true},
		// Spades
		{"SA", card.SA, true},
		{"S2", card.S2, true},
		{"S3", card.S3, true},
		{"S4", card.S4, true},
		{"S5", card.S5, true},
		{"S6", card.S6, true},
		{"S7", card.S7, true},
		{"S8", card.S8, true},
		{"S9", card.S9, true},
		{"ST", card.ST, true},
		{"SJ", card.SJ, true},
		{"SQ", card.SQ, true},
		{"SK", card.SK, true},
		// invalid
		{"invalid Suited Card", card.Card(52), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.IsValid(); got != tt.want {
				t.Errorf("Card.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Name(t *testing.T) {
	tests := []struct {
		name string
		card card.Card
		want string
	}{
		// Clubs
		{"CA", card.CA, "Ace of Clubs"},
		{"C2", card.C2, "Two of Clubs"},
		{"C3", card.C3, "Three of Clubs"},
		{"C4", card.C4, "Four of Clubs"},
		{"C5", card.C5, "Five of Clubs"},
		{"C6", card.C6, "Six of Clubs"},
		{"C7", card.C7, "Seven of Clubs"},
		{"C8", card.C8, "Eight of Clubs"},
		{"C9", card.C9, "Nine of Clubs"},
		{"CT", card.CT, "Ten of Clubs"},
		{"CJ", card.CJ, "Jack of Clubs"},
		{"CQ", card.CQ, "Queen of Clubs"},
		{"CK", card.CK, "King of Clubs"},
		// Diamonds
		{"DA", card.DA, "Ace of Diamonds"},
		{"D2", card.D2, "Two of Diamonds"},
		{"D3", card.D3, "Three of Diamonds"},
		{"D4", card.D4, "Four of Diamonds"},
		{"D5", card.D5, "Five of Diamonds"},
		{"D6", card.D6, "Six of Diamonds"},
		{"D7", card.D7, "Seven of Diamonds"},
		{"D8", card.D8, "Eight of Diamonds"},
		{"D9", card.D9, "Nine of Diamonds"},
		{"DT", card.DT, "Ten of Diamonds"},
		{"DJ", card.DJ, "Jack of Diamonds"},
		{"DQ", card.DQ, "Queen of Diamonds"},
		{"DK", card.DK, "King of Diamonds"},
		// Hearts
		{"HA", card.HA, "Ace of Hearts"},
		{"H2", card.H2, "Two of Hearts"},
		{"H3", card.H3, "Three of Hearts"},
		{"H4", card.H4, "Four of Hearts"},
		{"H5", card.H5, "Five of Hearts"},
		{"H6", card.H6, "Six of Hearts"},
		{"H7", card.H7, "Seven of Hearts"},
		{"H8", card.H8, "Eight of Hearts"},
		{"H9", card.H9, "Nine of Hearts"},
		{"HT", card.HT, "Ten of Hearts"},
		{"HJ", card.HJ, "Jack of Hearts"},
		{"HQ", card.HQ, "Queen of Hearts"},
		{"HK", card.HK, "King of Hearts"},
		// Spades
		{"SA", card.SA, "Ace of Spades"},
		{"S2", card.S2, "Two of Spades"},
		{"S3", card.S3, "Three of Spades"},
		{"S4", card.S4, "Four of Spades"},
		{"S5", card.S5, "Five of Spades"},
		{"S6", card.S6, "Six of Spades"},
		{"S7", card.S7, "Seven of Spades"},
		{"S8", card.S8, "Eight of Spades"},
		{"S9", card.S9, "Nine of Spades"},
		{"ST", card.ST, "Ten of Spades"},
		{"SJ", card.SJ, "Jack of Spades"},
		{"SQ", card.SQ, "Queen of Spades"},
		{"SK", card.SK, "King of Spades"},
		// invalid
		{"invalid Card", card.Card(52), "Invalid Card(52)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.Name(); got != tt.want {
				t.Errorf("Card.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Rank(t *testing.T) {
	tests := []struct {
		name string
		card card.Card
		want card.Rank
	}{
		// Clubs
		{"CA", card.CA, card.Ace},
		{"C2", card.C2, card.Two},
		{"C3", card.C3, card.Three},
		{"C4", card.C4, card.Four},
		{"C5", card.C5, card.Five},
		{"C6", card.C6, card.Six},
		{"C7", card.C7, card.Seven},
		{"C8", card.C8, card.Eight},
		{"C9", card.C9, card.Nine},
		{"CT", card.CT, card.Ten},
		{"CJ", card.CJ, card.Jack},
		{"CQ", card.CQ, card.Queen},
		{"CK", card.CK, card.King},
		// Diamonds
		{"DA", card.DA, card.Ace},
		{"D2", card.D2, card.Two},
		{"D3", card.D3, card.Three},
		{"D4", card.D4, card.Four},
		{"D5", card.D5, card.Five},
		{"D6", card.D6, card.Six},
		{"D7", card.D7, card.Seven},
		{"D8", card.D8, card.Eight},
		{"D9", card.D9, card.Nine},
		{"DT", card.DT, card.Ten},
		{"DJ", card.DJ, card.Jack},
		{"DQ", card.DQ, card.Queen},
		{"DK", card.DK, card.King},
		// Hearts
		{"HA", card.HA, card.Ace},
		{"H2", card.H2, card.Two},
		{"H3", card.H3, card.Three},
		{"H4", card.H4, card.Four},
		{"H5", card.H5, card.Five},
		{"H6", card.H6, card.Six},
		{"H7", card.H7, card.Seven},
		{"H8", card.H8, card.Eight},
		{"H9", card.H9, card.Nine},
		{"HT", card.HT, card.Ten},
		{"HJ", card.HJ, card.Jack},
		{"HQ", card.HQ, card.Queen},
		{"HK", card.HK, card.King},
		// Spades
		{"SA", card.SA, card.Ace},
		{"S2", card.S2, card.Two},
		{"S3", card.S3, card.Three},
		{"S4", card.S4, card.Four},
		{"S5", card.S5, card.Five},
		{"S6", card.S6, card.Six},
		{"S7", card.S7, card.Seven},
		{"S8", card.S8, card.Eight},
		{"S9", card.S9, card.Nine},
		{"ST", card.ST, card.Ten},
		{"SJ", card.SJ, card.Jack},
		{"SQ", card.SQ, card.Queen},
		{"SK", card.SK, card.King},
		// invalid
		{"invalid Card", card.Card(52), card.Ace},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.Rank(); got != tt.want {
				t.Errorf("Card.Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		card card.Card
		want string
	}{
		// Clubs
		{"CA", card.CA, "CA"},
		{"C2", card.C2, "C2"},
		{"C3", card.C3, "C3"},
		{"C4", card.C4, "C4"},
		{"C5", card.C5, "C5"},
		{"C6", card.C6, "C6"},
		{"C7", card.C7, "C7"},
		{"C8", card.C8, "C8"},
		{"C9", card.C9, "C9"},
		{"CT", card.CT, "CT"},
		{"CJ", card.CJ, "CJ"},
		{"CQ", card.CQ, "CQ"},
		{"CK", card.CK, "CK"},
		// Diamonds
		{"DA", card.DA, "DA"},
		{"D2", card.D2, "D2"},
		{"D3", card.D3, "D3"},
		{"D4", card.D4, "D4"},
		{"D5", card.D5, "D5"},
		{"D6", card.D6, "D6"},
		{"D7", card.D7, "D7"},
		{"D8", card.D8, "D8"},
		{"D9", card.D9, "D9"},
		{"DT", card.DT, "DT"},
		{"DJ", card.DJ, "DJ"},
		{"DQ", card.DQ, "DQ"},
		{"DK", card.DK, "DK"},
		// Hearts
		{"HA", card.HA, "HA"},
		{"H2", card.H2, "H2"},
		{"H3", card.H3, "H3"},
		{"H4", card.H4, "H4"},
		{"H5", card.H5, "H5"},
		{"H6", card.H6, "H6"},
		{"H7", card.H7, "H7"},
		{"H8", card.H8, "H8"},
		{"H9", card.H9, "H9"},
		{"HT", card.HT, "HT"},
		{"HJ", card.HJ, "HJ"},
		{"HQ", card.HQ, "HQ"},
		{"HK", card.HK, "HK"},
		// Spades
		{"SA", card.SA, "SA"},
		{"S2", card.S2, "S2"},
		{"S3", card.S3, "S3"},
		{"S4", card.S4, "S4"},
		{"S5", card.S5, "S5"},
		{"S6", card.S6, "S6"},
		{"S7", card.S7, "S7"},
		{"S8", card.S8, "S8"},
		{"S9", card.S9, "S9"},
		{"ST", card.ST, "ST"},
		{"SJ", card.SJ, "SJ"},
		{"SQ", card.SQ, "SQ"},
		{"SK", card.SK, "SK"},
		// invalid
		{"invalid Card", card.Card(52), "!(4)A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.String(); got != tt.want {
				t.Errorf("Card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Suit(t *testing.T) {
	tests := []struct {
		name string
		card card.Card
		want card.Suit
	}{
		// Clubs
		{"CA", card.CA, card.Club},
		{"C2", card.C2, card.Club},
		{"C3", card.C3, card.Club},
		{"C4", card.C4, card.Club},
		{"C5", card.C5, card.Club},
		{"C6", card.C6, card.Club},
		{"C7", card.C7, card.Club},
		{"C8", card.C8, card.Club},
		{"C9", card.C9, card.Club},
		{"CT", card.CT, card.Club},
		{"CJ", card.CJ, card.Club},
		{"CQ", card.CQ, card.Club},
		{"CK", card.CK, card.Club},
		// Diamonds
		{"DA", card.DA, card.Diamond},
		{"D2", card.D2, card.Diamond},
		{"D3", card.D3, card.Diamond},
		{"D4", card.D4, card.Diamond},
		{"D5", card.D5, card.Diamond},
		{"D6", card.D6, card.Diamond},
		{"D7", card.D7, card.Diamond},
		{"D8", card.D8, card.Diamond},
		{"D9", card.D9, card.Diamond},
		{"DT", card.DT, card.Diamond},
		{"DJ", card.DJ, card.Diamond},
		{"DQ", card.DQ, card.Diamond},
		{"DK", card.DK, card.Diamond},
		// Hearts
		{"HA", card.HA, card.Heart},
		{"H2", card.H2, card.Heart},
		{"H3", card.H3, card.Heart},
		{"H4", card.H4, card.Heart},
		{"H5", card.H5, card.Heart},
		{"H6", card.H6, card.Heart},
		{"H7", card.H7, card.Heart},
		{"H8", card.H8, card.Heart},
		{"H9", card.H9, card.Heart},
		{"HT", card.HT, card.Heart},
		{"HJ", card.HJ, card.Heart},
		{"HQ", card.HQ, card.Heart},
		{"HK", card.HK, card.Heart},
		// Spades
		{"SA", card.SA, card.Spade},
		{"S2", card.S2, card.Spade},
		{"S3", card.S3, card.Spade},
		{"S4", card.S4, card.Spade},
		{"S5", card.S5, card.Spade},
		{"S6", card.S6, card.Spade},
		{"S7", card.S7, card.Spade},
		{"S8", card.S8, card.Spade},
		{"S9", card.S9, card.Spade},
		{"ST", card.ST, card.Spade},
		{"SJ", card.SJ, card.Spade},
		{"SQ", card.SQ, card.Spade},
		{"SK", card.SK, card.Spade},
		// invalid
		{"invalid Card", card.Card(52), 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.Suit(); got != tt.want {
				t.Errorf("Card.Suit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleWith() {
	sa := card.With(card.Spade, card.Ace)

	fmt.Println(sa)
	fmt.Println(sa == card.SA)
	// Output:
	// SA
	// true
}

func ExampleCard_Name() {
	fmt.Println(card.SA.Name())
	fmt.Println(card.D7.Name())
	// Output:
	// Ace of Spades
	// Seven of Diamonds
}

func ExampleCard_String() {
	fmt.Println(card.SuitSymbols)
	fmt.Println(card.SA)

	card.SuitSymbols = [4]string{
		"\u2663", // ♣
		"\u2662", // ♢
		"\u2661", // ♡
		"\u2660", // ♠
	}
	fmt.Println(card.SA)
	// Output:
	// [C D H S]
	// SA
	// ♠A
	card.SuitSymbols = [4]string{"C", "D", "H", "S"}
}

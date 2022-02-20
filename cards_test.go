package card_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/makabe/card"
)

func TestCards_Add(t *testing.T) {
	tests := []struct {
		cards card.Cards
		other card.Cards
		want  card.Cards
	}{
		{
			nil,
			card.Cards{},
			nil,
		},
		{
			nil,
			card.Cards{card.SA},
			card.Cards{card.SA},
		},
		{
			nil,
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.S2},
		},

		{
			card.Cards{},
			card.Cards{},
			card.Cards{},
		},
		{
			card.Cards{},
			card.Cards{card.SA},
			card.Cards{card.SA},
		},
		{
			card.Cards{},
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.S2},
		},

		{
			card.Cards{card.SA},
			card.Cards{},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.SA},
			card.Cards{card.SA},
			card.Cards{card.SA, card.SA},
		},
		{
			card.Cards{card.SA},
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.SA, card.S2},
		},

		{
			card.Cards{card.SA, card.S2},
			card.Cards{},
			card.Cards{card.SA, card.S2},
		},
		{
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA},
			card.Cards{card.SA, card.S2, card.SA},
		},
		{
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.S2, card.SA, card.S2},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cards, tt.other), func(t *testing.T) {
			clone := tt.cards.Clone()
			got := tt.cards.Add(tt.other...)
			// expected value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.Add() = %v, want %v", got, tt.want)
			}
			// non-destructive
			if !reflect.DeepEqual(clone, tt.cards) {
				t.Errorf("Cards.Add() is not non-destructive(%v)", tt.cards)
			}
			// not sharing underlying array
			if shareArray(got, tt.cards) {
				t.Error("Cards.Add() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Any(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		predName  predicateName
		want      bool
	}{
		{nilCards, isValid, false},
		{nilCards, is6Pips, false},
		{nilCards, isSpade, false},
		{nilCards, isHeart, false},

		{empty, isValid, false},
		{empty, is6Pips, false},
		{empty, isSpade, false},
		{empty, isHeart, false},

		{standardDeck, isValid, true},
		{standardDeck, is6Pips, true},
		{standardDeck, isSpade, true},
		{standardDeck, isHeart, true},

		{piquetPack, isValid, true},
		{piquetPack, is6Pips, false},
		{piquetPack, isSpade, true},
		{piquetPack, isHeart, true},

		{fourAces, isValid, true},
		{fourAces, is6Pips, false},
		{fourAces, isSpade, true},
		{fourAces, isHeart, true},

		{singleSpadeAce, isValid, true},
		{singleSpadeAce, is6Pips, false},
		{singleSpadeAce, isSpade, true},
		{singleSpadeAce, isHeart, false},

		{doubleSpadeAce, isValid, true},
		{doubleSpadeAce, is6Pips, false},
		{doubleSpadeAce, isSpade, true},
		{doubleSpadeAce, isHeart, false},

		{singleInvalidAce, isValid, false},
		{singleInvalidAce, is6Pips, false},
		{singleInvalidAce, isSpade, false},
		{singleInvalidAce, isHeart, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cardsName, tt.predName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			pred := predicate(tt.predName)
			if got := cs.Any(pred); got != tt.want {
				t.Errorf("Cards.Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCards_Bottom(t *testing.T) {
	sa := card.SA
	tests := []struct {
		cardsName cardsName
		want1     card.Cards
		want2     *card.Card
	}{
		{nilCards, nil, nil},
		{empty, card.Cards{}, nil},
		{fourAces, card.Cards{card.HA, card.CA, card.DA}, &sa},
		{singleSpadeAce, card.Cards{}, &sa},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.cardsName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			// expected values
			got1, got2 := cs.Bottom()
			if !reflect.DeepEqual(got1, tt.want1) || !equal(got2, tt.want2) {
				t.Errorf("Cards.Bottom() = (%v, %v), want (%v, %v)", got1, got2, tt.want1, tt.want2)
			}
			// non-destructive
			if !reflect.DeepEqual(cs, cards(tt.cardsName)) {
				t.Errorf("Cards.Bottom() is not non-destructive(%v)", cs)
			}
			// not sharing underlying array
			if shareArray(got1, cs) {
				t.Error("Cards.Bottom() returns Cards sharing underlying array")
			}
			// not returns a pointer to the underlying array
			for i := range cs {
				if got2 == &cs[i] {
					t.Error("Cards.Bottom() returns a pointer to the underlying array")
				}
			}
			for i := 0; i < cap(got1); i++ {
				if &got1[i] == got2 {
					t.Error("Cards.Bottom() returns a pointer to the other cards' underlying array")
				}
			}
		})
	}
}

func TestCards_Clone(t *testing.T) {
	tests := []cardsName{
		nilCards,
		empty,
		standardDeck,
		piquetPack,
		fourAces,
		singleSpadeAce,
		doubleSpadeAce,
		singleInvalidAce,
	}

	for _, cardsName := range tests {
		t.Run(fmt.Sprint(cardsName), func(t *testing.T) {
			cs := cards(cardsName)
			got := cs.Clone()
			// expected value
			if !reflect.DeepEqual(got, cs) {
				t.Errorf("Cards.Clone() = %v, want %v", got, cs)
			}
			// not sharing underlying array
			if shareArray(got, cs) {
				t.Error("Cards.Clone() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Empty(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		want      bool
	}{
		{nilCards, true},
		{empty, true},
		{standardDeck, false},
		{piquetPack, false},
		{fourAces, false},
		{singleSpadeAce, false},
		{doubleSpadeAce, false},
		{singleInvalidAce, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.cardsName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			if got := cs.Empty(); got != tt.want {
				t.Errorf("Cards.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCards_Every(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		predName  predicateName
		want      bool
	}{
		{nilCards, isValid, true},
		{nilCards, isAce, true},
		{nilCards, isSpade, true},
		{nilCards, isPiquetCard, true},

		{empty, isValid, true},
		{empty, isAce, true},
		{empty, isSpade, true},
		{empty, isPiquetCard, true},

		{standardDeck, isValid, true},
		{standardDeck, isAce, false},
		{standardDeck, isSpade, false},
		{standardDeck, isPiquetCard, false},

		{piquetPack, isValid, true},
		{piquetPack, isAce, false},
		{piquetPack, isSpade, false},
		{piquetPack, isPiquetCard, true},

		{fourAces, isValid, true},
		{fourAces, isAce, true},
		{fourAces, isSpade, false},
		{fourAces, isPiquetCard, true},

		{singleSpadeAce, isValid, true},
		{singleSpadeAce, isAce, true},
		{singleSpadeAce, isSpade, true},
		{singleSpadeAce, isPiquetCard, true},

		{doubleSpadeAce, isValid, true},
		{doubleSpadeAce, isAce, true},
		{doubleSpadeAce, isSpade, true},
		{doubleSpadeAce, isPiquetCard, true},

		{singleInvalidAce, isValid, false},
		{singleInvalidAce, isAce, true},
		{singleInvalidAce, isSpade, false},
		{singleInvalidAce, isPiquetCard, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cardsName, tt.predName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			pred := predicate(tt.predName)
			if got := cs.Every(pred); got != tt.want {
				t.Errorf("Cards.Every() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCards_Filter(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		predName  predicateName
		want      card.Cards
	}{
		{nilCards, isValid, nil},
		{nilCards, isAce, nil},
		{nilCards, is6Pips, nil},
		{nilCards, isSpade, nil},
		{nilCards, isHeart, nil},
		{nilCards, isRed, nil},
		{nilCards, isPiquetCard, nil},

		{empty, isValid, card.Cards{}},
		{empty, isAce, card.Cards{}},
		{empty, is6Pips, card.Cards{}},
		{empty, isSpade, card.Cards{}},
		{empty, isHeart, card.Cards{}},
		{empty, isRed, card.Cards{}},
		{empty, isPiquetCard, card.Cards{}},

		{standardDeck, isValid, cards(standardDeck)},
		{standardDeck, isAce, card.Cards{card.HA, card.CA, card.DA, card.SA}},
		{standardDeck, is6Pips, card.Cards{card.H6, card.C6, card.D6, card.S6}},
		{standardDeck, isSpade, card.Cards{
			card.SK, card.SQ, card.SJ, card.ST, card.S9, card.S8, card.S7,
			card.S6, card.S5, card.S4, card.S3, card.S2, card.SA,
		}},
		{standardDeck, isHeart, card.Cards{
			card.HA, card.H2, card.H3, card.H4, card.H5, card.H6, card.H7,
			card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
		}},
		{standardDeck, isRed, card.Cards{
			card.HA, card.H2, card.H3, card.H4, card.H5, card.H6, card.H7,
			card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
			card.DK, card.DQ, card.DJ, card.DT, card.D9, card.D8, card.D7,
			card.D6, card.D5, card.D4, card.D3, card.D2, card.DA,
		}},
		{standardDeck, isPiquetCard, cards(piquetPack)},

		{piquetPack, isValid, cards(piquetPack)},
		{piquetPack, isAce, card.Cards{card.HA, card.CA, card.DA, card.SA}},
		{piquetPack, is6Pips, card.Cards{}},
		{piquetPack, isSpade, card.Cards{
			card.SK, card.SQ, card.SJ, card.ST, card.S9, card.S8, card.S7, card.SA,
		}},
		{piquetPack, isHeart, card.Cards{
			card.HA, card.H7, card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
		}},
		{piquetPack, isRed, card.Cards{
			card.HA, card.H7, card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
			card.DK, card.DQ, card.DJ, card.DT, card.D9, card.D8, card.D7, card.DA,
		}},
		{piquetPack, isPiquetCard, cards(piquetPack)},

		{fourAces, isValid, cards(fourAces)},
		{fourAces, isAce, cards(fourAces)},
		{fourAces, is6Pips, card.Cards{}},
		{fourAces, isSpade, card.Cards{card.SA}},
		{fourAces, isHeart, card.Cards{card.HA}},
		{fourAces, isRed, card.Cards{card.HA, card.DA}},
		{fourAces, isPiquetCard, cards(fourAces)},

		{singleSpadeAce, isValid, card.Cards{card.SA}},
		{singleSpadeAce, isAce, card.Cards{card.SA}},
		{singleSpadeAce, is6Pips, card.Cards{}},
		{singleSpadeAce, isSpade, card.Cards{card.SA}},
		{singleSpadeAce, isHeart, card.Cards{}},
		{singleSpadeAce, isRed, card.Cards{}},
		{singleSpadeAce, isPiquetCard, card.Cards{card.SA}},

		{doubleSpadeAce, isValid, card.Cards{card.SA, card.SA}},
		{doubleSpadeAce, isAce, card.Cards{card.SA, card.SA}},
		{doubleSpadeAce, is6Pips, card.Cards{}},
		{doubleSpadeAce, isSpade, card.Cards{card.SA, card.SA}},
		{doubleSpadeAce, isHeart, card.Cards{}},
		{doubleSpadeAce, isRed, card.Cards{}},
		{doubleSpadeAce, isPiquetCard, card.Cards{card.SA, card.SA}},

		{singleInvalidAce, isValid, card.Cards{}},
		{singleInvalidAce, isAce, cards(singleInvalidAce)},
		{singleInvalidAce, is6Pips, card.Cards{}},
		{singleInvalidAce, isSpade, card.Cards{}},
		{singleInvalidAce, isHeart, card.Cards{}},
		{singleInvalidAce, isRed, card.Cards{}},
		{singleInvalidAce, isPiquetCard, card.Cards{}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cardsName, tt.predName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			pred := predicate(tt.predName)
			got := cs.Filter(pred)
			// expected value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.Filter() = %v, want %v", got, tt.want)
			}
			// non-destructive
			if !reflect.DeepEqual(cs, cards(tt.cardsName)) {
				t.Errorf("Cards.Filter() is not non-destructive(%v)", cs)
			}
			// not sharing underlying array
			if shareArray(got, cs) {
				t.Error("Cards.Filter() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Include(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		otherName cardsName
		want      bool
	}{
		{nilCards, empty, true},
		{nilCards, standardDeck, false},
		{nilCards, piquetPack, false},
		{nilCards, fourAces, false},
		{nilCards, singleSpadeAce, false},
		{nilCards, doubleSpadeAce, false},
		{nilCards, singleInvalidAce, false},

		{empty, empty, true},
		{empty, standardDeck, false},
		{empty, piquetPack, false},
		{empty, fourAces, false},
		{empty, singleSpadeAce, false},
		{empty, doubleSpadeAce, false},
		{empty, singleInvalidAce, false},

		{standardDeck, empty, true},
		{standardDeck, standardDeck, true},
		{standardDeck, piquetPack, true},
		{standardDeck, fourAces, true},
		{standardDeck, singleSpadeAce, true},
		// duplicate cards are also considered as one card
		{standardDeck, doubleSpadeAce, true},
		{standardDeck, singleInvalidAce, false},

		{piquetPack, empty, true},
		{piquetPack, standardDeck, false},
		{piquetPack, piquetPack, true},
		{piquetPack, fourAces, true},
		{piquetPack, singleSpadeAce, true},
		{piquetPack, doubleSpadeAce, true}, // considered as one SA
		{piquetPack, singleInvalidAce, false},

		{fourAces, empty, true},
		{fourAces, standardDeck, false},
		{fourAces, piquetPack, false},
		{fourAces, fourAces, true},
		{fourAces, singleSpadeAce, true},
		{fourAces, doubleSpadeAce, true}, // considered as one SA
		{fourAces, singleInvalidAce, false},

		{singleSpadeAce, empty, true},
		{singleSpadeAce, standardDeck, false},
		{singleSpadeAce, piquetPack, false},
		{singleSpadeAce, fourAces, false},
		{singleSpadeAce, singleSpadeAce, true},
		{singleSpadeAce, doubleSpadeAce, true}, // considered as one SA
		{singleSpadeAce, singleInvalidAce, false},

		{doubleSpadeAce, empty, true},
		{doubleSpadeAce, standardDeck, false},
		{doubleSpadeAce, piquetPack, false},
		{doubleSpadeAce, fourAces, false},
		{doubleSpadeAce, singleSpadeAce, true},
		{doubleSpadeAce, doubleSpadeAce, true}, // considered as one SA
		{doubleSpadeAce, singleInvalidAce, false},

		{singleInvalidAce, empty, true},
		{singleInvalidAce, standardDeck, false},
		{singleInvalidAce, piquetPack, false},
		{singleInvalidAce, fourAces, false},
		{singleInvalidAce, singleSpadeAce, false},
		{singleInvalidAce, doubleSpadeAce, false},
		{singleInvalidAce, singleInvalidAce, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cardsName, tt.otherName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			other := cards(tt.otherName)
			if got := cs.Include(other...); got != tt.want {
				t.Errorf("Cards.Include() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCards_Move(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		otherName cardsName
		n         uint
		want1     card.Cards
		want2     card.Cards
	}{
		{nilCards, nilCards, 0, nil, nil},
		{nilCards, nilCards, 1, nil, nil},
		{nilCards, empty, 0, card.Cards{}, nil},
		{nilCards, empty, 1, card.Cards{}, nil},
		{nilCards, fourAces, 0, card.Cards{card.HA, card.CA, card.DA, card.SA}, nil},
		{nilCards, fourAces, 1, card.Cards{card.HA, card.CA, card.DA, card.SA}, nil},

		{empty, nilCards, 0, nil, card.Cards{}},
		{empty, nilCards, 1, nil, card.Cards{}},
		{empty, empty, 0, card.Cards{}, card.Cards{}},
		{empty, empty, 1, card.Cards{}, card.Cards{}},
		{empty, fourAces, 0, card.Cards{card.HA, card.CA, card.DA, card.SA}, card.Cards{}},
		{empty, fourAces, 1, card.Cards{card.HA, card.CA, card.DA, card.SA}, card.Cards{}},

		{fourAces, nilCards, 0, nil, card.Cards{card.HA, card.CA, card.DA, card.SA}},
		{fourAces, nilCards, 1, card.Cards{card.HA}, card.Cards{card.CA, card.DA, card.SA}},
		{fourAces, nilCards, 4, card.Cards{card.HA, card.CA, card.DA, card.SA}, card.Cards{}},
		{fourAces, nilCards, 5, card.Cards{card.HA, card.CA, card.DA, card.SA}, card.Cards{}},
		{fourAces, empty, 0, card.Cards{}, card.Cards{card.HA, card.CA, card.DA, card.SA}},
		{fourAces, empty, 1, card.Cards{card.HA}, card.Cards{card.CA, card.DA, card.SA}},
		{fourAces, empty, 4, card.Cards{card.HA, card.CA, card.DA, card.SA}, card.Cards{}},
		{fourAces, empty, 5, card.Cards{card.HA, card.CA, card.DA, card.SA}, card.Cards{}},
		{
			fourAces, fourAces, 0,
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			card.Cards{card.HA, card.CA, card.DA, card.SA},
		},
		{
			fourAces, fourAces, 1,
			card.Cards{card.HA, card.CA, card.DA, card.SA, card.HA},
			card.Cards{card.CA, card.DA, card.SA},
		},
		{
			fourAces, fourAces, 4,
			card.Cards{card.HA, card.CA, card.DA, card.SA, card.HA, card.CA, card.DA, card.SA},
			card.Cards{},
		},
		{
			fourAces, fourAces, 5,
			card.Cards{card.HA, card.CA, card.DA, card.SA, card.HA, card.CA, card.DA, card.SA},
			card.Cards{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d,%s", tt.cardsName, tt.n, tt.otherName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			to := cards(tt.otherName)
			got1, got2 := cs.Move(tt.n, to)
			// expected values
			if !reflect.DeepEqual(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Cards.Move() = (%v, %v), want (%v, %v)", got1, got2, tt.want1, tt.want2)
			}
			// non-destructive
			if !reflect.DeepEqual(cs, cards(tt.cardsName)) {
				t.Errorf("Cards.Move() is not non-destructive(%v)", cs)
			}
			// not sharing underlying array
			if shareArray(got1, cs) || shareArray(got2, cs) || shareArray(got1, got2) {
				t.Error("Cards.Move() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Partition(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		predName  predicateName
		want1     card.Cards
		want2     card.Cards
	}{
		{nilCards, isAce, nil, nil},
		{nilCards, isSpade, nil, nil},
		{nilCards, isRed, nil, nil},
		{nilCards, isPiquetCard, nil, nil},

		{empty, isAce, card.Cards{}, card.Cards{}},
		{empty, isSpade, card.Cards{}, card.Cards{}},
		{empty, isRed, card.Cards{}, card.Cards{}},
		{empty, isPiquetCard, card.Cards{}, card.Cards{}},

		{
			standardDeck,
			isAce,
			card.Cards{
				card.HA, card.CA, card.DA, card.SA,
			},
			card.Cards{
				card.H2, card.H3, card.H4, card.H5, card.H6, card.H7,
				card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
				card.C2, card.C3, card.C4, card.C5, card.C6, card.C7,
				card.C8, card.C9, card.CT, card.CJ, card.CQ, card.CK,
				card.DK, card.DQ, card.DJ, card.DT, card.D9, card.D8,
				card.D7, card.D6, card.D5, card.D4, card.D3, card.D2,
				card.SK, card.SQ, card.SJ, card.ST, card.S9, card.S8,
				card.S7, card.S6, card.S5, card.S4, card.S3, card.S2,
			},
		},
		{
			standardDeck,
			isSpade,
			card.Cards{
				card.SK, card.SQ, card.SJ, card.ST, card.S9, card.S8, card.S7,
				card.S6, card.S5, card.S4, card.S3, card.S2, card.SA,
			},
			card.Cards{
				card.HA, card.H2, card.H3, card.H4, card.H5, card.H6, card.H7,
				card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
				card.CA, card.C2, card.C3, card.C4, card.C5, card.C6, card.C7,
				card.C8, card.C9, card.CT, card.CJ, card.CQ, card.CK,
				card.DK, card.DQ, card.DJ, card.DT, card.D9, card.D8, card.D7,
				card.D6, card.D5, card.D4, card.D3, card.D2, card.DA,
			},
		},
		{
			standardDeck,
			isRed,
			card.Cards{
				card.HA, card.H2, card.H3, card.H4, card.H5, card.H6, card.H7,
				card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
				card.DK, card.DQ, card.DJ, card.DT, card.D9, card.D8, card.D7,
				card.D6, card.D5, card.D4, card.D3, card.D2, card.DA,
			},
			card.Cards{
				card.CA, card.C2, card.C3, card.C4, card.C5, card.C6, card.C7,
				card.C8, card.C9, card.CT, card.CJ, card.CQ, card.CK,
				card.SK, card.SQ, card.SJ, card.ST, card.S9, card.S8, card.S7,
				card.S6, card.S5, card.S4, card.S3, card.S2, card.SA,
			},
		},
		{
			standardDeck,
			isPiquetCard,
			card.Cards{
				card.HA, card.H7, card.H8, card.H9, card.HT, card.HJ, card.HQ, card.HK,
				card.CA, card.C7, card.C8, card.C9, card.CT, card.CJ, card.CQ, card.CK,
				card.DK, card.DQ, card.DJ, card.DT, card.D9, card.D8, card.D7, card.DA,
				card.SK, card.SQ, card.SJ, card.ST, card.S9, card.S8, card.S7, card.SA,
			},
			card.Cards{
				card.H2, card.H3, card.H4, card.H5, card.H6,
				card.C2, card.C3, card.C4, card.C5, card.C6,
				card.D6, card.D5, card.D4, card.D3, card.D2,
				card.S6, card.S5, card.S4, card.S3, card.S2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cardsName, tt.predName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			pred := predicate(tt.predName)
			got1, got2 := cs.Partition(pred)
			// expected values
			if !reflect.DeepEqual(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Cards.Partition() = (%v, %v), want (%v, %v)", got1, got2, tt.want1, tt.want2)
			}
			// non-destructive
			if !reflect.DeepEqual(cs, cards(tt.cardsName)) {
				t.Errorf("Cards.Partition() is not non-destructive(%v)", cs)
			}
			// not sharing underlying array
			if shareArray(got1, cs) || shareArray(got2, cs) || shareArray(got1, got2) {
				t.Error("Cards.Partition() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Remove(t *testing.T) {
	tests := []struct {
		cards card.Cards
		other card.Cards
		want  card.Cards
	}{
		{
			nil,
			card.Cards{},
			nil,
		},
		{
			nil,
			card.Cards{card.SA},
			nil,
		},
		{
			nil,
			card.Cards{card.SA, card.S2},
			nil,
		},
		{
			nil,
			card.Cards{card.SA, card.SA}, // considered as one SA
			nil,
		},

		{
			card.Cards{},
			card.Cards{},
			card.Cards{},
		},
		{
			card.Cards{},
			card.Cards{card.SA},
			card.Cards{},
		},
		{
			card.Cards{},
			card.Cards{card.SA, card.S2},
			card.Cards{},
		},
		{
			card.Cards{},
			card.Cards{card.SA, card.SA}, // considered as one SA
			card.Cards{},
		},

		{
			card.Cards{card.SA},
			card.Cards{},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.SA},
			card.Cards{card.SA},
			card.Cards{},
		},
		{
			card.Cards{card.SA},
			card.Cards{card.SA, card.S2},
			card.Cards{},
		},
		{
			card.Cards{card.SA},
			card.Cards{card.SA, card.SA}, // considered as one SA
			card.Cards{},
		},

		{
			card.Cards{card.SA, card.S2},
			card.Cards{},
			card.Cards{card.SA, card.S2},
		},
		{
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA},
			card.Cards{card.S2},
		},
		{
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.S2},
			card.Cards{},
		},
		{
			card.Cards{card.SA, card.S2},
			card.Cards{card.SA, card.SA}, // considered as one SA
			card.Cards{card.S2},
		},

		{
			card.Cards{card.SA, card.SA},
			card.Cards{},
			card.Cards{card.SA, card.SA},
		},
		{
			card.Cards{card.SA, card.SA},
			card.Cards{card.SA},
			card.Cards{}, // Both SA will be removed
		},
		{
			card.Cards{card.SA, card.SA},
			card.Cards{card.SA, card.S2},
			card.Cards{},
		},
		{
			card.Cards{card.SA, card.SA},
			card.Cards{card.SA, card.SA}, // considered as one SA
			card.Cards{},                 // Both SA will be removed
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cards, tt.other), func(t *testing.T) {
			clone := tt.cards.Clone()
			got := tt.cards.Remove(tt.other...)
			// expected value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.Remove() = %v, want %v", got, tt.want)
			}
			// non-destructive
			if !reflect.DeepEqual(clone, tt.cards) {
				t.Errorf("Cards.Remove() is not non-destructive(%v)", tt.cards)
			}
			// not sharing underlying array
			if shareArray(got, tt.cards) {
				t.Error("Cards.Remove() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Reverse(t *testing.T) {
	tests := []struct {
		cards card.Cards
		want  card.Cards
	}{
		{
			nil,
			nil,
		},
		{
			card.Cards{},
			card.Cards{},
		},
		{
			card.Cards{card.SA},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.CA, card.C3, card.HA, card.H2, card.SA},
			card.Cards{card.SA, card.H2, card.HA, card.C3, card.CA},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.cards), func(t *testing.T) {
			clone := tt.cards.Clone()
			got := tt.cards.Reverse()
			// expected value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.Reverse() = %v, want %v", tt.cards, tt.want)
			}
			// non-destructive
			if !reflect.DeepEqual(clone, tt.cards) {
				t.Errorf("Cards.Reverse() is not non-destructive(%v)", tt.cards)
			}
			// not sharing underlying array
			if shareArray(got, tt.cards) {
				t.Error("Cards.Reverse() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Shuffle(t *testing.T) {
	stdDeck := card.NewStandardDeck()
	clone := stdDeck.Clone()

	t.Run("should not be destructive", func(t *testing.T) {
		r := rand.New(rand.NewSource(1))
		_ = stdDeck.Shuffle(r)
		if !reflect.DeepEqual(clone, stdDeck) {
			t.Fatalf("Cards.Shuffle() is not non-destructive(%v)", stdDeck)
		}
	})

	t.Run("returns new card", func(t *testing.T) {
		r := rand.New(rand.NewSource(1))
		got := stdDeck.Shuffle(r)
		if shareArray(got, stdDeck) {
			t.Error("Cards.Shuffle() returns Cards sharing underlying array")
		}
	})

	t.Run("returns shuffled card", func(t *testing.T) {
		r := rand.New(rand.NewSource(1))
		got := stdDeck.Shuffle(r)
		if reflect.DeepEqual(got, stdDeck) {
			t.Errorf("Cards.Shuffle() is not working well (%v)", got)
		}
	})

	t.Run("returns card in a different order each time it is called", func(t *testing.T) {
		r := rand.New(rand.NewSource(1))
		got1 := stdDeck.Shuffle(r)
		got2 := stdDeck.Shuffle(r)
		if reflect.DeepEqual(got1, got2) {
			t.Errorf("Cards.Shuffle() arranges in a fixed order each time (%v)", got1)
		}
	})

	t.Run("returns the card in the same order each time if the seed is fixed", func(t *testing.T) {
		r1 := rand.New(rand.NewSource(1))
		r2 := rand.New(rand.NewSource(1))
		got1 := stdDeck.Shuffle(r1)
		got2 := stdDeck.Shuffle(r2)
		if !reflect.DeepEqual(got1, got2) {
			t.Errorf("Cards.Shuffle() returns the card in a different order even if the seed is fixed (%v)", got1)
		}
	})

	t.Run("uses the default source if nil is passed for random source", func(t *testing.T) {
		got1 := stdDeck.Shuffle(nil)
		got2 := stdDeck.Shuffle(nil)
		if reflect.DeepEqual(got1, stdDeck) {
			t.Errorf("Cards.Shuffle() is not working well (%v)", got1)
		}
		if reflect.DeepEqual(got1, got2) {
			t.Errorf("Cards.Shuffle() arranges in a fixed order each time (%v)", got2)
		}
	})

	t.Run("returns nil when called on nil", func(t *testing.T) {
		var nilCards card.Cards
		r := rand.New(rand.NewSource(1))
		got := nilCards.Shuffle(r)
		if got != nil {
			t.Errorf("Cards.Shuffle() = %v, want nil", got)
		}
	})
}

func TestCards_Size(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		want      uint
	}{
		{nilCards, 0},
		{empty, 0},
		{standardDeck, 52},
		{piquetPack, 32},
		{fourAces, 4},
		{singleSpadeAce, 1},
		{doubleSpadeAce, 2},
		{singleInvalidAce, 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.cardsName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			if got := cs.Size(); got != tt.want {
				t.Errorf("Cards.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCards_Sort(t *testing.T) {
	cs := card.Cards{
		card.SA,
		card.SK,
		card.CA,
		card.H2,
		card.HA,
		card.C3,
	}

	tests := []struct {
		cards    card.Cards
		compName comparatorName
		want     card.Cards
	}{
		{cs, suitAscRankAsc, card.Cards{card.CA, card.C3, card.HA, card.H2, card.SA, card.SK}},
		{cs, rankAscSuitAsc, card.Cards{card.CA, card.HA, card.SA, card.H2, card.C3, card.SK}},
		{nil, suitAscRankAsc, nil},
		{card.Cards{}, suitAscRankAsc, card.Cards{}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.cards, tt.compName), func(t *testing.T) {
			comp := comparator(tt.compName)
			clone := tt.cards.Clone()
			got := tt.cards.Sort(comp)
			// expected value
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cards.Sort() = %v, want %v", got, tt.want)
			}
			// non-destructive
			if !reflect.DeepEqual(clone, tt.cards) {
				t.Errorf("Cards.Sort() is not non-destructive(%v)", tt.cards)
			}
			// not sharing underlying array
			if shareArray(got, cs) {
				t.Error("Cards.Sort() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Take(t *testing.T) {
	tests := []struct {
		cards card.Cards
		n     uint
		want1 card.Cards
		want2 card.Cards
	}{
		{
			nil,
			0,
			nil,
			nil,
		},
		{
			nil,
			1,
			nil,
			nil,
		},

		{
			card.Cards{},
			0,
			card.Cards{},
			card.Cards{},
		},
		{
			card.Cards{},
			1,
			card.Cards{},
			card.Cards{},
		},

		{
			card.Cards{card.SA},
			0,
			card.Cards{},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.SA},
			1,
			card.Cards{card.SA},
			card.Cards{},
		},
		{
			card.Cards{card.SA},
			2,
			card.Cards{card.SA},
			card.Cards{},
		},

		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			0,
			card.Cards{},
			card.Cards{card.HA, card.CA, card.DA, card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			3,
			card.Cards{card.HA, card.CA, card.DA},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			4,
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			card.Cards{},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			5,
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			card.Cards{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d", tt.cards, tt.n), func(t *testing.T) {
			clone := tt.cards.Clone()
			got1, got2 := tt.cards.Take(tt.n)
			// expected values
			if !reflect.DeepEqual(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Cards.Take() = (%v, %v), want (%v, %v)", got1, got2, tt.want1, tt.want2)
			}
			// non-destructive
			if !reflect.DeepEqual(clone, tt.cards) {
				t.Errorf("Cards.Take() is not non-destructive(%v)", tt.cards)
			}
			// not sharing underlying array
			if shareArray(got1, tt.cards) || shareArray(got2, tt.cards) || shareArray(got1, got2) {
				t.Error("Cards.Take() returns Cards sharing underlying array")
			}
		})
	}
}

func TestCards_Top(t *testing.T) {
	ha, sa := card.HA, card.SA
	tests := []struct {
		cardsName cardsName
		want1     *card.Card
		want2     card.Cards
	}{
		{nilCards, nil, nil},
		{empty, nil, card.Cards{}},
		{fourAces, &ha, card.Cards{card.CA, card.DA, card.SA}},
		{singleSpadeAce, &sa, card.Cards{}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.cardsName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			clone := cs.Clone()
			got1, got2 := cs.Top()
			// expected values
			if !equal(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Cards.Top() = (%v, %v), want (%v, %v)", got1, got2, tt.want1, tt.want2)
			}
			// non-destructive
			if !reflect.DeepEqual(cs, clone) {
				t.Errorf("Cards.Top() is not non-destructive(%v)", cs)
			}
			// not returns a pointer to the underlying array
			for i := 0; i < cap(cs); i++ {
				if got1 == &cs[i] {
					t.Error("Cards.Top() returns a pointer to the underlying array")
				}
			}
			// not sharing underlying array
			if shareArray(got2, cs) {
				t.Error("Cards.Top() returns Cards sharing underlying array")
			}
			for i := 0; i < cap(got2); i++ {
				if got1 == &got2[i] {
					t.Error("Cards.Top() returns a pointer to the other cards' underlying array")
				}
			}
		})
	}
}

func TestCards_haveDuplicates(t *testing.T) {
	tests := []struct {
		cardsName cardsName
		want      bool
	}{
		{nilCards, false},
		{empty, false},
		{standardDeck, false},
		{piquetPack, false},
		{fourAces, false},
		{singleSpadeAce, false},
		{doubleSpadeAce, true},
		{singleInvalidAce, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.cardsName), func(t *testing.T) {
			cs := cards(tt.cardsName)
			got := card.ExportCardsHaveDuplicates(cs) // Cards.haveDuplicates
			if got != tt.want {
				t.Errorf("Cards.haveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCards_splitAt(t *testing.T) {
	tests := []struct {
		cards card.Cards
		n     int
		want1 card.Cards
		want2 card.Cards
	}{
		{
			nil,
			0,
			nil,
			nil,
		},
		{
			nil,
			1,
			nil,
			nil,
		},
		{
			nil,
			-1,
			nil,
			nil,
		},

		{
			card.Cards{},
			0,
			card.Cards{},
			card.Cards{},
		},
		{
			card.Cards{},
			1,
			card.Cards{},
			card.Cards{},
		},
		{
			card.Cards{},
			-1,
			card.Cards{},
			card.Cards{},
		},

		{
			card.Cards{card.SA},
			0,
			card.Cards{},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.SA},
			1,
			card.Cards{card.SA},
			card.Cards{},
		},
		{
			card.Cards{card.SA},
			2,
			card.Cards{card.SA},
			card.Cards{},
		},
		{
			card.Cards{card.SA},
			-1,
			card.Cards{},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.SA},
			-2,
			card.Cards{},
			card.Cards{card.SA},
		},

		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			0,
			card.Cards{},
			card.Cards{card.HA, card.CA, card.DA, card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			3,
			card.Cards{card.HA, card.CA, card.DA},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			4,
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			card.Cards{},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			5,
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			card.Cards{},
		},

		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			-1,
			card.Cards{card.HA, card.CA, card.DA},
			card.Cards{card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			-3,
			card.Cards{card.HA},
			card.Cards{card.CA, card.DA, card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			-4,
			card.Cards{},
			card.Cards{card.HA, card.CA, card.DA, card.SA},
		},
		{
			card.Cards{card.HA, card.CA, card.DA, card.SA},
			-5,
			card.Cards{},
			card.Cards{card.HA, card.CA, card.DA, card.SA},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%d", tt.cards, tt.n), func(t *testing.T) {
			clone := tt.cards.Clone()
			got1, got2 := card.ExportCardsSplitAt(tt.cards, tt.n) // Cards.splitAt
			// expected values
			if !reflect.DeepEqual(got1, tt.want1) || !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Cards.splitAt() = (%v, %v), want (%v, %v)", got1, got2, tt.want1, tt.want2)
			}
			// non-destructive
			if !reflect.DeepEqual(clone, tt.cards) {
				t.Errorf("Cards.splitAt() is not non-destructive(%v)", tt.cards)
			}
			// not sharing underlying array
			if shareArray(got1, tt.cards) || shareArray(got2, tt.cards) || shareArray(got1, got2) {
				t.Error("Cards.splitAt() returns Cards sharing underlying array")
			}
		})
	}
}

package card_test

import (
	"fmt"
	"math/rand"

	"github.com/makabe/card"
)

func ExampleCards_Add() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Add())
	fmt.Println(hand.Add(card.S2))
	fmt.Println(hand.Add(card.S2, card.SA))
	// Output:
	// [SA SK SQ SJ ST]
	// [SA SK SQ SJ ST S2]
	// [SA SK SQ SJ ST S2 SA]
}

func ExampleCards_Add_nilEmpty() {
	nilCards := card.Cards(nil)
	emptyCards := card.Cards{}

	// returns a clone if no card are added
	added := nilCards.Add()
	fmt.Println(added, added == nil) // nil returns nil

	added = emptyCards.Add()
	fmt.Println(added, added == nil) // empty returns empty

	// returns non-empty cards if at least one card is added
	added = nilCards.Add(card.SA)
	fmt.Println(added)

	added = emptyCards.Add(card.SA)
	fmt.Println(added)
	// Output:
	// [] true
	// [] false
	// [SA]
	// [SA]
}

func ExampleCards_Add_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	added := hand.Add(card.S2)

	fmt.Println(added)
	fmt.Println(hand) // not mutated

	// the return value is a newly allocated Cards
	added[0] = card.S3

	fmt.Println(added) // mutated
	fmt.Println(hand)  // not affected
	// Output:
	// [SA SK SQ SJ ST S2]
	// [SA SK SQ SJ ST]
	// [S3 SK SQ SJ ST S2]
	// [SA SK SQ SJ ST]
}

func ExampleCards_Any() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }
	isQueen := func(c card.Card) bool { return c.Rank() == card.Queen }

	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ, // Queen!
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Any(isHeart))
	fmt.Println(hand.Any(isQueen))
	// Output:
	// false
	// true
}

func ExampleCards_Any_nilEmpty() {
	isCard := func(c card.Card) bool { return true }

	nilCards := card.Cards(nil)
	fmt.Println(nilCards.Any(isCard))

	emptyCards := card.Cards{}
	fmt.Println(emptyCards.Any(isCard))
	// Output:
	// false
	// false
}

func ExampleCards_Bottom() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Bottom())
	// Output:
	// [SA SK SQ SJ] ST
}

func ExampleCards_Bottom_nilEmpty() {
	nilCards := card.Cards(nil)
	remaining, bottom := nilCards.Bottom()

	fmt.Println(remaining, bottom, remaining == nil) // nil returns (nil, nil)

	emptyCards := card.Cards{}
	remaining, bottom = emptyCards.Bottom()

	fmt.Println(remaining, bottom, remaining == nil) // empty returns (empty, nil)
	// Output:
	// [] <nil> true
	// [] <nil> false
}

func ExampleCards_Bottom_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	remaining, bottom := hand.Bottom()

	fmt.Println(remaining, bottom)
	fmt.Println(hand) // not mutated

	// the return values share no addresses with the Cards
	remaining[0] = card.S2
	*bottom = card.S3

	fmt.Println(remaining, bottom) // mutated
	fmt.Println(hand)              // not affected

	// both the return values are independent of each other
	fmt.Println(cap(remaining) == len(remaining)) // no hidden capacity of the underlying array
	// Output:
	// [SA SK SQ SJ] ST
	// [SA SK SQ SJ ST]
	// [S2 SK SQ SJ] S3
	// [SA SK SQ SJ ST]
	// true
}

func ExampleCards_Clone() {
	fmt.Println(card.Cards{card.SA})
	fmt.Println(card.Cards{card.SA, card.S2})
	// Output:
	// [SA]
	// [SA S2]
}

func ExampleCards_Clone_nilEmpty() {
	nilCards := card.Cards(nil)
	clone := nilCards.Clone()

	fmt.Println(clone, clone == nil) // nil returns nil

	emptyCards := card.Cards{}
	clone = emptyCards.Clone()

	fmt.Println(clone, clone == nil) // empty returns empty
	// Output:
	// [] true
	// [] false
}

func ExampleCards_Clone_nonDestructive() {
	original := card.Cards{card.SA, card.S2}

	clone := original.Clone()

	fmt.Println(clone)

	// the return value is a newly allocated Cards
	clone[0] = card.S3

	fmt.Println(clone)    // mutated
	fmt.Println(original) // not affected
	// Output:
	// [SA S2]
	// [S3 S2]
	// [SA S2]
}

func ExampleCards_Empty() {
	p := func(cs card.Cards) {
		fmt.Println(cs.Empty())
	}

	p(nil)
	p(card.Cards{})
	p(card.Cards{card.SA})
	// Output:
	// true
	// true
	// false
}

func ExampleCards_Every() {
	isSpade := func(c card.Card) bool { return c.Suit() == card.Spade }
	isQueen := func(c card.Card) bool { return c.Rank() == card.Queen }

	hand := card.Cards{
		card.SA, // Spade!
		card.SK, // Spade!
		card.SQ, // Spade! Queen!
		card.SJ, // Spade!
		card.ST, // Spade!
	}

	fmt.Println(hand.Every(isSpade))
	fmt.Println(hand.Every(isQueen))
	// Output:
	// true
	// false
}

func ExampleCards_Every_nilEmpty() {
	isNotCard := func(c card.Card) bool { return false }

	nilCards := card.Cards(nil)
	fmt.Println(nilCards.Every(isNotCard))

	emptyCards := card.Cards{}
	fmt.Println(emptyCards.Every(isNotCard))
	// Output:
	// true
	// true
}

func ExampleCards_Filter() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }
	isFaceCard := func(c card.Card) bool { return c.Rank() >= card.Jack }

	hand := card.Cards{
		card.HA, // Heart!
		card.CK, //        FaceCard!
		card.HQ, // Heart! FaceCard!
		card.SJ, //        FaceCard!
		card.ST,
	}

	fmt.Println(hand.Filter(isHeart))
	fmt.Println(hand.Filter(isFaceCard))
	// Output:
	// [HA HQ]
	// [CK HQ SJ]
}

func ExampleCards_Filter_nilEmpty() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }

	nilCards := card.Cards(nil)
	hearts := nilCards.Filter(isHeart)

	fmt.Println(hearts, hearts == nil) // nil returns nil

	emptyCards := card.Cards{}
	hearts = emptyCards.Filter(isHeart)

	fmt.Println(hearts, hearts == nil) // empty returns empty
	// Output:
	// [] true
	// [] false
}

func ExampleCards_Filter_nonDestructive() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }

	hand := card.Cards{
		card.HA, // Heart!
		card.CK,
		card.HQ, // Heart!
		card.SJ,
		card.ST,
	}

	hearts := hand.Filter(isHeart)

	fmt.Println(hearts)
	fmt.Println(hand) // not mutated

	// the return value is a newly allocated Cards
	hearts[0] = card.SA

	fmt.Println(hearts) // mutated
	fmt.Println(hand)   // not affected
	// Output:
	// [HA HQ]
	// [HA CK HQ SJ ST]
	// [SA HQ]
	// [HA CK HQ SJ ST]
}

func ExampleCards_Include() {
	cs := card.Cards{card.SA, card.S2}

	fmt.Println(cs.Include())
	fmt.Println(cs.Include(card.SA))
	fmt.Println(cs.Include(card.S2))
	fmt.Println(cs.Include(card.S3))          // not included
	fmt.Println(cs.Include(card.SA, card.SA)) // considered as one SA
	fmt.Println(cs.Include(card.SA, card.S2))
	fmt.Println(cs.Include(card.SA, card.S3)) // includes SA, not S3
	// Output:
	// true
	// true
	// true
	// false
	// true
	// true
	// false
}

func ExampleCards_Include_nilEmpty() {
	nilCards := card.Cards(nil)
	fmt.Println(nilCards.Include())
	fmt.Println(nilCards.Include(card.SA))

	emptyCards := card.Cards{}
	fmt.Println(emptyCards.Include())
	fmt.Println(emptyCards.Include(card.SA))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleCards_Move() {
	deck := card.NewPiquetPack()
	hand, deck := deck.Take(3)

	fmt.Println(deck.Move(0, hand))
	fmt.Println(deck.Move(1, hand))
	fmt.Println(deck.Move(2, hand))
	// Output:
	// [HA H7 H8] [H9 HT HJ HQ HK CA C7 C8 C9 CT CJ CQ CK DK DQ DJ DT D9 D8 D7 DA SK SQ SJ ST S9 S8 S7 SA]
	// [HA H7 H8 H9] [HT HJ HQ HK CA C7 C8 C9 CT CJ CQ CK DK DQ DJ DT D9 D8 D7 DA SK SQ SJ ST S9 S8 S7 SA]
	// [HA H7 H8 H9 HT] [HJ HQ HK CA C7 C8 C9 CT CJ CQ CK DK DQ DJ DT D9 D8 D7 DA SK SQ SJ ST S9 S8 S7 SA]
}

func ExampleCards_Move_nilEmpty() {
	p := func(dst, src card.Cards) {
		fmt.Println(dst, src, dst == nil, src == nil)
	}

	nilCards := card.Cards(nil)
	emptyCards := card.Cards{}

	p(nilCards.Move(0, nilCards))
	p(nilCards.Move(1, nilCards))
	p(nilCards.Move(1, emptyCards))

	p(emptyCards.Move(0, nilCards))
	p(emptyCards.Move(0, emptyCards))
	p(emptyCards.Move(1, emptyCards))
	// Output:
	// [] [] true true
	// [] [] true true
	// [] [] false true
	// [] [] true false
	// [] [] false false
	// [] [] false false
}

func ExampleCards_Move_nonDestructive() {
	var (
		from = card.Cards{card.HA, card.H2, card.H3}
		to   = card.Cards{card.SA, card.S2, card.S3}
	)

	dst, src := from.Move(2, to)

	fmt.Println(dst, src)
	fmt.Println(to, from) // not mutated

	// the return values are newly allocated Cards
	dst[0] = card.SK
	src[0] = card.HK

	fmt.Println(dst, src) // mutated
	fmt.Println(to, from) // not affected

	// both the return values are independent of each other
	cd, cs := cap(dst), cap(src)
	fmt.Println(&dst[:cd][cd-1] != &src[:cs][cs-1]) // not share the underlying array
	// Output:
	// [SA S2 S3 HA H2] [H3]
	// [SA S2 S3] [HA H2 H3]
	// [SK S2 S3 HA H2] [HK]
	// [SA S2 S3] [HA H2 H3]
	// true
}

func ExampleCards_Partition() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }

	hand := card.Cards{
		card.HA, // Heart!
		card.CK,
		card.HQ, // Heart!
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Partition(isHeart))
	// Output:
	// [HA HQ] [CK SJ ST]
}

func ExampleCards_Partition_nilEmpty() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }

	nilCards := card.Cards(nil)
	hearts, others := nilCards.Partition(isHeart)

	fmt.Println(hearts, others, hearts == nil, others == nil) // nil returns nil

	emptyCards := card.Cards{}
	hearts, others = emptyCards.Partition(isHeart)

	fmt.Println(hearts, others, hearts == nil, others == nil) // empty returns empty
	// Output:
	// [] [] true true
	// [] [] false false
}

func ExampleCards_Partition_nonDestructive() {
	isHeart := func(c card.Card) bool { return c.Suit() == card.Heart }

	hand := card.Cards{
		card.HA, // Heart!
		card.CK,
		card.HQ, // Heart!
		card.SJ,
		card.ST,
	}

	hearts, others := hand.Partition(isHeart)

	fmt.Println(hearts, others)
	fmt.Println(hand) // not mutated

	// the return values are newly allocated Cards
	hearts[0] = card.SA
	others[0] = card.HK

	fmt.Println(hearts, others) // mutated
	fmt.Println(hand)           // not affected

	// both the return values are independent of each other
	ch, co := cap(hearts), cap(others)
	fmt.Println(&hearts[:ch][ch-1] != &others[:co][co-1]) // not share the underlying array
	// Output:
	// [HA HQ] [CK SJ ST]
	// [HA CK HQ SJ ST]
	// [SA HQ] [HK SJ ST]
	// [HA CK HQ SJ ST]
	// true
}

func ExampleCards_Remove() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Remove())
	fmt.Println(hand.Remove(card.SA))
	fmt.Println(hand.Remove(card.SA, card.ST))
	fmt.Println(hand.Remove(card.SA, card.S2)) // includes SA, not S2
	// Output:
	// [SA SK SQ SJ ST]
	// [SK SQ SJ ST]
	// [SK SQ SJ]
	// [SK SQ SJ ST]
}

func ExampleCards_Remove_duplicates() {
	hand := card.Cards{
		card.SK, // duplicate
		card.SK, // duplicate
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Remove(card.SK))          // removes both of them
	fmt.Println(hand.Remove(card.SQ, card.SQ)) // considered as one SQ
	// Output:
	// [SQ SJ ST]
	// [SK SK SJ ST]
}

func ExampleCards_Remove_nilEmpty() {
	nilCards := card.Cards(nil)
	removed := nilCards.Remove()

	fmt.Println(removed, removed == nil) // nil returns nil

	emptyCards := card.Cards{}
	removed = emptyCards.Remove()

	fmt.Println(removed, removed == nil) // empty returns empty
	// Output:
	// [] true
	// [] false
}

func ExampleCards_Remove_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	removed := hand.Remove(card.SJ, card.ST)

	fmt.Println(removed)
	fmt.Println(hand) // not mutated

	// the return value is a newly allocated Cards
	removed[0] = card.S2

	fmt.Println(removed) // mutated
	fmt.Println(hand)    // not affected
	// Output:
	// [SA SK SQ]
	// [SA SK SQ SJ ST]
	// [S2 SK SQ]
	// [SA SK SQ SJ ST]
}

func ExampleCards_Reverse() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Reverse())
	// Output:
	// [ST SJ SQ SK SA]
}

func ExampleCards_Reverse_nilEmpty() {
	nilCards := card.Cards(nil)
	reversed := nilCards.Reverse()

	fmt.Println(reversed, reversed == nil) // nil returns nil

	emptyCards := card.Cards{}
	reversed = emptyCards.Reverse()

	fmt.Println(reversed, reversed == nil) // empty returns empty
	// Output:
	// [] true
	// [] false
}

func ExampleCards_Reverse_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	reversed := hand.Reverse()

	fmt.Println(reversed)
	fmt.Println(hand) // not mutated

	// the return value is a newly allocated Cards
	reversed[0] = card.S2

	fmt.Println(reversed) // mutated
	fmt.Println(hand)     // not affected
	// Output:
	// [ST SJ SQ SK SA]
	// [SA SK SQ SJ ST]
	// [S2 SJ SQ SK SA]
	// [SA SK SQ SJ ST]
}

func ExampleCards_Shuffle() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	r := rand.New(rand.NewSource(0))

	for i := 0; i < 10; i++ {
		fmt.Println(hand.Shuffle(r))
	}
	// Output:
	// [SQ SJ SK SA ST]
	// [SQ SJ SA ST SK]
	// [SQ SK SJ SA ST]
	// [SQ SA SJ SK ST]
	// [ST SK SQ SA SJ]
	// [SQ SA SJ ST SK]
	// [SJ SA ST SQ SK]
	// [ST SQ SA SK SJ]
	// [SK SJ ST SQ SA]
	// [SQ SA ST SK SJ]
}

func ExampleCards_Shuffle_defaultSource() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	// uses the default source of random numbers when passed the nil *rand.Rand.
	rand.Seed(0)
	fmt.Println(hand.Shuffle(nil))
	// Output:
	// [SQ SJ SK SA ST]
}

func ExampleCards_Shuffle_nilEmpty() {
	r := rand.New(rand.NewSource(0))

	nilCards := card.Cards(nil)
	shuffled := nilCards.Shuffle(r)

	fmt.Println(shuffled, shuffled == nil) // nil returns nil

	emptyCards := card.Cards{}
	shuffled = emptyCards.Shuffle(r)

	fmt.Println(shuffled, shuffled == nil) // empty returns empty
	// Output:
	// [] true
	// [] false
}

func ExampleCards_Shuffle_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	r := rand.New(rand.NewSource(0))
	shuffled := hand.Shuffle(r)

	fmt.Println(shuffled)
	fmt.Println(hand) // not mutated

	// the return value is a newly allocated Cards
	shuffled[0] = card.S2

	fmt.Println(shuffled) // mutated
	fmt.Println(hand)     // not affected
	// Output:
	// [SQ SJ SK SA ST]
	// [SA SK SQ SJ ST]
	// [S2 SJ SK SA ST]
	// [SA SK SQ SJ ST]
}

func ExampleCards_Size() {
	p := func(cs card.Cards) {
		fmt.Println(cs.Size())
	}

	p(nil)
	p(card.Cards{})
	p(card.Cards{card.SA})
	p(card.Cards{card.SA, card.SK, card.SQ, card.SJ, card.ST})
	// Output:
	// 0
	// 0
	// 1
	// 5
}

func ExampleCards_Sort() {
	// Suit: C < D < H < S
	// Rank: A < 2 < ... < K
	suitAscRankAsc := func(c1, c2 card.Card) int {
		return int(c1) - int(c2)
	}

	hand := card.Cards{
		card.HA,
		card.CK,
		card.HQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Sort(suitAscRankAsc))
	// Output:
	// [CK HA HQ ST SJ]
}

func ExampleCards_Sort_nilEmpty() {
	suitAscRankAsc := func(c1, c2 card.Card) int {
		return int(c1) - int(c2)
	}

	nilCards := card.Cards(nil)
	sorted := nilCards.Sort(suitAscRankAsc)

	fmt.Println(sorted, sorted == nil) // nil returns nil

	emptyCards := card.Cards{}
	sorted = emptyCards.Sort(suitAscRankAsc)

	fmt.Println(sorted, sorted == nil) // empty returns empty
	// Output:
	// [] true
	// [] false
}

func ExampleCards_Sort_nonDestructive() {
	// Suit: C < D < H < S
	// Rank: A < 2 < ... < K
	suitAscRankAsc := func(c1, c2 card.Card) int { return int(c1) - int(c2) }

	hand := card.Cards{
		card.HA,
		card.CK,
		card.HQ,
		card.SJ,
		card.ST,
	}

	sorted := hand.Sort(suitAscRankAsc)

	fmt.Println(sorted)
	fmt.Println(hand) // not mutated

	// the return value is a newly allocated Cards
	sorted[0] = card.SK

	fmt.Println(sorted) // mutated
	fmt.Println(hand)   // not affected
	// Output:
	// [CK HA HQ ST SJ]
	// [HA CK HQ SJ ST]
	// [SK HA HQ ST SJ]
	// [HA CK HQ SJ ST]
}

func ExampleCards_Take() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Take(0)) // |o o o o o
	fmt.Println(hand.Take(1)) //  o|o o o o
	fmt.Println(hand.Take(4)) //  o o o o|o
	fmt.Println(hand.Take(5)) //  o o o o o|
	fmt.Println(hand.Take(6)) //  o o o o o .|
	// Output:
	// [] [SA SK SQ SJ ST]
	// [SA] [SK SQ SJ ST]
	// [SA SK SQ SJ] [ST]
	// [SA SK SQ SJ ST] []
	// [SA SK SQ SJ ST] []
}

func ExampleCards_Take_nilEmpty() {
	nilCards := card.Cards(nil)
	taken, remaining := nilCards.Take(0)

	fmt.Println(taken, remaining, taken == nil, remaining == nil) // nil returns nil

	emptyCards := card.Cards{}
	taken, remaining = emptyCards.Take(0)

	fmt.Println(taken, remaining, taken == nil, remaining == nil) // empty returns empty
	// Output:
	// [] [] true true
	// [] [] false false
}

func ExampleCards_Take_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	taken, remaining := hand.Take(2)

	fmt.Println(taken, remaining)
	fmt.Println(hand) // not mutated

	// the return values are newly allocated Cards
	taken[0] = card.S2
	remaining[0] = card.S3

	fmt.Println(taken, remaining) // mutated
	fmt.Println(hand)             // not affected

	// both the return values are independent of each other
	ct, cr := cap(taken), cap(remaining)
	fmt.Println(&taken[:ct][ct-1] != &remaining[:cr][cr-1]) // not share the underlying array
	// Output:
	// [SA SK] [SQ SJ ST]
	// [SA SK SQ SJ ST]
	// [S2 SK] [S3 SJ ST]
	// [SA SK SQ SJ ST]
	// true
}

func ExampleCards_Top() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	fmt.Println(hand.Top())
	// Output:
	// SA [SK SQ SJ ST]
}

func ExampleCards_Top_nilEmpty() {
	nilCards := card.Cards(nil)
	top, others := nilCards.Top()

	fmt.Println(top, others, others == nil) // nil returns (nil, nil)

	emptyCards := card.Cards{}
	top, others = emptyCards.Top()

	fmt.Println(top, others, others == nil) // empty returns (nil, empty)
	// Output:
	// <nil> [] true
	// <nil> [] false
}

func ExampleCards_Top_nonDestructive() {
	hand := card.Cards{
		card.SA,
		card.SK,
		card.SQ,
		card.SJ,
		card.ST,
	}

	top, remaining := hand.Top()

	fmt.Println(top, remaining)
	fmt.Println(hand) // not mutated

	// the return values share no addresses with the Cards
	*top = card.S2
	remaining[0] = card.S3

	fmt.Println(top, remaining) // mutated
	fmt.Println(hand)           // not affected
	// Output:
	// SA [SK SQ SJ ST]
	// [SA SK SQ SJ ST]
	// S2 [S3 SQ SJ ST]
	// [SA SK SQ SJ ST]
}

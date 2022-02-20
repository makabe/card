# card

[![Go Reference](https://pkg.go.dev/badge/github.com/makabe/card.svg)](https://pkg.go.dev/github.com/makabe/card)
[![CI](https://github.com/makabe/card/actions/workflows/ci.yml/badge.svg)](https://github.com/makabe/card/actions/workflows/ci.yml)
[![Codecov](https://codecov.io/gh/makabe/card/branch/main/graph/badge.svg?token=JWWE7F83DN)](https://codecov.io/gh/makabe/card)

A [French-suited playing cards](https://en.wikipedia.org/wiki/French-suited_playing_cards) library for Go.

## Install

```sh
go get github.com/makabe/card
```

## Example

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/makabe/card"
)

func suitAscRankAsc(c1, c2 card.Card) int {
	return int(c1) - int(c2)
}

func main() {
	// create new deck
	deck := card.NewStandardDeck()

	// shuffle deck
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	deck = deck.Shuffle(r)

	// deal hands
	pHand, deck := deck.Take(5)
	oHand, _ := deck.Take(5)

	// sort hands
	pHand = pHand.Sort(suitAscRankAsc)
	oHand = oHand.Sort(suitAscRankAsc)

	// print
	fmt.Println("You:", pHand)
	fmt.Println("Opponent:", oHand)
}
```

## License

[MIT](https://github.com/makabe/card/blob/main/LICENSE)

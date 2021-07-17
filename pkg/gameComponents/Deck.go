
package gameComponents

import (
	// "fmt"
	"math/rand"
	"time"
)

type Deck struct {
    numberOfDecksInUse int
    cards []Card
}



func NewDeck(numberOfDecks int) Deck {
    rand.Seed(time.Now().UnixNano())

    if (numberOfDecks < 1) {
        numberOfDecks = 1
    } else if (numberOfDecks > 8) {
        numberOfDecks = 8
    }

    d := Deck {numberOfDecks, getShuffledNewDeckOfCards(numberOfDecks)}

    return d
}

func getShuffledNewDeckOfCards(numberOfDecks int) []Card {

    cards:= make([]Card, 52*numberOfDecks+1)
    for d := 0; d < numberOfDecks; d++ {
        for j := 0; j < 4; j++ {
            for i := 0; i < 13; i++ {
                cards[52*d + i + 13*j] = NewCard(i + 13*j)
                // fmt.Printf("Card: %d  %d  %d\n", e.cards[i + 13*j].Val, e.cards[i + 13*j].Ord, e.cards[i + 13*j].Suit)
            }
        }
    }
    Shuffle(&cards)

    cardCut := GetCardCut()
    positionForCardCut := int(float64(len(cards)) * 0.8)
    cards = append(cards, cardCut)
    copy(cards[positionForCardCut+1:], cards[positionForCardCut:])
    cards[positionForCardCut] = cardCut

    return cards
}


func (d *Deck) Draw() Card {
    // fmt.Printf("Drawing from the %d cards.\n", len(e.cards))
	// fmt.Printf("Drawing, using %d decks.\n", d.numberOfDecksInUse)

    // c := e.cards[len(e.cards) - 1]
    // e.cards = e.cards[:len(e.cards)-1]
    //
    // copy(e.cards[len(e.cards) - 1:], e.cards[len(e.cards):])
    // e.cards[len(e.cards)-1] = Card{} // or the zero value of T
    // e.cards = e.cards[:len(e.cards)-1]
    card := d.cards[len(d.cards)-1]
    if IsCardCut(card) {
        d.cards = getShuffledNewDeckOfCards(d.numberOfDecksInUse)
        card = d.cards[len(d.cards)-1]
    }

    d.cards = d.cards[:len(d.cards)-1]

    return card
}


func Shuffle(cards *[]Card) {
    rand.Shuffle(len(*cards), func(i, j int) {
        (*cards)[i], (*cards)[j] = (*cards)[j], (*cards)[i]
    })
}

/*
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}
*/

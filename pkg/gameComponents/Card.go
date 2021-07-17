
package gameComponents

import (
    "strconv"
)

type Card struct {
    Suit int
    Ord int
    FaceUp bool
}
func NewCard(num int) Card {
    card := Card {}
    card.Suit = num / 13 + 1
    card.Ord = num % 13 + 1
    card.FaceUp = false
    return card
}
func GetCardCut() Card {
    return Card{0, 0, false}
}

func IsCardCut(card Card) bool {
    return card.Suit == 0
}

func (c Card) String() string {
    if IsCardCut(c) {
        return "c-c"
    }
    if !c.FaceUp {
        return "_-_"
    }

    suits := []string{"S", "H", "C", "D"}
    ord := strconv.Itoa(c.Ord)
    switch(ord) {
    case "1":
        ord = "A"
    case "11":
        ord = "J"
    case "12":
        ord = "Q"
    case "13":
        ord = "K"
    }

    return ord + "-" + suits[c.Suit-1]
}


package players
import "blackjack_test/pkg/gameComponents"

import "fmt"

type Player struct {
	name string
	hand []gameComponents.Card
}
func NewPlayer(name string) *Player {
    return &Player{name, make([]gameComponents.Card, 0)}
}

type IPlayer interface {
	CardsInHand() int
	AddToHand(c gameComponents.Card)
	GetHand(onlyFaceUp bool) []gameComponents.Card
	PrintHand()
	RevealHand()
	GetScore(aceRaisedIfPossible bool, includeFaceDownCard bool) int
	GetName() string
	Reset()

	WantsToHit(opponentsHand []gameComponents.Card) bool
}



func (p *Player) CardsInHand() int {
    // fmt.Printf("%d cards in hand\n", len(e.hand))
	return len(p.hand)
}

func (p *Player) AddToHand(card gameComponents.Card) {
	p.hand = append(p.hand, card)
	// fmt.Printf("AddToHand, hand now: ", e.hand, "\n")
}

func (p *Player) GetHand(includeFaceDown bool) ([]gameComponents.Card) {
	if !includeFaceDown && (len(p.hand) >= 2) {
		faceDownIdx := -1
		for index, card := range p.hand {
			if card.FaceUp == false {
				faceDownIdx = index
				break
			}
		}
		fmt.Printf("Found index of face down: %d\n", faceDownIdx)
		if faceDownIdx >= 0 {
			fmt.Printf("1: %v\n", p.hand)
			faceUpCards := make([]gameComponents.Card, 0)
			faceUpCards = append(faceUpCards, p.hand[:faceDownIdx]...)
			faceUpCards = append(faceUpCards, p.hand[faceDownIdx+1:]...)
			fmt.Printf("2: %v\n", faceUpCards)
			fmt.Printf("3: %v\n", len(faceUpCards))
			return faceUpCards
		} else {
			return p.hand
		}
	}
	return p.hand
}

func (p *Player) PrintHand() {
	fmt.Printf("%s hand: %v\n", p.name, p.GetHand(true))
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) Reset() {
	p.hand = make([]gameComponents.Card, 0)
}


func (p *Player) GetScore(aceRaisedIfPossible bool, includeFaceDownCard bool) int {
    // fmt.Printf("GetScore \n")
	return ScoreCards(p.GetHand(true), aceRaisedIfPossible, includeFaceDownCard)
}

func ScoreCards(cards []gameComponents.Card, aceRaisedIfPossible bool, includeFaceDownCard bool) int {
    // fmt.Printf("ScoreCards \n")
	score := 0
    hasAce := false
	// fmt.Printf("cards to score: %v\n", cards)
	for _, card := range cards {
		// fmt.Printf("card: %v (%d)\n", card, card.Ord)
		if (card.FaceUp || includeFaceDownCard) {
			if card.Ord == 1 {
	            score += 1
	            hasAce = true
	        } else if card.Ord >= 10 {
	            score += 10
	        } else {
				score += card.Ord
			}
		}
	}

    // remember that only one ace's value can change, if the player has more than one. The second must always be 1.
    if hasAce && aceRaisedIfPossible {
        if score + 10 <= 21 {
            score += 10
        }
    }

    return score
}


func (p *Player) RevealHand() {
	for index, _ := range p.hand {
		p.hand[index].FaceUp = true
	}
}

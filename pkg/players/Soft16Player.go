
package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type Soft16Player struct {
  *Player
}

func (p *Soft16Player) WantsToHit(opponentHand []gameComponents.Card) bool {
    // if ScoreCards(opponentHand) > 8 {
    //     return true
    // }
    if p.GetScore(false, true) < 16 {
        return true
    }
    return false
}

func NewSoft16Player() *Soft16Player {
    return &Soft16Player {NewPlayer("Soft16Player")}
}


package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type Soft17Player struct {
  *Player
}

func (p *Soft17Player) WantsToHit(opponentHand []gameComponents.Card) bool {
    // if ScoreCards(opponentHand) > 8 {
    //     return true
    // }
    if p.GetScore(true, true) < 17 {
        return true
    }
    return false
}

func NewSoft17Player() *Soft17Player {
    return &Soft17Player {NewPlayer("Soft17Player")}
}

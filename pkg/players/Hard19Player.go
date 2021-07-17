
package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type Hard19Player struct {
  *Player
}

func (p *Hard19Player) WantsToHit(opponentHand []gameComponents.Card) bool {
    // if ScoreCards(opponentHand) > 8 {
    //     return true
    // }
    if p.GetScore(true, true) < 19 {
        return true
    }
    return false
}

func NewHard19Player() *Hard19Player {
    return &Hard19Player {NewPlayer("Hard19Player")}
}

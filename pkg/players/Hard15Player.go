
package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type Hard15Player struct {
  *Player
}

func (p *Hard15Player) WantsToHit(opponentHand []gameComponents.Card) bool {
    // if ScoreCards(opponentHand) > 8 {
    //     return true
    // }
    if p.GetScore(true, true) < 15 {
        return true
    }
    return false
}

func NewHard15Player() *Hard15Player {
    return &Hard15Player {NewPlayer("Hard15Player")}
}


package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type Hard18Player struct {
  *Player
}

func (p *Hard18Player) WantsToHit(opponentHand []gameComponents.Card) bool {
    // if ScoreCards(opponentHand) > 8 {
    //     return true
    // }
    if p.GetScore(true, true) < 18 {
        return true
    }
    return false
}

func NewHard18Player() *Hard18Player {
    return &Hard18Player {NewPlayer("Hard18Player")}
}

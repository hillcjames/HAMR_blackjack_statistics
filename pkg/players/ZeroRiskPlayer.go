
package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type ZeroRiskPlayer struct {
  *Player
}

func (p *ZeroRiskPlayer) WantsToHit(opponentHand []gameComponents.Card) bool {
    // if ScoreCards(opponentHand) > 8 {
    //     return true
    // }
    if p.GetScore(true, true) < 12 {
        return true
    }
    return false
}

func NewZeroRiskPlayer() *ZeroRiskPlayer {
    return &ZeroRiskPlayer {NewPlayer("ZeroRiskPlayer")}
}


package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type Hard17Player struct {
  *Player
}

func (d *Hard17Player) WantsToHit(opponentHand []gameComponents.Card) bool {
    // fmt.Printf("Starting Hard17Player \n")
    if d.GetScore(true, true) < 17 {
        return true
    }
    return false
}

func NewHard17Player() *Hard17Player {
    return &Hard17Player {NewPlayer("Hard17Player")}
}

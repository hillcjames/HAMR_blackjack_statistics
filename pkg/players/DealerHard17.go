
package players

import "blackjack_test/pkg/gameComponents"

// import "fmt"


type DealerHard17 struct {
  *Player
}

func (d *DealerHard17) WantsToHit(opponentHand []gameComponents.Card) bool {
    // fmt.Printf("Starting DealerHard17 \n")
    if d.GetScore(true, true) < 17 {
        return true
    }
    return false
}

func NewDealerHard17() *DealerHard17 {
    return &DealerHard17 {NewPlayer("DealerHard17")}
}

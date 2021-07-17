
package players

import "blackjack_test/pkg/gameComponents"

// import "math/rand"
// import "fmt"
// Float64 - like, if it's a 9 showing, then only hit whatever percentage of time matches the probability of my current number being lower than that are. I think.


type BasicStratPlayer struct {
  *Player
}

func (p *BasicStratPlayer) WantsToHit(opponentHand []gameComponents.Card) bool {
    // fmt.Printf("Starting BasicStratPlayer \n")
    myScore := p.GetScore(true, true)
    myDroppedAcesScore := p.GetScore(false, true)
    oppScore := ScoreCards(opponentHand, true, false)
    if myScore < 12 {
       // fmt.Printf("\tstrat 1 (%d, %d)\n", myScore, oppScore)
        return true
    } else if myScore < 15 && myDroppedAcesScore < 5 {
        return true
    } else if oppScore + 10 <= 16 && oppScore + 10 > 11 {
       // fmt.Printf("\tstrat 3 (%d, %d)\n", myScore, oppScore)
        return false
    } else if myScore < oppScore + 10 && myScore < 18 {
       // fmt.Printf("\tstrat 2 (%d, %d)\n", myScore, oppScore)
        return true
    }
   // fmt.Printf("\tstrat 5 (%d, %d)\n", myScore, oppScore)
    return false
}

func NewBasicStratPlayer() *BasicStratPlayer {
    return &BasicStratPlayer {NewPlayer("BasicStratPlayer")}
}

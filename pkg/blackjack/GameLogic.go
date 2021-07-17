
package blackjack
import "blackjack_test/pkg/players"
import "blackjack_test/pkg/gameComponents"

import "fmt"

const (  // iota is reset to 0
        playerWin = iota  // == 0
        dealerWin = iota  // == 1
        tie = iota  // == 2
        error = iota  // == 3
)

type GameState struct {
    playerHand []gameComponents.Card
    dealerHand []gameComponents.Card
}

type Game struct {
    deck gameComponents.Deck
    player players.IPlayer
    dealer players.IPlayer
    quiet bool
}

func NewGame(p1 players.IPlayer, p2 players.IPlayer, quiet bool) Game {
    return Game {gameComponents.NewDeck(6), p1, p2, quiet}
}

func (g *Game) StartGame() int {

    g.player.Reset()
    g.dealer.Reset()
    g.Deal()
    //
	// fmt.Printf("Winner is: %v\n", g.dealer.GetHand(false))

    //
    // g.dealer.RevealHand()
    // g.PrintGameState()

    // initialGameState := GameState{g.player.GetHand(true), g.dealer.GetHand(true)}
    results := g.coreLoop()
    g.player.RevealHand()
    g.dealer.RevealHand()
    // endGameState := GameState{g.player.GetHand(true), g.dealer.GetHand(true)}

    if !g.quiet {
        g.PrintGameState()
        fmt.Printf("Winner is: %d\n", results)
    }
    // fmt.Printf("setup is: %v\n", initialGameState)
	// fmt.Printf("end is: %v\n", endGameState)
    return results
}

func (g *Game) coreLoop() int {
    if playerHasNatural21(g.player) {
        return playerWin
    }

    temp := g.dealer.GetHand(true)

    playerScore := g.takeTurn(g.player, temp)

    if playerScore > 21 {
        return dealerWin
    }

    dealerScore := g.takeTurn(g.dealer, g.player.GetHand(true))

    if dealerScore > 21 || playerScore > dealerScore {
        return playerWin
    } else if dealerScore > playerScore {
        return dealerWin
    } else if dealerScore == playerScore {
        return tie
    }

    g.PrintGameState()
    return error
}

func playerHasNatural21(player players.IPlayer) bool {
    return len(player.GetHand(true)) == 2 && player.GetScore(true, true) == 21
    // return false
}

func (g *Game) takeTurn(player players.IPlayer, opponentHand []gameComponents.Card) int {

    if !g.quiet {
        fmt.Printf("%s turn.\n", player.GetName())
        fmt.Printf("Hand: %v\n", player.GetHand(true))
    	// fmt.Printf("\t%s wants to hit: %s \n", player.GetName(), player.WantsToHit(opponentHand))
    }

    for (player.WantsToHit(opponentHand) && player.GetScore(false, true) <= 21) {
        g.hit(player, true)

        if !g.quiet {
            fmt.Printf("Hand: %v\n", player.GetHand(true))
        }
    	// fmt.Printf("\t%s wants to hit: %s \n", player.GetName(), player.WantsToHit(opponentHand))
    }
    score := player.GetScore(true, true)
    if !g.quiet {
        fmt.Printf("%s ended with: %d\n\n", player.GetName(), score)
    }
    return score
}


func (g *Game) Deal() {
    g.hit(g.player, true)
    g.hit(g.player, true)

    g.hit(g.dealer, false)
    g.hit(g.dealer, true)
}
func (g *Game) hit(player players.IPlayer, faceUp bool) {
    card := g.deck.Draw()
    card.FaceUp = faceUp
    player.AddToHand(card)
}

func (g *Game) PrintGameState() {
    fmt.Printf("PrintGameState\n")
    g.player.PrintHand()
    fmt.Printf("Score: %v\n", g.player.GetScore(true, true))

    g.dealer.PrintHand()
    fmt.Printf("Score: %v\n\n", g.dealer.GetScore(true, true))
}

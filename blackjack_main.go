/**
Simple system of pitting two systems of play against each other - simplest difference to start out with, would be just whether to hold on 16 or 17 or whatever.

After that, I want to see the stats of how the loses in each pairing break down, in the face up cards (especially the initial card)

So like, if the dealer holds on 16, and they have a six showing - what percentage of their losses, is having a six showing?


*/

package main

// import "fmt"


import "blackjack_test/pkg/metrics"


func main() {
    metrics.StartSimulation()
}

/**
class game {
  deck
  dealer, player

  getShowingCards(playerPerspective): visibleGameState

  playMatch

  playRound {
    if player.wantsToHit
      player.hit
    if busted(player)

    if dealer.wantsToHit
      dealer.hit
    if busted(dealer)

  }

  endGame(winner, previousGameState)
    add to tally

  addToTally()
    what cards were showing for both players right before someone made the winning/losing move.
    I think to start off, just make a histogram of each card value saying ho often that person lost against a given opponent when that card showed up first


}

class individualEvaluation(player)

    Or... Just for each person in isolation, what's the likelihood of them getting to 21 with any particular initial card showing.
      How many turns does it take them to bust? I have 1-that.

  while true
    if player.wantsToHit
      player.hit(deck)
    if busted(player)

class visibleGameState:
  list p1 cards
  list p2 cards


class deck:
  4 of each 1-9
  16 10s
  shuffle
  reset

card:
  value
  suit
  rank

player
string key or name or something
  abstract wantsToHit(visibleGameState)
  getHandSum(visibleGameState) // automatically shift ace - start high.
*/

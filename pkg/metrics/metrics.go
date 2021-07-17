/*
Results, with 100000 trials
For Soft16Player vs DealerHard17:
        0:  0.404930
        1:  0.498280
        2:  0.096790
        3:  0.000000
        Expected win/loss, for 100 games @ $25 min:      -225
For ZeroRiskPlayer vs DealerHard17:
        0:  0.425480
        1:  0.519440
        2:  0.055080
        3:  0.000000
        Expected win/loss, for 100 games @ $25 min:      -225
For Hard15Player vs DealerHard17:
        0:  0.420890
        1:  0.495280
        2:  0.083830
        3:  0.000000
        Expected win/loss, for 100 games @ $25 min:      -175
For BasicStratPlayer vs DealerHard17:
        0:  0.425270
        1:  0.485060
        2:  0.089670
        3:  0.000000
        Expected win/loss, for 100 games @ $25 min:      -150
For BasicStratPlayer vs Soft16Player:
        0:  0.453210
        1:  0.470510
        2:  0.076280
        3:  0.000000
        Expected win/loss, for 100 games @ $25 min:      -50
For Soft16Player vs BasicStratPlayer:
        0:  0.433290
        1:  0.489900
        2:  0.076810
        3:  0.000000
        Expected win/loss, for 100 games @ $25 min:      -125
Conclusions:

    To start off with the probably-obvious - the house always wins.
    More specifically though, the two biggest factors seem to be: the fact that the player has to go first (not suprising), and the rule of always treating aces as high unless
    forced to drop them.
*/
package metrics

import (
    "fmt"
)

import "blackjack_test/pkg/blackjack"
// import "blackjack_test/pkg/gameComponents"
import "blackjack_test/pkg/players"

type Experiment struct {
    player players.IPlayer
    dealer players.IPlayer
    trials int
    rawResults []int
}

func newExperiment(player players.IPlayer, dealer players.IPlayer, trials int) Experiment {
    return Experiment{player, dealer, trials, make([]int, 4)}
}

func StartSimulation() {

    trials := 100000

    experiments := []Experiment {
        newExperiment(players.NewSoft16Player(), players.NewDealerHard17(), trials),
        newExperiment(players.NewZeroRiskPlayer(), players.NewDealerHard17(), trials),
        newExperiment(players.NewHard15Player(), players.NewDealerHard17(), trials),
        newExperiment(players.NewBasicStratPlayer(), players.NewDealerHard17(), trials),
        newExperiment(players.NewBasicStratPlayer(), players.NewSoft16Player(), trials),
        newExperiment(players.NewSoft16Player(), players.NewBasicStratPlayer(), trials),
    }

    fmt.Printf("Starting experiments\n")

    progressBarWidth := 80
    progressSinceBarUpdate := 0
    for i := 0; i < progressBarWidth-1; i++ {
        fmt.Printf("_")
    }
    fmt.Printf("|\r")
    for expI, _ := range experiments {
        for i := 0; i < trials; i++ {
            // fmt.Printf("Experiment %d, trial %d\n", expI+1, i + 1)
            game := blackjack.NewGame(experiments[expI].player, experiments[expI].dealer, true)
            result := game.StartGame()
            experiments[expI].rawResults[result] += 1

            progressSinceBarUpdate += 1
            if progressSinceBarUpdate > (len(experiments) * trials) / progressBarWidth {
                fmt.Printf("█")
                progressSinceBarUpdate = 0
            }
        }
    }
    fmt.Printf("\n")
    fmt.Printf("Done with experiments\n")

    fmt.Printf("Results, with %d trials\n", trials)
    for _, experiment := range experiments {
        fmt.Printf("For %s vs %s:\n", experiment.player.GetName(), experiment.dealer.GetName())
        for i := 0; i < len(experiment.rawResults); i++ {
            winRate := float64(experiment.rawResults[i]) / float64(experiment.trials)
            fmt.Printf("\t%d:  %f\n", i, winRate)
        }
        fmt.Printf("\tExpected win/loss, for 100 games @ $25 min: \t %d \n", 25 * (int(100*float64(experiment.rawResults[0]) / float64(experiment.trials)) - int(100*float64(experiment.rawResults[1]) / float64(experiment.trials)))  )
    }
}


package metrics

import (
    "fmt"
    "math"
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
    return Experiment{player, dealer, trials, make([]int, 3)}
}

func StartSimulation() {
    resultStrings := []string{
        "Plr",
        "Dlr",
        "Tie",
    }

    trials := 1000000

    // This could be improved with a player-generator, assuming Go has generators
    experiments := []Experiment {
        // newExperiment(players.NewZeroRiskPlayer(), players.NewHard17Player(), trials),
        // newExperiment(players.NewHard15Player(), players.NewHard17Player(), trials),
        // newExperiment(players.NewSoft16Player(), players.NewHard17Player(), trials),
        // newExperiment(players.NewSoft17Player(), players.NewHard17Player(), trials),
        // newExperiment(players.NewHard17Player(), players.NewHard17Player(), trials),
        // newExperiment(players.NewHard18Player(), players.NewHard17Player(), trials),
        // newExperiment(players.NewHard19Player(), players.NewHard17Player(), trials),
        newExperiment(players.NewBasicStratPlayer(), players.NewHard17Player(), trials),
        newExperiment(players.NewBasicStratPlayer(), players.NewSoft17Player(), trials),
        newExperiment(players.NewBasicStratPlayer(), players.NewSoft16Player(), trials),
        newExperiment(players.NewBasicStratPlayer(), players.NewZeroRiskPlayer(), trials),
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
            fmt.Printf("\t%s:  %f\n", resultStrings[i], winRate)
        }
        winRate := float64(experiment.rawResults[0])
        lossRate := float64(experiment.rawResults[1])
        fmt.Printf("\tNaive Expected win/loss, for 100 games @ $25 min: \t %d \n", 25 * (int(100*winRate / float64(experiment.trials)) - int(100*lossRate / float64(experiment.trials)))  )
        fmt.Printf("\tExpected longest winning streak, during 100 games: \t %d \n", expectedLongestStreak(100, winRate/(winRate+lossRate)) )
        fmt.Printf("\tExpected longest losing streak, during 100 games: \t %d \n", expectedLongestStreak(100, lossRate/(winRate+lossRate)) )
        fmt.Printf("\n")
    }
}

func expectedLongestStreak(trials int, rate float64) int {
    /*
    This attempts to invert+modify the formula `e = 2^(n+1) -2` for expected trials to get N consecutive heads, using blackjack's uneven odds.
    https://math.stackexchange.com/questions/364038/expected-number-of-coin-tosses-to-get-five-consecutive-heads
    */
    inverseRate := 1/rate

    return int(math.Log(float64(trials) + inverseRate) / math.Log(inverseRate) - 1)
}

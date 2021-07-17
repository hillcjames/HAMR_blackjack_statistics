# HAMR_blackjack_statistics
Calculate expected loss for different blackjack strategies.

Requires Go to be installed.

To run:
 `go run blackjack_main.go`


 New player strategies can be easily added by extending Player - just copy any of the existing player classes, change the name both in the filename and throughout the file, and edit WantsToHit() as you see fit.
 Add a new experiment using that player, in `pkg/metrics.go`.

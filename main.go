package main

import (
	"github.com/ZAF07/go-design-pattern/strategy"
)


func main() {
	// One way to execute the logic
	algo := strategy.NewComputerAlgo("Human")
	logic := strategy.NewMainStruct(2,3,algo)
	logic.AddToStartValue(3)
	logic.CalculateMain()

	// Can execute the logic this way also
	// strategy.RunStrategyClient("COMPUTER", "AlphaZero")
}
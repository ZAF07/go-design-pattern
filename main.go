package main

import (
	"github.com/ZAF07/go-design-pattern/strategy"
)


func main() {
	// ONE WAY TO EXECUTE THE STRATEGY
	
	// Init the initial algo strategy  
	algo := strategy.NewComputerAlgo("Computer")
	logic := strategy.NewMainStruct(2,3,algo)
	logic.AddToStartValue(3)
	logic.CalculateMain()

	// Switching the algo strategy
	algo2 := strategy.NewHumanAlgo("Human")
	logic.SwitchAlgo(algo2)
	logic.AddToEndValue(10)
	logic.CalculateMain()

	// Can execute the logic this way also
	// strategy.RunStrategyClient("COMPUTER", "AlphaZero")
	// strategy.RunStrategyClient("HUMAN", "Zaffere")
}
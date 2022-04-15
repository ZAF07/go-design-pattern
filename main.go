package main

import (
	"github.com/ZAF07/go-design-pattern/adapter"
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

	// ADAPTER PATTERN

	macbook := adapter.NewMacBook()
	
	// Init a new appple charger
	appleCharger := adapter.NewCharger()
	appleCharger.StartCharging(macbook)
	
	// Windows products would'nt work with the appleCharger
	acerLaptop := adapter.NewWindowsLaptop()
	/* 
		THIS WOULD CAUSE AN ERROR, AS APPLE CHARGERS CAN ONLY WORK WITH APPLE PRODUCTS
		WE NEED AN ADAPTER!!
	*/ 
	// appleCharger.StartCharging(acerLaptop)

	// Solution here, pass the windows product via an adapter
	windowsMacAdapter := adapter.NewWindowsChargerAdapter(*acerLaptop)
	appleCharger.StartCharging(windowsMacAdapter)
}
package strategy

import (
	"fmt"
)

// MAIN STRUCT TO BE USED BY CLIENTS, HAS PROPERTIES AND METHODS EXPOSED
type MainStruct struct {
	CalculateAlgo Maths // Anything that implements the Maths interface IS the Maths interface
	Start         int
	End           int
}

// Initiates a new MainStruct
func NewMainStruct(startNum, endNum int, algo Maths) *MainStruct {
	return &MainStruct{
		CalculateAlgo: algo,
		Start:         startNum,
		End:           endNum,
	}
}

func (s *MainStruct) CalculateMain() {
	s.CalculateAlgo.calc(s)
}

func (s *MainStruct) AddToStartValue(i int) {
	s.Start += i
}

func (s *MainStruct) AddToEndValue(i int) {
	s.End += i
}

func (s *MainStruct) SwitchAlgo(algo Maths) {
	s.CalculateAlgo = algo
}

func RunStrategyClient(strategy, name string) {
	switch strategy {
	case "HUMAN":
		fmt.Printf("STRATEGY TYPE : %s", strategy)
		fmt.Println()
		algo := NewHumanAlgo(name)
		exposedStruct := NewMainStruct(1, 2, algo)
		// Calling exposed struct's method, which calls the method of the given struct
		exposedStruct.CalculateMain()
	case "COMPUTER":
		fmt.Printf("STRATEGY TYPE : %s", strategy)
		fmt.Println()
		algo := NewComputerAlgo(name)
		exposedStruct := NewMainStruct(1, 2, algo)
		// Calling exposed struct's method, which calls the method of the given struct
		exposedStruct.CalculateMain()
	}
}

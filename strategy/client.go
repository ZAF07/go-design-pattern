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

func (s *MainStruct) CalculateMain() {
	s.CalculateAlgo.calc(s)
}

func (s *MainStruct) Add(i int) {
	s.Start += 1
}

// Initiates a new MainStruct
func newMainStruct(startNum, endNum int, algo Maths) *MainStruct {
	return &MainStruct{
		CalculateAlgo: algo,
		Start:         startNum,
		End:           endNum,
	}
}

func RunStrategyClient(strategy, name string) {
	switch strategy {
	case "HUMAN":
		fmt.Printf("STRATEGY TYPE : %s", strategy)
		fmt.Println()
		algo := NewHuman(name)
		exposedStruct := newMainStruct(1, 2, algo)
		// Calling exposed struct's method, which calls the method of the given struct
		exposedStruct.CalculateMain()
	case "COMPUTER":
		fmt.Printf("STRATEGY TYPE : %s", strategy)
		fmt.Println()
		algo := NewComputer(name)
		exposedStruct := newMainStruct(1, 2, algo)
		// Calling exposed struct's method, which calls the method of the given struct
		exposedStruct.CalculateMain()
	}
}

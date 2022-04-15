package strategy

import (
	"fmt"

	"github.com/ZAF07/go-design-pattern/strategy/strategies"
)

// MAIN STRUCT TO BE USED BY CLIENTS, HAS PROPERTIES AND METHODS EXPOSED
type MainStruct struct {
	CalculateAlgo maths // Anything that implements the maths interface IS the maths interface
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
func newMainStruct(startNum, endNum int, algo maths) *MainStruct {
	return &MainStruct{
		CalculateAlgo: algo,
		Start:         startNum,
		End:           endNum,
	}
}

func Run(strategy, name string) {
	switch strategy {
	case "HUMAN":
		fmt.Printf("STRATEGY TYPE : %s", strategy)
		fmt.Println()
		algo := strategies.NewHuman(name)
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
	// Initiates a new struct
	// zaf := NewHuman("ZAFFERE")
	// com := NewComputer("AlphaZero")

	// fmt.Printf("STRATEGIES AVAILABLE : [%s, %s]", zaf.Name, com.Name)
	// fmt.Println()

}

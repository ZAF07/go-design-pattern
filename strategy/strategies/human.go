package strategies

import (
	"fmt"

	"github.com/ZAF07/go-design-pattern/strategy"
)

//  STRUCT AND IT'S METHODS
type Human struct {
	Name string
}

// This method takes in as an argument, (the main struct that was exposed to our clients (MainStruct, and all its properties and arguments))
func (s *Human) calc(i *strategy.MainStruct) {
	fmt.Printf("NAME OF STRATEGY STRUCT USED : %s", s.Name)
	fmt.Println()
	fmt.Println("Human calculated : ", i.Start+i.End)
}

//  Initiates a new Human struct
func NewHuman(name string) *Human {
	return &Human{
		Name: name,
	}
}
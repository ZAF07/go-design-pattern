package strategy

import (
	"fmt"
)

//  STRUCT AND IT'S METHODS
type Human struct {
	Name string
}

//  Initiates a new Human struct
func NewHumanAlgo(name string) *Human {
	return &Human{
		Name: name,
	}
}

// This method takes in as an argument, (the main struct that was exposed to our clients (MainStruct aka Client, and all its properties and arguments))
func (s *Human) calc(i *MainStruct) {
	fmt.Printf("NAME OF STRATEGY STRUCT USED : %s", s.Name)
	fmt.Println()
	fmt.Println("Human calculated : ", i.Start+i.End)
}
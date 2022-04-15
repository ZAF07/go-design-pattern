package strategy

import (
	"fmt"
)

// //  STRUCT AND THIER METHODS
// type Human struct {
// 	Name string
// }

type Computer struct {
	Name string
}

// This method takes in as an argument, (the main struct that was exposed to our clients (MainStruct, and all its properties and methods))
func (s *Computer) calc(i *MainStruct) {
	fmt.Printf("NAME OF STRATEGY STRUCT USED : %s", s.Name)
	fmt.Println()
	fmt.Println("Computer calculated : ", i.Start+i.End)
}

// This method takes in as an argument, (the main struct that was exposed to our clients (MainStruct, and all its properties and arguments))
// func (s *Human) calc(i *MainStruct) {
// 	fmt.Printf("NAME OF STRATEGY STRUCT USED : %s", s.Name)
// 	fmt.Println()
// 	fmt.Println("Human calculated : ", i.Start+i.End)
// }

// //  Initiates a new Human struct
// func NewHuman(name string) *Human {
// 	return &Human{
// 		Name: name,
// 	}
// }

// Initiates a new Computer struct
func NewComputer(name string) *Computer {
	return &Computer{
		Name: name,
	}
}
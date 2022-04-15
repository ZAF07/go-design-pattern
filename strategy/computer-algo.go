package strategy

import "fmt"

type Computer struct {
	Name string
}

// Initiates a new Computer struct
func NewComputerAlgo(name string) *Computer {
	return &Computer{
		Name: name,
	}
}

// This method takes in as an argument, the main struct that is exposed to our clients (MainStruct aka Client, and all its properties and methods)
func (s *Computer) calc(i *MainStruct) {
	fmt.Printf("NAME OF STRATEGY STRUCT USED : %s", s.Name)
	fmt.Println()
	fmt.Println("Computer calculated : ", i.Start+i.End)
}
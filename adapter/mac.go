package adapter

import "fmt"

type Mac struct {}

func NewMacBook() *Mac {
	return &Mac{}
}

func (m *Mac) InsertMacChargerHead() {
	fmt.Println("Mac charger head inserted")
}
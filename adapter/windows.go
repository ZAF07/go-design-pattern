package adapter

import "fmt"

type Windows struct {

}

func NewWindowsLaptop() *Windows {
	return &Windows{}
}

func (w *Windows) InsertWindowsCharger() {
	fmt.Println("Windows charger head inserted")
}
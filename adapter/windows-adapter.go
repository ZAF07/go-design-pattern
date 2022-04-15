package adapter

import "fmt"

type WindowsChargerAdapter struct {
	NonMacMachine Windows
}

func NewWindowsChargerAdapter(i Windows) *WindowsChargerAdapter {
	return &WindowsChargerAdapter{
		NonMacMachine: i,
	}
}

// This method is where we adapt top the interface
// We make the struct implement the AppleCharge interface and in the method, we call the windows charge method
func (cd *WindowsChargerAdapter) InsertMacChargerHead() {
	fmt.Println("Inserted Mac charger head via adapter")
	cd.NonMacMachine.InsertWindowsCharger()
}
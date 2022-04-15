package adapter

import "fmt"

type Charger struct {}

func NewCharger() *Charger {
	return&Charger{}
}

func (c Charger) StartCharging(i AppleCharge) {
	fmt.Println("Start Charging")
	i.InsertMacChargerHead()
}
package main

import (
	"fmt"
)

type car struct {
	gasPedal uint16
	breakePedal uint16
	steeringPedal int16
	topSpeedKmh float64
}

func main() {
	instCar := car{gasPedal: 200, breakePedal: 0, steeringPedal: 150, topSpeedKmh: 200}
	fmt.Println(instCar.gasPedal)
}
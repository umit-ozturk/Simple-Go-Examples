package main

import (
	"fmt"
)

const uSixTeenBitMax float64  = 65535
const kmhMultiple float64  = 1.60943

type car struct {
	gasPedal uint16
	breakePedal uint16
	steeringPedal int16
	topSpeedKmh float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/uSixTeenBitMax)
	
}

func main() {
	instCar := car{gasPedal: 20000, breakePedal: 0, steeringPedal: 150, topSpeedKmh: 200}
	fmt.Println(instCar.gasPedal)
	fmt.Println(instCar.kmh())
}
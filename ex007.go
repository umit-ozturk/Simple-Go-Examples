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

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/uSixTeenBitMax/kmhMultiple)

}

func (c *car) newTopSpeed(newSpeed float64) float64 {
	c.topSpeedKmh = newSpeed
	return float64(c.gasPedal) * (c.topSpeedKmh/uSixTeenBitMax/kmhMultiple)

}

func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	instCar := car{gasPedal: 20000, breakePedal: 0, steeringPedal: 150, topSpeedKmh: 200}
	fmt.Println(instCar.gasPedal)
	fmt.Println(instCar.kmh())
	fmt.Println(instCar.mph())
	instCar.newTopSpeed(500)
	instCar = newerTopSpeed(instCar, 100)
	fmt.Println(instCar.kmh())
	fmt.Println(instCar.mph())
}
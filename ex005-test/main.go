package main

import (
	"fmt"
)

func Calculate(x int) (result int){
	result = x + 2
	return result
}

func main()  {
	fmt.Println("Home Endpoint Hit.")
	result := Calculate(5)
	fmt.Println(result)
}
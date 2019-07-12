package main

import "fmt"


func main(){
	grades := make(map[string]float32)
	grades["Timmy"] = 50
	grades["Jenny"] = 50
	grades["Steve"] = 50

	TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	for k, v := range grades{
		fmt.Println(k, ":", v)
	}
}

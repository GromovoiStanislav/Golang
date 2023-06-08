package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	
	fmt.Println(rand.Intn(10) + 1)

	var pi64 = math.Pi
	fmt.Println(pi64) // Выводит: 3.141592653589793
	
	var pi32 float32 = math.Pi
	fmt.Println(pi32) // Выводит: 3.1415927

	fmt.Println(math.Abs(1 - 10)) // 9

	fmt.Println(math.Sqrt(9)) // 3 
	
	fmt.Println(math.Pow(2, 2)) // 4
	fmt.Println(math.Pow(2, 3)) // 8

	fmt.Println(math.Round(2.35555)) // 2
	fmt.Println(math.Round(2.4999)) // 2
	fmt.Println(math.Round(2.5)) //3

	fmt.Println(math.Floor(2.999)) // 2
	fmt.Println(math.Ceil(2.0001)) // 3
}

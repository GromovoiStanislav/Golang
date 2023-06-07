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

}

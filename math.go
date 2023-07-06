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

	fmt.Println(math.Hypot(6, 8)) // 10 гипотенуза
	fmt.Println(math.Sqrt(math.Pow(6, 2) + math.Pow(8, 2)))


	fmt.Println("MaxInt8: ", math.MaxInt8) // MaxInt8:  127
	fmt.Println("MaxInt8: ", math.MaxInt8) // MinInt8:  -128
	fmt.Println("MinInt8: ", math.MinInt8) // MaxInt16:  32767
	fmt.Println("MaxInt16: ", math.MaxInt16) // MinInt16:  -32768
	fmt.Println("MinInt16: ", math.MinInt16) // MaxInt32:  2147483647
	fmt.Println("MaxInt32: ", math.MaxInt32) // MinInt32:  -2147483648
	fmt.Println("MinInt32: ", math.MinInt32) // MaxInt64:  9223372036854775807
	fmt.Println("MaxInt64: ", math.MaxInt64) // MaxInt64:  9223372036854775807
	fmt.Println("MinInt64: ", math.MinInt64) // MinInt64:  -9223372036854775808
	fmt.Println("MaxUint8: ", math.MaxUint8) // MaxUint8:  255
	fmt.Println("MaxUint16: ", math.MaxUint16) // MaxUint16:  65535
	fmt.Println("MaxUint32: ", math.MaxUint32) // MaxUint32:  4294967295
	fmt.Println("MaxUint64: ", uint64(math.MaxUint64)) //MaxUint64:  18446744073709551615
	
}

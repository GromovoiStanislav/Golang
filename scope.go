package main

import "fmt"

func ExampleScope1() {
	var v int = 1

	{ //внутренний лексический блок
		var v string = "2"
		fmt.Println(v) // 2 
	}

	fmt.Println(v) // 1
}

var i string = "Строка"


func main() {
	ExampleScope1()
	// Output:
	// 2
	// 1

	i := "Другая строка"
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		i := true
		fmt.Println(i)
	}
	// 0
	// true
	// 1
	// true
	// 2
	// true
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		i := 20 // Объявляем НОВУЮ переменую
		fmt.Println(i)
	}
	// 0
	// 20
	// 1
	// 20
	// 2
	// 20
	fmt.Println(i)
	// Другая строка
}
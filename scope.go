package main

import "fmt"



var i string = "Строка"


func main() {
	
	v := 1
	{ //внутренний лексический блок
		var v string = "str"
		fmt.Println(v) // str 
	}
	fmt.Println(v) // 1




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
package main

import "fmt"

/*
Также есть функция Sprintf() которая работает как и Printf(), 
за исключением того что она ничего не печатает, 
а возвращает результат форматирования
*/

func main() {
	var a float64 = 100.123456789
	result := fmt.Sprintf("%.2f", a)
	fmt.Printf("%q", result) // вывод: "100.12"
	// result будет типа string
}
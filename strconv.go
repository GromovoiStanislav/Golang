package main

import (
	"fmt"
	"strconv"
)

func main() {
	

	s := strconv.Itoa(-42) // "-42"
	fmt.Println(s + "**") // -42**

	

	i, _ := strconv.Atoi("-42")  // -42
	fmt.Println(i+1) // -41

	n, _ := strconv.Atoi("-42**")  // -42
	fmt.Println(n) // 0


	num, err := strconv.ParseFloat("4.1", 32)
	if err != nil {
		fmt.Println("Вы ввели не число! Ошибка")
	}
	fmt.Println(num) //4.099999904632568

	num, err := strconv.ParseFloat("4.1", 64)
	if err != nil {
		fmt.Println("Вы ввели не число! Ошибка")
	}
	fmt.Println(num) // 4.1
}
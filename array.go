package main

import "fmt"

func main() {
	var a [3]int
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3} // длина массива будет определена Go в зависимости от количества указанных при инициализации значений
	d := [3]int{1: 12} //явно указать значение, которое должно быть присвоено элементу массива с определенным индексом
	e := [6]int{1: 12, 2: 11, 5: 3}

	fmt.Println(a) // [0 0 0]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]
	fmt.Println(e) // [0 12 11 0 0 3]

	numbers := [5]int{1,2,3,4,5}
	fmt.Println(numbers)    // [1 2 0 0 0]

	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue


	fmt.Println(b == c) // true
	fmt.Println(a == b) // false
	fmt.Println(d == e) // invalid operation: d == e (mismatched types [3]int and [6]int)
}
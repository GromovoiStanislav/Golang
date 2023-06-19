package main

import "fmt"

func main() {
	var a [3]int
	a2 := [3]int{}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3} // длина массива будет определена Go в зависимости от количества указанных при инициализации значений
	d := [3]int{1: 12} //явно указать значение, которое должно быть присвоено элементу массива с определенным индексом
	e := [6]int{1: 12, 2: 11, 5: 3}

	fmt.Println(a) // [0 0 0]
	fmt.Println(a2) // [0 0 0]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]
	fmt.Println(e) // [0 12 11 0 0 3]

	
	
	fmt.Println(b == c) // true
	fmt.Println(a == b) // false
	// fmt.Println(d == e) // invalid operation: d == e (mismatched types [3]int and [6]int)


	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors) // [red green blue]
	colors[0], colors[2] = colors[2], colors[0]
	fmt.Println(colors) // [blue green red]


	numbers := [5]int{1,2}
	fmt.Println(numbers)    // [1 2 0 0 0]
	fmt.Println(numbers[0]) // 1  
    fmt.Println(numbers[4]) // 0   
    numbers[4] = 5   
    fmt.Println(numbers)    // [1 2 0 0 5]
    fmt.Println(numbers[len(numbers) - 1]) // 5


	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		// 1
		// 2
		// 0
		// 0
		// 5
	}

	for idx := range numbers {
		fmt.Println(numbers[idx])
		// 1
		// 2
		// 0
		// 0
		// 5
	}

	for _, elem := range numbers {
		fmt.Println(elem)
		// 1
		// 2
		// 0
		// 0
		// 5
	}

	for idx, elem := range numbers {
		fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
		// Элемент с индексом 0: 1
		// Элемент с индексом 1: 2
		// Элемент с индексом 2: 0
		// Элемент с индексом 3: 0
		// Элемент с индексом 4: 5

		elem = 100 // Не изменит массив !!!
		numbers[idx] = 100 // Изменит массив !!!
	}


	for range numbers {
		fmt.Println("hello")
		// hello
		// hello
		// hello
		// hello
		// hello
	}
}
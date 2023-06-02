package main

import "fmt"


func star() {
    fmt.Println("**********")
}

func mult(a int, b int) {
    fmt.Println(a * b)
} 


func plus(a, b, c int){
    fmt.Println(a + b + c)
}


func square(num int) int {
    return num * num
}


func isEven(num int) bool {
    return num % 2 == 0
}


func concat(a, b string) string {
    return a + b
}

func plusMinus(x, y int) (int, int) {
    return x + y, x - y
}


func swap(x, y int) (int, int) {
    return y, x
} 


func sum(nums ...int) (res int) {
    for _, n := range nums {
        res += n
    }
    return
}


func welcome() {
    fmt.Println("Добро пожаловать!")
}


func main() {
    star() // выведет: **********
	
	mult(2, 3) // выведет: 6
	plus(2, 3, 4) // выведет: 9

	res := square(4)
    fmt.Println(res) // выведет: 16


	fmt.Println(plusMinus(15, 5)) // выведет: 20 10

	a, b := swap(42, 8)
    fmt.Println(a, b) // выведет 8 42


	fmt.Println(sum())        // 0
    fmt.Println(sum(1))       // 1
    fmt.Println(sum(1, 2))    // 3
    fmt.Println(sum(1, 2, 3)) // 6

	defer welcome()
    fmt.Println("Привет!") // сначала выведет "Привет! " и только после этого выведет "Добро пожаловать!"

	defer fmt.Println("Раз")
	defer fmt.Println("Два")
	fmt.Println("Три") 
	// Три
	// Два
	// Раз

	fmt.Println("начало")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("конец")
	// начало
	// конец
	// 4
	// 3
	// 2
	// 1
	// 0
} 
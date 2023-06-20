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

func add(x, y int, a, b, c float32){
    var z = x + y
    var d = a + b + c
    fmt.Println("x + y = ", z)
    fmt.Println("a + b + c = ", d)
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


func sum(nums  ...int) (res int) {
	// nums  - это slice !!!
    for _, v := range nums  {
        res += v 
    }
    return
}

func foo(name string, numbers ...int) (string, []int) {
	return name, numbers
}

func foo1(name string, args ...interface{})(string, []interface{}) {
	return name, args
}




func welcome() {
    fmt.Println("Добро пожаловать!")
}


func fn() (int, error) {
	// Какая-то полезная работа
	return 0, nil
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

	fmt.Println(sum([]int{1, 2}...))    // 3
    fmt.Println(sum([]int{1, 2, 3}...)) // 6
	
	slice := []int{42, 33, 22, 11}
	fmt.Println(sum(slice...)) // 108
	arr := [4]int{42, 33, 22, 11}
	fmt.Println(sum(arr[:]...)) // 108

	fmt.Println(foo("Это срез:", 1, 2, 3, 4, 5)) // Это срез: [1 2 3 4 5]

	name, args := foo1("Это срез:", "fff", 2, "ddd", 4, 5)
	fmt.Println(name, args)
	for _, arg := range args {
		fmt.Printf("%T %v\n", arg, arg)
	}
	// Это срез: [fff 2 ddd 4 5]
	// string fff
	// int 2
	// string ddd
	// int 4
	// int 5




	// Игнорирование возвращаемых значений
	fn()

	i, _ := fn()
	fmt.Println(i) // 0

	_, err := fn()
	if err == nil {
		fmt.Println("Ошибок нет") // Ошибок нет
	}





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
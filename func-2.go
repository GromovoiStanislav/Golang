package main

import "fmt"

// Тип функции:



// тип: func(int, int) int
func add(x int, y int) int {
	return x + y
}

// тип: func(int, int) int
func subtract(x int, y int) int {
	 return x - y
}

// тип: func(int, int) int
func multiply(x int, y int) int {
    return x * y
}

// тип: func(string)
func display(message string){
    fmt.Println(message)
}


// Функция как результат другой функции
func selectFn(n int) (func(int, int) int){
    if n==1 {
        return add
    }else if n==2{
        return subtract
    }else{
        return multiply
    }
}
 

// Функция как параметры другой функции
func action(n1 int, n2 int, operation func(int, int) int){
    result := operation(n1, n2)
    fmt.Println(result)
}


func main() {

	// var f func(int, int) int = add
	//или так:
	f := add        //теперь переменная f указывает на функцию add с типом func(int, int) int
	fmt.Println(f(3, 4)) // 7
	fmt.Println(f(4, 5)) //9

	f = multiply    // теперь переменная f указывает на функцию multiply с ТЕМ ЖЕ типом func(int, int) int
    fmt.Println(f(3, 4)) // 12

	// f = display      // ошибка, так как функция display имеет тип func(string)
	//var f1 func(string) = display   // норм 
	//или так:
	f1 := display   // норм 
	f1("hello")


	action(10, 25, add)     // 35
    action(5, 6, multiply)  // 30


	f2 := selectFn(1)
	fmt.Println(f2(3, 4))           // 7
	fmt.Println(selectFn(1)(3, 4)) // 7

	f3 := selectFn(3)
	fmt.Println(f3(3, 4))           // 12
	fmt.Println(selectFn(3)(3, 4)) // 12


	slice := []int{-2, 4, 3, -1, 7, -4, 23}
     
    sumOfEvens := sum(slice, isEven)    // сумма четных чисел
    fmt.Println(sumOfEvens)             // -2
     
    sumOfPositives := sum(slice, isPositive)    // сумма положительных чисел
    fmt.Println(sumOfPositives)                 // 37

}


func isEven(n int) bool{
    return n % 2 == 0
}


func isPositive(n int) bool{
    return n > 0
}
 
// Функция как параметры другой функции
func sum(numbers []int, criteria func(int) bool) int{
    result := 0
    for _, val := range numbers{
        if(criteria(val)){
            result += val
        }
    }
    return result
}



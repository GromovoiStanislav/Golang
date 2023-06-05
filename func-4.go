package main

import "fmt"

// Рекурсивные функции:
 

func factorial(n uint) uint{
 
    if n == 0{
        return 1
    }
    return n * factorial(n - 1)
}


func fibonacci (n uint) uint{
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return fibonacci (n - 1) + fibonacci (n - 2)
}


func main() {
     
    fmt.Println(factorial(4))   // 24
    fmt.Println(factorial(5))   // 120
    fmt.Println(factorial(6))   // 720

	fmt.Println(fibonacci(4))  // 3
    fmt.Println(fibonacci(5))  // 5
    fmt.Println(fibonacci(6))  // 8

}



// без рекурсии
func getFibonacci(n int) int {
	var fib []int = []int{0, 1}
	for i := 1; i < n; i++ {
		fib = append(fib, fib[i]+fib[i-1])
	}
	return fib[n]
}
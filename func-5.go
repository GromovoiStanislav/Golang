package main

import "fmt"

// Передача указателей в функции:

func change(x int) {
    x = 22
}

func changePointer(pointer *int) {
    *pointer = 22
}

func doublePointer(pointer *int){
    *pointer *= 2
}


func main() {
    num := 14

    change(num) // не изменила значение
    fmt.Println(num) // выведет 14

    changePointer(&num) // изменила значение
    fmt.Println(num) // выведет 22

	doublePointer(&num) // изменила значение
    fmt.Println(num) // выведет 44
}
package main

import "fmt"

/*
Оператор defer позволяет выполнить определенную операцию после каких-то действий (даже если сработает panic)

Команда defer помещает вызов функции в стек. Поэтому они выполняются в очередности -LIFO (Last-In, First-Out)

defer запоминает значения переменных, переданных в функцию, на момент объявления defer, а не на момент его вызова. То есть условно:
a:=5
defer myFunc(a) // когда вызовется myFunc - будет передано значение 5, а не 7
a = 7
*/

func finish() {
	fmt.Println("Program has been finished")
}

func myFunc(a int) {
	fmt.Println("myFunc", a)
}

func main() {
	defer finish()

	a := 3
	defer myFunc(a)

	a = 2
	defer myFunc(a)

	defer func() {
		for i := 0; i < 3; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println("!")
	}()

	defer fmt.Println("Program has been started")

	fmt.Println("Program is working")

	myFunc(1)

	panic("panic! panic! panic!")
}

// Program is working
// myFunc 1
// Program has been started
// 0 1 2 !
// myFunc 2
// myFunc 3
// Program has been finished
// panic: panic! panic! panic!
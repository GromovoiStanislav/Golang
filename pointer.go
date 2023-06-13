package main

import "fmt"

func main() {
	
	x := 200
	y := &x // var y *int = &x
	*y++
	fmt.Println(x) // 201


	z := new(int)
	fmt.Println(*z, z) // 0 0xc00001c030
	*z = 1
	fmt.Println(*z, z) // 1 0xc00001c030




	a := 23
	p := &a

	fmt.Println(a)  //23
	fmt.Println(p)  // 0xc00001c030
	fmt.Println(*p) //23

	*p = 11

	fmt.Println(*p)  // 11
	fmt.Println(a)   // 11
	fmt.Println(&a) // 0xc00001c030
	fmt.Println(*&a) // 11




	// указатель на указатель
	a := 200
	aPtr := &a
	*aPtr++
	bPtr := &aPtr
	**bPtr++ // указатель на указатель
	*aPtr++
	fmt.Println(a) // 203






	// Передача указателей в функции:
	num := 14

    change(num) // не изменила значение
    fmt.Println(num) // выведет 14

    changePointer(&num) // изменила значение
    fmt.Println(num) // выведет 22

	doublePointer(&num) // изменила значение
    fmt.Println(num) // выведет 44
	
}



func change(x int) {
    x = 22
}

func changePointer(pointer *int) {
    *pointer = 22
}

func doublePointer(pointer *int){
    *pointer *= 2
}
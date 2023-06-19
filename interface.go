package main

import (
	"fmt"
)

// пустой интерфейс interface{} означает тип данных, под который подходит любой другой тип
func Print(arg interface{}) {
	fmt.Println(arg)
}

func main() {
	Print("hello!") // hello!
	Print(123) // 123
	Print([]int{1, 5, 10}) // [1 5 10]
	


}

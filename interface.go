package main

import (
	"fmt"
)

/*
Нулевое значение для интерфейса — это пустое значение nil.
*/


// пустой интерфейс interface{} означает тип данных, под который подходит любой другой тип
func Print(arg interface{}) {
	fmt.Println(arg)
}



// объявление интерфейса
type Printer interface {
    Print()
}

// нигде не указано, что User реализует интерфейс Printer
type User struct {
    email string
}
// структура User имеет метод Print, как в интерфейсе Printer. 
// Следовательно, во время компиляции запишется связь между User и Printer
func (u *User) Print() {
    fmt.Println("My email is", u.email)
}


// функция принимает как аргумент интерфейс Printer !!!
func TestPrint(p Printer) {
    p.Print()
}





func main() {
	Print("hello!") // hello!
	Print(123) // 123
	Print([]int{1, 5, 10}) // [1 5 10]
	

	// в функцию TestPrint передается структура User, 
	// а так как она реализует интерфейс Printer, все работает без ошибок
	TestPrint(&User{email: "test@test.com"})
}

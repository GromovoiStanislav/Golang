package main

import "fmt"
 
type mile int
type kilometer int
 
func distanceToEnemy (distance mile){
    fmt.Println("расстояние для противника:",distance, "миль")

}
 

type library []string
func (l library) print(){
    for _, val := range l{
        fmt.Println(val)
    }
}
 
func printBooks(lib library){
    for _, value := range lib{
        fmt.Println(value)
    }
}



type counter int
// передается указатель, чтобы можно было изменить состояние счетчика "c"
func (c *counter) inc() {
	*c++
}



type errorCode string



func main() {  
    var distance mile = 5
    distanceToEnemy(distance) // расстояние для противника: 5 миль
    
	// var distance2 kilometer = 5
    // distanceToEnemy(distance2)   // ! ошибка

	// var distance3 int = 5
	// distanceToEnemy(distance3) // ! ошибка

	distanceToEnemy(5) // расстояние для противника: 5 миль


	myLibrary := library{"Book1", "Book2", "Book3"}
	printBooks(myLibrary)

	myLibrary2 := []string{"Book1", "Book2", "Book3"}
	printBooks(myLibrary2)


	// myLibrary3 := []int{1, 2, 3}
	// printBooks(myLibrary3) // ! ошибка


	lib := library{"Book1", "Book2", "Book3"}
	lib.print()
	// Book1
	// Book2
	// Book3


	c := counter(0)
	c.inc()
	c.inc()
	// (&c).inc() // передается указатель на счетчик &c
	fmt.Println(c) // 2


	// Алиас можно конвертировать в оригинальный тип и обратно:
	ec := errorCode("internal")
	fmt.Println(ec) // internal
	fmt.Println(string(ec)) // internal

}
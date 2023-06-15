package main

import "fmt"
 
type mile int
type kilometer int
 
func distanceToEnemy (distance mile){
    fmt.Println("расстояние для противника:",distance, "миль")

}
 

type library []string
 
func printBooks(lib library){
    for _, value := range lib{
        fmt.Println(value)
    }
}


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


}
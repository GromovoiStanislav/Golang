package main

import "fmt"


type Test struct {
    a int
    b int
}
func (x Test) do() int {
    return x.a - x.b
}


type Contact struct {
    name string
    age  int
} 
func (x Contact) welcome() {
    fmt.Println("Welcome", x.name)
} 
func (ptr *Contact) increase(val int) {
	ptr.age += val
}


// обычная функция
func welcome2(x Contact) {
	fmt.Println("welcome", x.name)
}


func main() {
    t := Test{11, 2}
    fmt.Println(t.do()) // 9
    
    

    a := Contact{"Андрей", 33}
    // a := Contact{name: "Андрей", age: 33}
	fmt.Println(a) // {Андрей 33}
    a.welcome() // Welcome Андрей
    welcome2(a) // Welcome Андрей
    a.increase(5)
	fmt.Println(a.age) // 38




    
    a.age = 27
	fmt.Println(a.age)  // 27
	fmt.Println(a.name) // Андрей


     // указаьтель на структуру
    aPtr := &a
	fmt.Println(aPtr.age) // 27
	fmt.Println((*aPtr).age) // 27
    fmt.Println(aPtr) // &{Андрей 27}
	fmt.Println(*aPtr) // {Андрей 27}


    // указатель при создании новой структуры
    p := &Contact{"Андрей", 22}
	fmt.Println(p.age)    // 22
	fmt.Println((*p).age) // 22
	fmt.Println(p)        // &{Андрей 22}
	fmt.Println(*p)       //{ Андрей 22}
}



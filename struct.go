package main

import (
	"fmt"
	"math"
)


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
func (cPtr *Contact) welcome() {
    fmt.Println("Welcome", cPtr.name)
} 
func (cPtr *Contact) increase(val int) {
	cPtr.age += val
}


// обычная функция
func welcome2(x Contact) {
	fmt.Println("welcome", x.name)
}



type coord struct {
	x, y float64 
}
type Circle struct {
	center coord
	radius float64 
}
func (c *Circle) area() float64 {
    return math.Pi * c.radius  * c.radius 
}
// обычная функция
func circleArea(c *Circle) float64 {
	return math.Pi * c.radius * c.radius
}



type Person struct {
	Name string
}
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
type Android struct {
	Person // Встраиваемые типы
	Model string
}






func main() {
    t := Test{11, 2}
    fmt.Println(t.do()) // 9
    
    var c Circle 
    // c := Circle{}
    cPtr := new(Circle) 
    fmt.Println(c) // {{0 0} 0}
    fmt.Println(cPtr) // &{{0 0} 0}


    circle1 := Circle{
		center: coord{4, 5},
		radius: 7,
	}
	fmt.Println(circle1) // {{4 5} 7}
    fmt.Println(circle1.center) // {4 5}
	fmt.Println(circle1.center.x, circle1.center.y) // 4 5
    
    fmt.Println(circle1.area()) // 153.93804002589985
    fmt.Println(circleArea(&circle1)) // 153.93804002589985



    var andr = Android{
		Model: "R2-D2",
		Person: Person{
			Name: "name",
		},
	}
	andr.Person.Talk() // Hi, my name is Artoo-Detoo
	andr.Talk() // Hi, my name is Artoo-Detoo
	fmt.Println(andr.Name, andr.Person.Name) // Artoo-Detoo Artoo-Detoo
	fmt.Println(andr) //{{Artoo-Detoo} R2-D2}





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



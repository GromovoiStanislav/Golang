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
	Person // Встраиваемый тип
	Model string
}






func main() {
    t := Test{11, 2}
    fmt.Println(t.do()) // 9
    
    var c Circle 
    // c := Circle{}
	fmt.Println(c)     // {{0 0} 0}

	//при инициализации указателя на структуру можно присвоить адрес безымянного объекта
    cPtr := new(Circle) 
    fmt.Println(cPtr)  // &{{0 0} 0}
	fmt.Println(*cPtr) // {{0 0} 0}


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
	//var aPtr *Contact = &a
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




	//можно определять указатели на отдельные поля структуры:
	sam := Contact{name: "Sam", age: 22}
	var agePointer *int = &sam.age // указатель на поле sam.age
	*agePointer = 35               // изменяем значение поля
	fmt.Println(sam)               // {Sam 35}




	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	printNodeValue(current)
	// 4
	// 5
	// 6

	//или
	printNodes(current)
	// 4
	// 5
	// 6

	//или
    for current != nil{
        fmt.Println(current.value)
        current = current.next
    }
	// 4
	// 5
	// 6
}


// поле ссылается на тот же тип
type node struct{
    value int
    next *node //  ! поле должно представлять указатель на структуру
	//next node // ! неправильно
}

// рекурсивный вывод списка
func printNodeValue(n *node){
    fmt.Println(n.value)
    if n.next != nil{
        printNodeValue(n.next)
    }
}
//или
func printNodes(n *node) {
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}
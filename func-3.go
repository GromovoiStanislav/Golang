package main
 
import "fmt"


 
func main() {
	
	// Анонимные функции:
    f1 := func(x, y int) int { return x + y} // тип: func(int, int) int 
	//var f func(int, int) int = func(x, y int) int { return x + y}
    fmt.Println(f1(3, 4))        // 7
    fmt.Println(f1(6, 7))        // 13

    var add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2)) // 3




	// Анонимная функция как аргумент функции
	action(10, 25, func (x int, y int) int { return x + y })    // 35
    action(5, 6, func (x int, y int) int { return x * y })      // 30


	f2 := selectFn(1)
    fmt.Println(f2(2, 3))        	// 5
    fmt.Println(selectFn(1)(4, 5))  // 9
    fmt.Println(selectFn(2)(7, 6))  // 1


	// Замыкание
	f := square()
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25

	fmt.Println(square()())// 9
	fmt.Println(square()())// 9
}



func action(n1 int, n2 int, operation func(int, int) int){
    result := operation(n1, n2)
    fmt.Println(result)
}

// Анонимная функция как результат функции
func selectFn(n int) (func(int, int) int){
    if n==1 {
        return func(x int, y int) int{ return x + y}
    }else if n==2{
        return func(x int, y int) int{ return x - y}
    }else{
        return func(x int, y int) int{ return x * y}
    }
}


// Замыкание
func square() func() int { 
    var x int = 2
    return func() int { 
        x++
        return x * x
    }
}
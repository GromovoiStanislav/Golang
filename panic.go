package main
import "fmt"
 

func divide(x, y float64) float64{
    if y == 0 { 
        panic("Division by zero!")
    }
    return x / y
}


func main() {
    fmt.Println(divide(15, 5)) // 3
    fmt.Println(divide(4, 0)) // panic: Division by zero!
    fmt.Println("Program has been finished")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i) // Передаем копию i в горутину
	}

	time.Sleep(100 * time.Millisecond)
}

// go run .\goroutine-4.go

/*
4
0
3
1
2
*/


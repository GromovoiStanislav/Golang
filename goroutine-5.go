package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("отправка значения 10 в канал")
		ch <- 10
		fmt.Println("отправка значения 20 в канал")
		ch <- 20
		fmt.Println("отправка значения 30 в канал")
		ch <- 30
		fmt.Println("отправка значения 40 в канал")
		ch <- 40
		fmt.Println("отправка значения 50 в канал")
		ch <- 50
		fmt.Println("отправка значения 60 в канал")
		ch <- 60
		fmt.Println("отправка значения 70 в канал")
		ch <- 70

		close(ch)
	}()

	for x := range ch {
		fmt.Println(x) // Чтение из канала
	}
}

// go run .\goroutine-5.go

/*
отправка значения 10 в канал
отправка значения 20 в канал
10
20
отправка значения 30 в канал
отправка значения 40 в канал
30
40
отправка значения 50 в канал
отправка значения 60 в канал
50
60
отправка значения 70 в канал
70
*/


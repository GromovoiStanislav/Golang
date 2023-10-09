package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(id int, phrase string, done chan <- bool) {
	defer func() { done <- true }() // Отправляем сигнал о завершении горутины в канал done

	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)

		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	done := make(chan bool)

	go say(1, "go is awesome", done)
	go say(2, "cats are cute", done)

	// Ожидаем завершения обоих горутин
	<-done
	<-done

	close(done) // Закрываем канал done
}

// go run .\goroutine-3.go

/*
Worker #2 says: cats...
Worker #1 says: go...
Worker #1 says: is...
Worker #2 says: are...
Worker #1 says: awesome...
Worker #2 says: cute...
*/


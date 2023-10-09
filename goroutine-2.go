package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func say(id int, phrase string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины

	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)

		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Увеличиваем счетчик WaitGroup на количество горутин

	go say(1, "go is awesome", &wg)
	go say(2, "cats are cute", &wg)

	wg.Wait() // Ожидаем завершения обоих горутин
}

// go run .\goroutine-2.go

/*
Worker #2 says: cats...
Worker #1 says: go...
Worker #1 says: is...
Worker #2 says: are...
Worker #1 says: awesome...
Worker #2 says: cute...
*/


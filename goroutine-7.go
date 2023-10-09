package main

/*
Задача: Напишите программу на Go, которая параллельно вычисляет факториал двух числа и выводит общую сумму
*/

import (
    "fmt"
    "sync"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
      var num1, num2 int
    fmt.Scanln(&num1, &num2)

    // Создаем каналы для передачи результатов из горутин
    resultChan := make(chan int)


    var wg sync.WaitGroup

    // Горутина для вычисления факториала первого числа
    wg.Add(1)
    go func() {
        defer wg.Done()
        result := factorial(num1)
        resultChan <- result
    }()

    // Горутина для вычисления факториала второго числа
    wg.Add(1)
    go func() {
        defer wg.Done()
        result := factorial(num2)
        resultChan <- result
    }()

    // Чтение результатов из каналов
    result1 := <-resultChan
    result2 := <-resultChan
    
     // Закрываем канал после отправки всех результатов
    close(resultChan)

    // Ожидаем завершения обоих горутин
    wg.Wait()

    // Суммируем результаты
    total := result1 + result2

    // Выводим общую сумму
    fmt.Println(total)
}

// go run .\goroutine-7.go

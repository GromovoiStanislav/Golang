package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	var input int
	_, err := fmt.Scan(&input) // функция Scan возвращает два параметра, но нам сейчас важно проверить только ошибку
	
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
	} else {
		d, err := divide(2, 0)
		if err != nil {
			fmt.Println("error:", err) // Функция вернула непустую ошибку
		} else {
			fmt.Println(d) // Деление прошло без ошибок
		}
	}



	myError := errors.New("my error")
	fmt.Println("", myError) // my error

	err2 := fmt.Errorf("my error: %s", "<error>")
	fmt.Println(err2) // my error: <error>
}
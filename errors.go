package main

import (
	"errors"
	"fmt"
)


func validateName(name string) error {
    if name == "" {
        
		// return errors.New("empty name")
		
		// или оборачиваем в виде "[название метода]: %w". %w — это плейсхолдер для ошибки
		return fmt.Errorf("validateName: %w", errors.New("empty name"))
    }

    if len([]rune(name)) > 50 {
        return errors.New("a name cannot be more than 50 characters")
    }

    return nil
}

// errors.New создает новый объект ошибки
var errDivisionByZero = errors.New("division by zero")

func divide(a int, b int) (int, error) {
	if b == 0 {
		// return -1, errDivisionByZero

		// или оборачиваем в виде "[название метода]: %w". %w — это плейсхолдер для ошибки
		return -1, fmt.Errorf("divide: %w", errDivisionByZero)
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
		if err != nil {// Функция вернула непустую ошибку
			
			if errors.Is(err, errDivisionByZero) {//проверки типов конкретных ошибок
				fmt.Println("error:", err) 
			}

		} else {
			fmt.Println(d) // Деление прошло без ошибок
		}
	}


	// errors.New создает новый объект ошибки
	myError := errors.New("my error")
	fmt.Println("", myError) // my error

	err2 := fmt.Errorf("my error: %s", "<error>")
	fmt.Println(err2) // my error: <error>

	
}
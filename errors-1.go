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



type DivisionByZeroErr struct{}
func (e DivisionByZeroErr) Error() string {
	return "division by zero"
}


func divide(a int, b int) (int, error) {
	if b == 0 {
		// return 0, errDivisionByZero    // for errors.Is
		// return 0, DivisionByZeroErr{}  // for errors.As

		// или оборачиваем в виде "[название метода]: %w". %w — это плейсхолдер для ошибки
		return 0, fmt.Errorf("divide: %w", errDivisionByZero) // for errors.Is
		return 0, fmt.Errorf("divide: %w", DivisionByZeroErr{}) // for errors.As
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
			} else {
				fmt.Println("unknown error:", err)
			}


			if errors.As(err, &DivisionByZeroErr{}) {//Иногда нужно проверить не конкретную ошибку, а целый тип
				fmt.Println("error:", err)
			} else {
				fmt.Println("unknown error:", err)
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
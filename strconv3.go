package main

import (
	"fmt"
	"log"
	"strconv"
)


// Функция intToString преобразует целочисленное значение в строку 
func intToString() {
	old := 50
	oldStr := strconv.Itoa(old) // Преобразование целого числа в строку
	fmt.Println("Hello I'm " + oldStr + " years old")
}

// Функция floatToString преобразует значение с плавающей запятой в строку и выводит результат
func floatToString() {
	points := 133.132
	fmt.Println("Sammy has " + fmt.Sprint(points) + " points")// Преобразование значения с плавающей запятой в строку
}

// Функция stringToInt преобразует строку в целочисленное значение 
func stringToInt() {
	str := "23345"
	num, err := strconv.Atoi(str) // Преобразование строки в целое число

	if err != nil {
		log.Fatal("не удалось конвертировать тип!") // В случае ошибки выводим сообщение и завершаем программу
	}
	fmt.Println(float32(num) + 1.3) // Преобразование целого числа в число с плавающей запятой и вывод результата
}

// Функция stringToFloat преобразует строку в значение с плавающей запятой и выводит результат
func stringToFloat() {
	str := "123.12"
	num, err := strconv.ParseFloat(str, 5) // Преобразование строки в число с плавающей запятой с указанием точности

	if err != nil {
		log.Fatal("не удалось конвертировать тип!") // В случае ошибки выводим сообщение и завершаем программу
	}
	fmt.Println(num + 1) // Сложение числа с плавающей запятой и вывод результата
}

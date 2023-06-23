package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	// количество unicode символов
	fmt.Println("runes:", utf8.RuneCountInString("привет😀")) // 7



	// функции ниже принимают на вход тип rune

	// проверка символа на цифру
	fmt.Println(unicode.IsDigit('1')) // true
	fmt.Println(unicode.IsDigit('g')) // false

	// проверка символа на букву
	fmt.Println(unicode.IsLetter('a')) // true
	fmt.Println(unicode.IsLetter('1'))  // false
	fmt.Println(unicode.IsLetter('\t')) // false

	// проверка символа на нижний регистр
	fmt.Println(unicode.IsLower('A')) // false

	// проверка символа на верхний регистр
	fmt.Println(unicode.IsUpper('A')) // true

	// проверка символа на пробел
	// пробел это не только ' ', но и: '\t', '\n', '\v', '\f', '\r' 
	fmt.Println(unicode.IsSpace('\t')) // true

	// С помощью функции Is можно проверять на кастомный RangeTable:
	// например, проверка на латиницу:
	fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false
	fmt.Println(unicode.Is(unicode.Latin, 'g')) // true

	// функции преобразований
	fmt.Println(string(unicode.ToLower('F'))) // f
	fmt.Println(string(unicode.ToUpper('f'))) // F
	fmt.Println(string(unicode.ToLower('Ф'))) // ф
	fmt.Println(string(unicode.ToUpper('ф'))) // Ф

	
}
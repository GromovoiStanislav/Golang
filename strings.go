package main

import (
	"fmt"
	"strings"
)

func main() {
	q := `
	SELECT *
	FROM person
	WHERE age > 18`
	fmt.Println(q)


	// обрезает символы, переданные вторым аргументом, с обеих сторон строки
	strings.Trim(" hello ", " ") // "hello"
	strings.Trim("_hello_", "_") // "hello"

	// преобразует все буквы в строке в нижний регистр
	strings.ToLower("пРиВеТ") // "привет"

	// преобразует все буквы в строке в верхний регистр
	strings.ToUpper("пРиВеТ") // "ПРИВЕТ"

	// озаглавливает первую букву в каждом слове в строке
	strings.Title("привет, джон") // "Привет, Джон"

	// функция проверяет, что строка "Mr. Doe" начинается с подстроки "Mr."
    strings.HasPrefix("Mr. Doe", "Mr.") // true

	// функция проверяет, что строка "Mr. Doe" заканчивается с подстроки "Doe"
	strings.HasSuffix("Mr. Doe", "Doe")// true

	// замены символов в строке
	strings.ReplaceAll("hello world!", "world!", "buddy!") // hello buddy!
	strings.ReplaceAll("one two three", " ", "_") // one_two_three
	
	strings.Replace("one two three", " ", "_", 1) // one_two three
	strings.Replace("one two three", " ", "_", 2) // one_two_three
}
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

	// убирает пробелы
	strings.TrimSpace(" hello ") // "hello"



	// обрезает символы, переданные вторым аргументом, с обеих сторон строки
	strings.Trim(" hello ", " ") // hello
	strings.Trim("_hello_", "_") // hello
	strings.Trim("tetstet", "te") // s
	strings.Trim("stest", "te") // stes
	strings.Trim("test", "te") // s
	strings.Trim("test", "et") // s


	// Возвращает строку c нижним регистром
	strings.ToLower("пРиВеТ") // привет

	// Возвращает строку c верхним регистром
	strings.ToUpper("пРиВеТ") // ПРИВЕТ

	// озаглавливает первую букву в каждом слове в строке
	strings.Title("привет, джон") // Привет, Джон

	// Начинается ли строка с префикса 
    strings.HasPrefix("Mr. Doe", "Mr.") // true
	strings.HasPrefix("test", "te") // true

	// Количество вхождений 
	strings.Count("hello", string('l')) // 2   тут передается rune !!!
	strings.Count("hello", "H") // 0
	strings.Count("hello", "р") // 1
	strings.Count("hello", "he") //1


	// Заканчивается ли строка суффиксом
	strings.HasSuffix("Mr. Doe", "Doe")// true
	strings.HasSuffix("test", "st") // true

	// заменяет ВСЕ хожденя подстроки old в строке на подстроку new
	strings.ReplaceAll("hello world!", "world!", "buddy!") // hello buddy!
	strings.ReplaceAll("one two three", " ", "_") // one_two_three
	
	// заменяет первые n вхождений подстроки old в строке на подстроку new
	strings.Replace("one two three", " ", "_", 1) // one_two three
	strings.Replace("one two three", " ", "_", 2) // one_two_three
	// Если значение n равно -1, то будут заменены все вхождения
	strings.Replace("one two three", " ", "_", -1) // one_two_three
	strings.Replace("blanotblanot", "not", "***", -1) // bla***bla***


	// Содержится ли подстрока в строке  
	strings.Contains("test", "es") // true

	// Кол-во подстрок в строке
    strings.Count("test", "t") // 2

	// Возвращает начальный индекс подстроки в строке, а при отсутствии вхождения возвращает -1
    strings.Index("test", "e") // 1
    strings.Index("test", "a") // -1


	// Повторяет строку n раз подряд
    strings.Repeat("a", 5) // aaaaa


	// объединяет срез строк через символ
    strings.Join([]string{"hello","world"}, "-") // hello-world
	strings.Join([]string{"hello","world"}, " ") // hello world
	strings.Join([]string{"hello", "world"}, "") // helloworld

	// Разбивает строку согласно разделителю и возвращает срез
    strings.Split("a-b-c-d-e", "-") //  [a b c d e]          - len 5
	strings.Split("a-b-c-d-e", "")  // [a - b - c - d - e]   - len 9
	strings.Split("a-b-c-d-e", "*") //[a-b-c-d-e]            - len 1

	// Разбивает фразу по пробелам и возвращает срез
	strings.Fields("cats are cute") // [cats are cute]
	strings.Fields("Hello world!") // [Hello world!]


	/*
	Builder 
	Когда необходимо собрать большую строку по каким-то правилам, использование конкатенации — не лучшее решение,
	потому что каждая операция создает новую строку, что сильно влияет на производительность при большом количестве операций.
	Такая задача решается с помощью билдера:
	*/
	sb := &strings.Builder{}
	sb.WriteString("hello")
	sb.WriteString(" ")
	sb.WriteString("world")
	fmt.Println(sb.String()) // hello world


}
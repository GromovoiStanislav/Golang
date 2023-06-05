package main

import (
	"fmt"
	"math/rand"
)



func main() {

	//1й вариант:
	var i int
	fmt.Scan(&i)

	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}



	switch a := rand.Intn(10) + 1; a {
    case 9: fmt.Println("a = 9")
    case 8: fmt.Println("a = 8")
    case 7: fmt.Println("a = 7")
    case 6, 5, 4: fmt.Println("a = 6 или 5 или 4")
    default: 
        fmt.Println("значение переменной a не определено")
	}

	company := ""
	switch rand.Intn(3) {
	case 0: company = "Space Adventures"
	case 1: company = "SpaceX"
	case 2: company = "Virgin Galactic"
	}
	fmt.Println(company)


	x := 3
	switch x {
		case 2: 
			x += 1
		case 3:
			x += 1
		case 4:
			x -= 1
	}
	fmt.Println(x) // -> 4



	//2й вариант:
	var c uint32
	fmt.Scan(&c)

	switch {
	case 1 <= c && c <= 9:
		fmt.Println("от 1 до 9")
	case 100 <= c && c <= 250:
		fmt.Println("от 100 до 250")
	case 1000 <= c && c <= 6000:
		fmt.Println("от 1000 до 6000")
	}


	month := "Январь"

	switch {
		case month == "Декабрь" || month == "Январь" || month == "Февраль":
			fmt.Println("Зима")
		case month == "Март" || month == "Апрель" || month == "Май":
			fmt.Println("Весна")
	}





	//Оператор fallthrough
	v := 42

	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}
	// Вывод:
	// 42
	// 1
	// default


	day := 3 // сколько еще рабочих дней осталось до выходных?
	switch day {
		case 1:
			fmt.Println("Понедельник")
			fallthrough
		case 2:
			fmt.Println("Вторник")
			fallthrough
		case 3:
			fmt.Println("Среда")
			fallthrough
		case 4:
			fmt.Println("Четверг")
			fallthrough
		case 5:
			fmt.Println("Пятница")
		default:
			fmt.Println("Неправильный день")
	}
	
	/* выведет
	Среда
	Четверг
	Пятница
	*/



}
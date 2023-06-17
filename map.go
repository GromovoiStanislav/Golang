package main

import "fmt"

/*
Карта (также известна как ассоциативный массив или словарь) 
это неупорядоченная коллекция пар вида ключ-значение. хеш-таблица
*/


func main() {
	// map[string]int // Ключи представляют тип string, значения - тип int
	var people = map[string]int{ 
		"Tom": 1,
		"Bob": 2,
		"Sam": 4,
		"Alice": 8,
	}
	fmt.Println(people)     // map[Tom:1 Bob:2 Sam:4 Alice:8]

	// Для обращения к элементам нужно применять ключи:
	fmt.Println(people["Alice"]) // 8
	fmt.Println(people["Bob"])   // 2
	people["Bob"] = 32
	fmt.Println(people["Bob"]) // 32

	// Для проверки наличия элемента по определенному ключу можно применять выражение if:
	if val, ok := people["Bob"]; ok {
		fmt.Println(val) // 32
	}
	// или
	val, ok := people["Tom"];
	if ok{
		fmt.Println(val) // 1
	}



	// Для перебора элементов применяется цикл for:
	for key, value := range people{
		fmt.Println(key, value)
	}
	// Tom 1
	// Bob 32
	// Sam 4
	// Alice 8


	// Мапы всегда передаются по ссылке:
	modifyMap(people)
    fmt.Println(people) // map[Alice:8 Bob:32 Sam:4 Tom:1 added:100]



	// создание пустой хеш-таблицу
	// var people2 = map[string]int{}
	// people2 := map[string]int{}
	people2 := make(map[string]int)
	
	// добавление элементов
	people2["Kate"] = 128
	fmt.Println(people2) // map[Kate:128]

	//Для удаления применяется встроенная функция delete(map, key)
	delete(people, "Kate")
	fmt.Println(people) // map[]


}


// Мапы всегда передаются по ссылке:
func modifyMap(m map[string]int) {
    m["added"] = 100
}
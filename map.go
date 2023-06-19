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


	///////////////////
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"]) // Lithium

	fmt.Println(elements["Un"]) // пустая строка

	// с проверкой
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}


	elements1 := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements1["Li"]) // Lithium

	//карта строк для карты строк
	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	//тип нашей карты map[string]map[string]string


	if el, ok := elements2["Li"]; ok {
		fmt.Println(el["name"], el["state"]) // Lithium solid
	}




}


// Мапы всегда передаются по ссылке:
func modifyMap(m map[string]int) {
    m["added"] = 100
}
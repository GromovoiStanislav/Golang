package main

import (
	"fmt"
)

// Структура Person, содержащая возраст человека
type Person struct {
	Age uint8
}

// тип PersonList (слайс структур Person)
type PersonList []Person


// возвращает map, где ключ — возраст, а значение — кол-во таких возрастов
func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int)
	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}

func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}

	fmt.Println(pl.GetAgePopularity()) // map[18:2 44:1]

}
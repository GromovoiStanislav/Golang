package main

import "fmt"

func main() {
	var a []int
	var b []int = []int{1, 2, 3}
	c := []int{1, 2, 3}
	d := []int{1: 12}

	fmt.Println(a) // []
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12]


	sl_1 := make([]int, 10)
	fmt.Println(sl_1)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(sl_1)) // 10
	fmt.Println(cap(sl_1)) // 10


	sl_2 := make([]int, 0, 10)
	fmt.Println(sl_2)      // []
	fmt.Println(len(sl_2)) // 0
	fmt.Println(cap(sl_2)) // 10


	sl_3 := make([]int, 5, 10) 
	fmt.Println(sl_3)      // [0 0 0 0 0]
	fmt.Println(len(sl_3)) // 5
	fmt.Println(cap(sl_3)) // 10



	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"} // базовый массив

	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4]  // с 1-го по 4-й
	users3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users1) // [Kate Sam Tom Paul]
	fmt.Println(users2) // [Bob Alice Kate Sam]
	fmt.Println(users3) // [Sam Tom Paul Mike Robert]
	
}


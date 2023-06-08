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


	e := []int{1, 2, 3}
	e = append(e, 4, 5)
	fmt.Println(e) // [1 2 3 4 5]


	sl_1 := make([]int, 10)
	fmt.Println(sl_1)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(sl_1)) // 10
	fmt.Println(cap(sl_1)) // 10
	sl_1 = append(sl_1, 4, 5)
	fmt.Println(sl_1)      // [0 0 0 0 0 0 0 0 0 0 4 5]
	fmt.Println(len(sl_1)) // 12
	fmt.Println(cap(sl_1)) // 20

	sl_2 := make([]int, 5, 10) 
	fmt.Println(sl_2)      // [0 0 0 0 0]
	fmt.Println(len(sl_2)) // 5
	fmt.Println(cap(sl_2)) // 10
	sl_2 = append(sl_2, 4, 5)
	fmt.Println(sl_2)      // [0 0 0 0 0 4 5]
	fmt.Println(len(sl_2)) // 7
	fmt.Println(cap(sl_2)) // 10


	sl_3 := make([]int, 0, 10)
	fmt.Println(sl_3)      // []
	fmt.Println(len(sl_3)) // 0
	fmt.Println(cap(sl_3)) // 10
	sl_3 = append(sl_3, 4, 5)
	fmt.Println(sl_3)      // [4 5]
	fmt.Println(len(sl_3)) // 2
	fmt.Println(cap(sl_3)) // 10

	sl_4 := make([]int, 0, 0)
	fmt.Println(sl_4)      // []
	fmt.Println(len(sl_4)) // 0
	fmt.Println(cap(sl_4)) // 0
	sl_2 = append(sl_4, 4, 5)
	fmt.Println(sl_4)      // [4 5]
	fmt.Println(len(sl_4)) // 2
	fmt.Println(cap(sl_4)) // 2



	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"} // базовый массив
	users1 := initialUsers[2:6] // с элемента с индексом 2 до индекса 6 (не включая 6) 
	users2 := initialUsers[:4]  // с начала до индекса 4 (не включая 4)
	users3 := initialUsers[3:]  // с индекса 3 до конца
	fmt.Println(users1) // [Kate Sam Tom Paul]
	fmt.Println(users2) // [Bob Alice Kate Sam]
	fmt.Println(users3) // [Sam Tom Paul Mike Robert]




	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	baseSlice := baseArray[5:8]

	fmt.Println(baseArray)      // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(baseSlice)      // [5 6 7]
	fmt.Println(len(baseSlice)) // 3
	fmt.Println(cap(baseSlice)) // 5
	pointer := fmt.Sprintf("%p", baseSlice)
	fmt.Println(pointer) // 0xc0000b8028

	baseArray[6] = 100
	fmt.Println(baseArray)      // [0 1 2 3 4 5 100  7 8 9]
	fmt.Println(baseSlice)      // [5 100 7] Изменился
	fmt.Println(len(baseSlice)) // 3
	fmt.Println(cap(baseSlice)) // 5
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice)) // true Тот же адрес

	baseSlice = append(baseSlice, 10)
	fmt.Println(baseArray)      // [0 1 2 3 4 5 100 7 10 9] Изменился
	fmt.Println(baseSlice)      // [5 100 7 10]
	fmt.Println(len(baseSlice)) // 4
	fmt.Println(cap(baseSlice)) // 5
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice)) // true Тот же адрес

	baseSlice = append(baseSlice, 11, 12, 13)
	fmt.Println(baseArray)      // [0 1 2 3 4 5 100 7 10 9] Не изменился!!!
	fmt.Println(baseSlice)      // [5 100 7 10 11 12 13]
	fmt.Println(len(baseSlice)) // 7
	fmt.Println(cap(baseSlice)) // 10
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice)) // false Новый адрес !!!

	baseArray[7] = 200
	fmt.Println(baseArray) // [0 1 2 3 4 5 100 200 10 9]
	fmt.Println(baseSlice) // [5 100 7 10 11 12 13] Не изменился!!!

	
	g := []int{1, 2, 3}
	g = append(g, g...)
	fmt.Println(g) // [1 2 3 1 2 3]


	// Использование append для удаления элемента из среза
	f := []int{1, 2, 3, 4, 5, 6, 7}
	f = append(f[0:2], f[3:]...)
	fmt.Println(f) // [1 2 4 5 6 7]




	s1 := []int{1, 2, 3}
	fmt.Println(s1) // [1 2 3]
	
	s2 := make([]int, 10)
	fmt.Println(s2) // [0 0 0 0 0 0 0 0 0 0]
	n := copy(s2, s1)
	fmt.Println(s2) // [1 2 3 0 0 0 0 0 0 0]
	fmt.Println(n) // 3 - число скопированных элементов:

	s3 := make([]int, 1)
	fmt.Println(s3) // [0]
	n = copy(s3, s1)
	fmt.Println(s3) // [1]
	fmt.Println(n) // 1 - число скопированных элементов:

	s4 := make([]int, len(s1), cap(s1))
	fmt.Println(s4) // [0 0 0]
	n = copy(s4, s1)
	fmt.Println(s4) // [1 2 3]
	fmt.Println(n)  // 3 - число скопированных элементов:



	s_1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(s_1) // [1 2 3 4 5]
	fmt.Println(s_1...) // 1 2 3 4 5

	s_2 := []interface{}{1, "ff", 3, "dd", 5}
	fmt.Println(s_2)    // [1 ff 3 dd 5]
	fmt.Println(s_2...) // 1 ff 3 dd 5

	s_3 := []interface{}{1, 2, s_1, 4, 5}
	fmt.Println(s_3)    // [1 2 [1 2 3 4 5] 4 5]
	fmt.Println(s_3...) // 1 2 [1 2 3 4 5] 4 

	s_4 := []interface{}{1, 2, 3, 4, 5}
	s_4 = append(s_4, s_2, s_1)
	fmt.Println(s_4)    // [1 2 3 4 5 [1 ff 3 dd 5] [1 2 3 4 5]]
	fmt.Println(s_4...) // 1 2 3 4 5 [1 ff 3 dd 5] [1 2 3 4 5]




	a1 := [3]int{1, 2, 3}
	fnA(a1)
	fmt.Println(a) // [1 2 3] Массив Не изменился!!!
	b1 := []int{1, 2, 3}
	fnB(b1)
	fmt.Println(b) // [1 15 3] Срез Изменился!!!

}


func fnA(a [3]int) {
	a[1] = 15
}

func fnB(a []int) {
	a[1] = 15
}
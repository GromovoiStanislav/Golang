package main

import (
	"fmt"
	"sort"
)


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
	fmt.Println(baseSlice[1])      // 6 
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

	

	num := [5]int{0, 2, 4, 6, 8}
	numSlice := num[1:3]
	fmt.Println(numSlice[1]) // 4
	numSlice[0] = 8
	fmt.Println(num) // [0 8 4 6 8]


	// обединение срезов
	g := []int{1, 2, 3}
	g = append(g, g...)
	fmt.Println(g) // [1 2 3 1 2 3]


	// Использование append для удаления элемента из среза
	f := []int{1, 2, 3, 4, 5, 6, 7}
	f = append(f[0:2], f[3:]...)
	fmt.Println(f) // [1 2 4 5 6 7]



	//Копирование слайсов
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




	// перебор их элементов
	s := make([]int, 5)
	s[1] = 2
	s[2] = 3
  
	for i, v := range s {
	  fmt.Println(i, v)
	}
	// 0 0
	// 1 2
	// 2 3
	// 3 0
	// 4 0
	for _, v := range a {
		fmt.Println(v)
	}
	// 0
	// 2
	// 3
	// 0
	// 0




	//Сортировка слайсов
	nums1 := []int{2, 1, 6, 5, 3, 4}
	//быстрая сортировка
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	fmt.Println(nums1) // [1 2 3 4 5 6]

	nums2 := []int{2, 1, 6, 5, 3, 4}
	//сортировка вставками
	sort.SliceStable(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	fmt.Println(nums2) // [1 2 3 4 5 6]



	//Построение многомерных срезов
	seaNames := [][]string{
		{"shark", "octopus", "squid", "mantis shrimp"},
		{"Sammy", "Jesse", "Drew", "Jamie"},
	}

	fmt.Println(seaNames) // [[shark octopus squid mantis shrimp] [Sammy Jesse Drew Jamie]]
	fmt.Println(seaNames[0]) // [shark octopus squid mantis shrimp]
	fmt.Println(seaNames[1]) // [Sammy Jesse Drew Jamie]
	fmt.Println(seaNames[0][0]) // shark
	fmt.Println(seaNames[1][0]) // Sammy



	fmt.Println(MergeNumberLists([]int{1, 2}, []int{3}, []int{4, 5})) // [1 2 3 4, 5]
}


func fnA(a [3]int) {
	a[1] = 15
}

func fnB(a []int) {
	a[1] = 15
}



// принимает вариативный список слайсов чисел и объединяет их в 1, сохраняя последовательность:
func MergeNumberLists(numberLists ...[]int) []int {
	mergedCap := 0
	for i := 0; i < len(numberLists); i++ {
		mergedCap += len(numberLists[i])
	}

	merged := make([]int, 0, mergedCap)
	for _, nl := range numberLists {
		merged = append(merged, nl...)
	}

	return merged
}
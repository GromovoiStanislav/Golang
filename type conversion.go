package main

import (
	"fmt"
)

func main() {
	/*
	Приведение целочисленных типов:
	*/

	var index int8 = 15
	var bigIndex int32
	bigIndex = int32(index)
	fmt.Printf("%v %T\n", bigIndex, bigIndex) // 15 int32

	// По аналогии выше легко понять как конвертировать в другие типы:
	var a int32 = 22
	var b = uint64(a)
	fmt.Printf("%v %T\n", b, b) // 22 uint64


	// Важно: Go позволяет преобразовывать типы с большим количеством бит в типы с меньшим количеством бит. 
	var big int64 = 64
	var little = int8(big)
	fmt.Printf("%v %T\n", little, little) // 64 int8
	
	// Однако, при преобразовании целых чисел, может быть превышено максимальное значение для данного типа данных,
	// и выполнится перенос (потеря данных):
	var big2 int64 = 129
	var little2 = int8(big2)
	fmt.Printf("%v %T\n", little2, little2) // -127 int8

	/*
	Приведение целых чисел и чисел с плавающей точкой
	*/

	var x int64 = 57
	var y = float64(x)
	fmt.Printf("%v %T\n", x, x) // 57 int64
	fmt.Printf("%v %T\n", y, y) // 57 float64

	// Go может преобразовывать float в int. Но делает это с потерей точности. Синтаксис преобразования не меняется.
	var f float64 = 56.231
	var i int = int(f)
	fmt.Printf("%v %T\n", f, f) // 56.231 float64
	fmt.Printf("%v %T\n", i, i) // 56 int

	// Числа, конвертируемые с помощью деления 
	a1 := 5 / 2
	fmt.Printf("%v %T\n", a1, a1) // 2 int
	//Если при делении используются числовые типы с плавающей точкой,
	// тогда все остальные типы будут автоматически объявляться как числа с плавающей точкой:
	a2 := 5.0 / 2
	fmt.Printf("%v %T\n", a2, a2) // 2.5 float64


	//Важно: Если используюьтся переменные то они должны быть 1 типа
	//преобразование происходит к типу превой переменной:
	c1 := 5.0
	d1 := c1 / 2
	fmt.Printf("%T %v\n", d1, d1) // float64 2.5

	c2 := 5
	d2 := c2 / 2.0
	fmt.Printf("%T %v\n", d2, d2) // int 2

	c3 := 5
	d3 := c3 / 2.1
	fmt.Printf("%T %v\n", d3, d3) // (untyped float constant) truncated to int

	c4 := 5.2
	d4 := c4 / 2
	fmt.Printf("%T %v\n", d4, d4) // float64 2.6

	c5 := 5.2 //5
	a5 := 2   //2.5
	d5 := c5 / a5
	fmt.Printf("%T %v\n", d5, d5) // invalid operation: c5 / a5 (mismatched types float64 and int)

}
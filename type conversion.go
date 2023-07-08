package main

import (
	"fmt"
	"strconv"
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


	/*
	Конвертация строк в байты/rune и обратно
	*/

	// Строка в Go это срез байтов, поэтому мы можем конвертировать байты в строку и наоборот:
	str1 := "str"
	bs1 := []byte(str1)
	str2 := string(bs1)
	fmt.Println(str1,str2) // str str
	fmt.Println(bs1) // [115 116 114] - побайтовый срез

	
	text := "строка"
	bs2 := []byte(text)
	fmt.Println(text, string(bs2)) //строка строка
	fmt.Println(bs2) // [209 129 209 130 209 128 208 190 208 186 208 176] - побайтовый срез
	
	// Тоже самое работает и со срезами типа rune:
	rs := []rune(text)
	fmt.Println(text, string(rs)) //строка строка
	fmt.Println(rs) // [1089 1090 1088 1086 1082 1072] - срез рун


	/*
	Конвертация в строки
	*/

	// Itoa
	{
		text := strconv.Itoa(2020)        // int -> string
		fmt.Printf("%T %v\n", text, text) // string 2020

		fmt.Println()
	}

	// FormatInt
	{
		var n int64 = 0xB                     // 'B' в шестнадцатеричной это 11 в десятичной системе
		fmt.Println(strconv.FormatInt(n, 2))  // 1011
		fmt.Println(strconv.FormatInt(n, 10)) // 11
		fmt.Println(strconv.FormatInt(n, 16)) // b

		fmt.Println()
	}

	// FormatUint
	{
		var n uint64 = 93101
		res := strconv.FormatUint(n, 10)
		fmt.Printf("%T %v\n", res, res) // string 93101

		fmt.Println()
	}

	// FormatFloat
	{
		var a float64 = 1.0123456789

		// 1 параметр - число для конвертации
		// fmt - форматирование
		// prec - точность (кол-во знаков после запятой)
		// bitSize - 32 или 64 (32 для float32, 64 для float64)
		fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

		// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
		fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

		// Возможные форматы fmt:
		// 'f' (-ddd.dddd, no exponent),
		// 'b' (-ddddp±ddd, a binary exponent),
		// 'e' (-d.dddde±dd, a decimal exponent),
		// 'E' (-d.ddddE±dd, a decimal exponent),
		// 'g' ('e' for large exponents, 'f' otherwise),
		// 'G' ('E' for large exponents, 'f' otherwise),
		// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
		// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
		var b float64 = 2222 * 1023 * 245 * 2 * 52
		fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10

		fmt.Println()
	}

	// Sprint and Sprintf
	{
		fmt.Println(fmt.Sprint(20.19) + "Попугаев") // 20.19Попугаев

		a := 20.20
		fmt.Println(fmt.Sprintf("%f", a) + "Попугаев") // 20.200000Попугаев

		fmt.Println()
		//Внимание! Использовать fmt для конвертации нежелательно из-за того что производительность ниже по сравнению с strconv.
	}

	// FormatBool
	{
		a := true
		res := strconv.FormatBool(a)
		fmt.Printf("%T %v", res, res) // string true
	}

}


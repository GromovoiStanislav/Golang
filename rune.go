package main

import (
    "fmt"
)

/*
Для работы с Юникод символами отличными от ASCII в Go представлен тип данных rune.
rune — это алиас к int32
Каждая руна представляет собой код символа стандарта Unicode (каждый может занимать до 4-х байт)
*/

func main() {

	emoji := []rune("привет😀")

	fmt.Println(emoji) //[1087 1088 1080 1074 1077 1090 128512]
	fmt.Println(string(emoji)) // привет😀
	fmt.Println(len(emoji)) //7


	for i := 0; i < len(emoji); i++ {
		fmt.Println(emoji[i], string(emoji[i])) // выводим код символа и его строковое представление
	}
	// 1087 п
	// 1088 р
	// 1080 и
	// 1074 в
	// 1077 е
	// 1090 т
	// 128512 😀


	for _, ch := range emoji {
		fmt.Println(ch, string(ch)) // выводим код символа и его строковое представление
	}
	// 1087 п
	// 1088 р
	// 1080 и
	// 1074 в
	// 1077 е
	// 1090 т
	// 128512 😀





	s := "привет😀"

	// range разбирает строку на rune. Т.е. строка автоматически будет преобразована в []rune
	for _, c := range s {
		fmt.Println(c) // выводит код символов Unicode
	}
	// 1087
	// 1088
	// 1080
	// 1074
	// 1077
	// 1090
	// 128512 😀

	for _, c := range s {
		fmt.Printf("%c \n", c)
		//fmt.Println(string(c))
	}
	// п 
	// р 
	// и 
	// в 
	// е 
	// т 
	// 😀



	rs := []rune("Это срез рун")
	// Итерируясь мы будем заменять символ 'р' на '*'
	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		} else if rs[i] == ' ' {
			rs[i] = '_'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs)) //Измененнный срез в виде строки: Это_с*ез_*ун
	fmt.Println(string(rs)) // Это_с*ез_*ун


	// Создадим новый []rune из 10 элементов
	new_rs := make([]rune, 10)
	fmt.Println(new_rs) // [0 0 0 0 0 0 0 0 0 0]
	new_rs = append(new_rs, 'Б')
	fmt.Println(new_rs) // [0 0 0 0 0 0 0 0 0 0 1041]
	new_rs[1] = 'Ю'
	fmt.Println(new_rs) // [0 1070 0 0 0 0 0 0 0 0 1041]



	// Создадим новый пустой []rune
	new_rs2 := make([]rune, 0)
	fmt.Println(new_rs2) // []
	new_rs2 = append(new_rs2, 'Б')
	fmt.Println(new_rs2) // [1041]


	// перевод rune в string
	str := string('H')
	fmt.Println(str)

}
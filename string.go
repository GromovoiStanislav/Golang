package main

import (
    "fmt"
    "unicode/utf8"
)


func main() {
	
	s := "привет"
		
	// range разбирает строку на так называемые руны (rune), т.е. символы в кодировке Unicode (каждый может занимать до 4-х байт)
	for _, c := range s {
		fmt.Println(c) // выводит код символов Unicode
	}
	// 1087
	// 1088
	// 1080
	// 1074
	// 1077
	// 1090

	for _, c := range s {
		fmt.Printf("%c \n", c)
	}
	// п 
	// р 
	// и 
	// в 
	// е 
	// т 


	// При этом индексируется строка s[i] побайтово!!!, а не посимвольно!!!, равно как получение длины строки через len вернёт количество байт, а не рун

	for _, c := range s {
		fmt.Printf("%c(%v) ", c, c)
	}
	// п(1087) р(1088) и(1080) в(1074) е(1077) т(1090) 
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c(%v) ", s[i], s[i])
	}
	// Ð(208) ¿(191) Ñ(209) (128) Ð(208) ¸(184) Ð(208) ²(178) Ð(208) µ(181) Ñ(209) (130) 
	fmt.Println()

	fmt.Println("runes:", utf8.RuneCountInString(s)) // runes: 6
	fmt.Println("bytes:", len(s)) // bytes: 12 

}
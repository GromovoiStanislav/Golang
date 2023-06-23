package main

import (
    "fmt"
    "unicode/utf8"
	"reflect"
)

// byte — это алиас к типу uint8 (0-255)

func main() {

	s := "привет"
		

	// просто байты
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	// 208
	// 191
	// 209
	// 128
	// 208
	// 184
	// 208
	// 178
	// 208
	// 181
	// 209
	// 130

	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
	// Ð
	// ¿
	// Ñ
	// 
	// Ð
	// ¸
	// Ð
	// ²
	// Ð
	// µ
	// Ñ
	// 


	// с ASCII символами
	s2 := "hello"
	for i := 0; i < len(s2); i++ {
		fmt.Println(string(s2[i]))
	}
	// h
	// e
	// l
	// l
	// o


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


	//s := "привет"

	//обычный срез возвращает строку !!!
	fmt.Println(s[:]) // привет
	fmt.Println(s[6:]) // вет
	fmt.Println(s[7:]) // �ет


	str := "hey"
	fmt.Println(str[0], str[1], str[2])   // 104 101 121
	fmt.Println(str[:], str[1:], str[2:]) // hey ey y
	fmt.Println(string(str[0]), string(str[1]), string(str[2])) // h e y  // толко с ASCII символами
	
	fmt.Printf("%T\n", str[0]) // uint8
	fmt.Printf("%T\n", str[:]) // string




	// Байтовй срез ([]byte)
	bs := []byte(s) 
	fmt.Println(bs) // [208 191 209 128 208 184 208 178 208 181 209 130]
	fmt.Println(string(bs)) // привет

	fmt.Printf("Так байтовый срез выглядит внутри: %v\n", bs) // Так байтовый срез выглядит внутри: [208 191 209 128 208 184 208 178 208 181 209 130]
	fmt.Printf("Байтовый срез в виде строки: %s\n", bs) // Байтовый срез в виде строки: привет
	
	bs2 := append(bs, bs...)
	fmt.Println(bs2) //[208 191 209 128 208 184 208 178 208 181 209 130 208 191 209 128 208 184 208 178 208 181 209 130]
	fmt.Printf("Объединенный байтовый срез в виде строки: %s", bs2) //Объединенный байтовый срез в виде строки: приветпривет
	fmt.Println(string(bs2)) // приветпривет

	// Байтовый срез можно изменить,
	// а затем напечатаем его в виде строки
	for i := range bs {
		if bs[i]%2 == 0 {
			bs[i] = bs[i] + 1
			continue
		}
		bs[i] = bs[i] - 1
	}
	fmt.Printf("Измененный байтовый срез в виде строки: %s", bs) //Измененный байтовый срез в виде строки: ѾЁѹѳѴЃ
	fmt.Println(string(bs)) // ѾЁѹѳѴЃ


	/*
	Отдельные ASCII символы можно объявлять сразу с типом byte.
	Для этого нужно обернуть символ в одинарные кавычки и указать тип byte:
	*/
	asciiCh := byte('a')
    asciiChStr := string(asciiCh)

	fmt.Println(reflect.TypeOf(asciiCh))    // uint8
	fmt.Printf("%T\n", asciiCh)				// uint8
	fmt.Println(reflect.TypeOf(asciiChStr)) // string
	fmt.Printf("%T\n", asciiChStr)			// string

	fmt.Println(asciiCh)                    // 97
	fmt.Println(asciiChStr)                 // a

	asciiCh++
	fmt.Println(asciiCh) //98
	fmt.Println(string(asciiCh)) //b

	asciiCh++
	fmt.Println(asciiCh) //99
	fmt.Println(string(asciiCh)) //c



}
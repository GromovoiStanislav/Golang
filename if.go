package main

import (
	"fmt"
	"math"
)

func main() {
	
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	// fmt.Scanln(&a,&b)


	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}


	height := 177;
	if height <= 165 {
		 fmt.Println("Человек невыского роста")
	} else if height > 165 && height <= 185 {
		 fmt.Println("Человек среднего роста")
	} else {
		 fmt.Println("Высокий человек")
	}
	// выведет: Человек среднего роста


	// If с краткой инструкцией
	if v := math.Pow(x, n); v < lim {
        return v
    }else{
		return lim
	}

}
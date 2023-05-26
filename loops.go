package main

import "fmt"


func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}


	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}


	i = 0
	for ; i < 10; {
		fmt.Println(i)
		i++
	}
	// OR
	i = 0
	for i < 10 {
    	fmt.Println(i)
    	i++
	}


	i = 0
	for {
		fmt.Println(i)
		i++
		if i > 10 {
			break            // если число больше 10 выходим из цикла
		}	
	}



	//Оператор continue
	sum_even := 0
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			continue             // число нечетное, то пропускаем его 
		}                        // и переходим к следующей итерации
		sum_even += i
	}
	fmt.Println("Сумма: ",sum_even)        // выведет: 30


	//Оператор break
	for i := 1; i <= 10; i++{
		if i % 5 == 0 {
			break            // если число кратно 5, то выходим из цикла
		}
		fmt.Println(i)
	}
	/* выведет:
	1
	2
	3
	4
	*/



	// Примеры:
	sum := 0
	for i := 0; i <= 30; i += 5 {
		sum += i
	} 
	fmt.Println(sum)
	

	var n int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n){// считываем числа пока не будет введен 0
		fmt.Println(n)
	}

}
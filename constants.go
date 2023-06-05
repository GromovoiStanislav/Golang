package main

import "fmt"


func main() {

	const (
		a = 1
		b
		c
		d = 3
		f
	)
	fmt.Println(a, b, c, d, f)      // 1, 1, 1, 3, 3




	/*
	iota
	*/

	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)
	fmt.Println(Sunday)    // 0
	fmt.Println(Saturday)  // 6
	
	
	
	

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	func main() {
		fmt.Println(Sunday)   // 0
		fmt.Println(Saturday) // 6
	}



	const (
	  c0 = iota  // c0 == 0
	  c1 = iota  // c1 == 1
	  c2 = iota  // c2 == 2
	)
	fmt.Println(c0, c1, c2) // 0 1 2


	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_  // 7
		Add
	)

	fmt.Println(Sunday)   // 0
	fmt.Println(Saturday) // 6
	fmt.Println(Add) // 8



	const (
		u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
		v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
		w         = iota * 42 // w == 84  (индекс - 2, поэтому 2 * 42 = 84)
	)

	// переменные ни в одном блоке const, поэтому индекс не увеличился
	const x = iota  // x == 0
	const y = iota  // y == 0

}
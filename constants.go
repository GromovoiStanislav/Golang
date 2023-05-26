package main

import "fmt"

/*
iota
*/

func main() {

	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)
	fmt.Println(Sunday)    // ����� 0
	fmt.Println(Saturday)  // ����� 6
	
	
	
	

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
		fmt.Println(Sunday)   // ����� 0
		fmt.Println(Saturday) // ����� 6
	}



	const (
	  c0 = iota  // c0 == 0
	  c1 = iota  // c1 == 1
	  c2 = iota  // c2 == 2
	)
	fmt.Println(c0, c1, c2) // �����: 0 1 2


	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_  // ���������� 7
		Add
	)

	fmt.Println(Sunday)   // �����: 0
	fmt.Println(Saturday) // �����: 6
	fmt.Println(Add) // �����: 8



	const (
		u         = iota * 42 // u == 0 (������ - 0, ������� 0 * 42 = 0)
		v float64 = iota * 42 // v == 42.0 (������ - 1, ������� 1.0 * 42 = 42.0)
		w         = iota * 42 // w == 84  (������ - 2, ������� 2 * 42 = 84)
	)

	// ���������� �� � ����� ����� const, ������� ������ �� ����������
	const x = iota  // x == 0
	const y = iota  // y == 0

}
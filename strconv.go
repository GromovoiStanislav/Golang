package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(-42) // "-42"
	fmt.Println(s + "**") // -42**

	i, _ := strconv.Atoi("-42")  // -42
	fmt.Println(i+1) // -41

	n, _ := strconv.Atoi("-42**")  // -42
	fmt.Println(n) // 0
}
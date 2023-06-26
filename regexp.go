package main

import (
	
	"fmt"

	"regexp"
)

func main() {
	password := "qwerty123"

	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{5,}$", password)
	if matched {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}
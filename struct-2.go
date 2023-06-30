package main

import (
	"fmt"
)

// Parent is a parent struct.
type Parent struct {
	Name     string
	Children []Child
}

// Child is a child struct.
type Child struct {
	Name string
	Age  int
}

// создает копию структуры Parent и возвращает ее:
func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}

	cp := *p

	cp.Children = make([]Child, len(p.Children))
	copy(cp.Children, p.Children)

	return cp
}



func main() {
	
	fmt.Println(CopyParent(nil)) // { []}

	
	p := Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}

	cp := CopyParent(&p)

	fmt.Println(p)  // {Harry [{Andy 18}]}
	fmt.Println(cp) // {Harry [{Andy 18}]}

	// при мутациях в копии "cp"изначальная структура "p" не изменяется
	cp.Children[0] = Child{}
	fmt.Println(p)  // {Harry [{Andy 18}]}
	fmt.Println(cp) // {Harry [{ 0}]}
}
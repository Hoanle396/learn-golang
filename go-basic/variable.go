package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	var a string // default = null
	var b int    //default = 0
	var c bool   //default = false

	var d, e, f, g int = 1, 3, 5, 7 // multiple variables
	var (
		h int
		i int    = 1
		k string = "hello"
	)

	// go allow define variable with camel case, pascal case, and snake case

	fmt.Println(a, b, c, d, e, f, g, h, i, k)

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

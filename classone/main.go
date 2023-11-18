package main

import (
	"classone/simplecalc"
	"fmt"
)

const Hello = "Hello World"

func main() {
	a, b := 15, 5
	c := true
	fmt.Println("Hello World")
	fmt.Println(simplecalc.Add(a, b))
	simplecalc.Mul(a, b)
	fmt.Println(c)
	fmt.Println(Hello)

}

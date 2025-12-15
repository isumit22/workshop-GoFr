package main

import "fmt"

func main() {
	var name string
	name = "GoFr Event"
	fmt.Printf("Welcome to %s\n", name)

	var (
		i int     = 20
		f float64 = 3.14
	)

	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("f: %v %T\n", f, f)

	flag := true
	fmt.Printf("b: %v %T\n", flag, flag)

	var x byte = 'A'
	fmt.Printf("x: %v %T\n", x, x)

	var y rune = '中'
	fmt.Printf("y: %v %T\n", y, y)
}
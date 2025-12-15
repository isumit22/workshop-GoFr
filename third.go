package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

func calculateSum(num1, num2 int) int {
	return num1 + num2
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func main() {
	sum := calculateSum(10, 20)
	fmt.Println("Sum:", sum)

	r := rectangle{length: 10, breadth: 20}
	fmt.Println("Area of rectangle:", r.area())
}
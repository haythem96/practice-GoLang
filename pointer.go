package main

import (
	"fmt"
)

func main() {
	x := 15
	a := &x // memory address of x

	fmt.Println("variable a: ", a)
	fmt.Println("content a: ", *a)

	*a = 5

	fmt.Println("variable x: ", x)

	*a = *a * *a

	fmt.Println("variable x:", x)
	fmt.Println("content a: ", *a)
}

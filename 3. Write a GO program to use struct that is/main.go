package main

import (
	"fmt"
	"myproject/week-3/newpack" // use your module name here
)

func main() {
	num := newpack.Number{Value: 5}
	fmt.Println("Square of 5 is:", num.Square())
}

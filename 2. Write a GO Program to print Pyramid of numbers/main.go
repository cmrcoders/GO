// 2. Write a GO Program to print Pyramid of numbers.
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of levels: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for s := 1; s <= n-i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

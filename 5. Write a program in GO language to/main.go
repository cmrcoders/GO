// 5. Write a GO Program to print Floyd’s Triangle.
package main

import "fmt"

func main() {
	var rows int
	number := 1
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)
	fmt.Println("\nFloyd’s Triangle:")
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", number)
			number++
		}
		fmt.Println()
	}
}

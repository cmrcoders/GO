// 1. Write a Go Program to find LCM and GCD of given two numbers.
package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println("GCD:", gcd(a, b))
	fmt.Println("LCM:", lcm(a, b))
}

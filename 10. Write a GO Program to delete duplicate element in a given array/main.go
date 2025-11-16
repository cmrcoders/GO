// 10. Write a GO Program to delete duplicate element in a given array.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 2, 4, 5, 1, 6, 3}
	fmt.Println("Original Array:", arr)
	unique := []int{}
	for _, v := range arr {
		found := false
		for _, u := range unique {
			if v == u {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, v)
		}
	}
	fmt.Println("Array after removing duplicates:", unique)
}

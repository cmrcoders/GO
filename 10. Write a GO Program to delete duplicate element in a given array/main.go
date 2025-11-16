// 10. Write a GO Program to delete duplicate element in a given array.
package main
import "fmt" 
func removeDuplicates(arr []int) []int { 
    exists := make(map[int]bool) 
    result := []int{} 
    for _, value := range arr { 
        if !exists[value] { 
            exists[value] = true 
            result = append(result, value) 
        } 
    } 
    return result 
} 
func main() { 
    arr := []int{1, 2, 3, 2, 4, 5, 1, 6, 3} 
    fmt.Println("Original Array:", arr) 
    arr = removeDuplicates(arr) 
    fmt.Println("Array after removing duplicates:", arr) 
} 

// 6. Write a GO Program to take user input and addition of two strings.
package main 
import "fmt" 
func main() { 
    var first, second string 
    fmt.Print("Enter First String: ") 
    fmt.Scanln(&first) 
    fmt.Print("Enter Second String: ") 
    fmt.Scanln(&second) 
    result := first + second 
    fmt.Println("Concatenation of two strings:", result) 
} 

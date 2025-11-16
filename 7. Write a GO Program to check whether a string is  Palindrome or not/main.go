// 7. Write a GO Program to check whether a string is Palindrome or not.
package main 
import "fmt" 
func main() { 
    var str string 
    fmt.Print("Enter a string: ") 
    fmt.Scanln(&str) 
    original := str 
    reversed := "" 
    for i := len(str) - 1; i >= 0; i-- { 
        reversed += string(str[i]) 
    } 
    if original == reversed { 
        fmt.Printf("\"%s\" is a Palindrome\n", original) 
    } else { 
        fmt.Printf("\"%s\" is NOT a Palindrome\n", original) 
    } 
} 

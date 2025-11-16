// 12. Write a GO program consisting of Contains, Contains Any, Count and Equal Fold String functions.
package main 
import ( 
    "fmt" 
    "strings" 
) 
func main() { 
    fmt.Println("ContainsAny Examples:") 
    fmt.Println(strings.ContainsAny("Germany", "G"))   // true 
    fmt.Println(strings.ContainsAny("Germany", "g"))   // false 
    fmt.Println() 
    fmt.Println("Contains Examples:") 
    fmt.Println(strings.Contains("Germany", "Ger"))    // true 
    fmt.Println(strings.Contains("Germany", "ger"))    // false 
    fmt.Println(strings.Contains("Germany", "er"))     // true 
    fmt.Println() 
    fmt.Println("Count Example:") 
    fmt.Println(strings.Count("cheese", "e"))          // 3 
    fmt.Println() 
    fmt.Println("EqualFold Examples:") 
    fmt.Println(strings.EqualFold("Cat", "cAt"))       // true 
    fmt.Println(strings.EqualFold("India", "Indiana")) // false 
} 

// 9. Write a GO Program to calculate average using arrays.
package main 
import "fmt" 
func main() { 
    var num [100]int 
    var n, sum int 
    fmt.Print("Enter number of elements: ") 
    fmt.Scanln(&n) 
    for i := 0; i < n; i++ { 
        fmt.Printf("Enter number %d: ", i+1) 
        fmt.Scanln(&num[i]) 
        sum += num[i] 
    } 
    avg := float64(sum) / float64(n) 
    fmt.Printf("\nAverage of %d number(s) is: %.2f\n", n, avg) 
} 

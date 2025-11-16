// 11. Write a GO Program with an example of Array reverse sort Functions for integers and strings.
package main 
import ( 
    "fmt" 
    "sort" 
) 
func main() { 
    fmt.Println("Integer Reverse Sort:") 
    nums := []int{50, 90, 30, 10, 50} 
    sort.Sort(sort.Reverse(sort.IntSlice(nums))) 
    fmt.Println("Sorted in descending order:", nums) 
    fmt.Println() 
    fmt.Println("String Reverse Sort:") 
    texts := []string{"Japan", "UK", "Germany", "Australia", "Pakistan"} 
    sort.Sort(sort.Reverse(sort.StringSlice(texts))) 
    fmt.Println("Sorted in descending order:", texts) 
}

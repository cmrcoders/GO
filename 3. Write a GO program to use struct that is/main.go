package main 
import ( 
    "fmt" 
    "newpack" 
) 
func main() { 
    num := newpack.Number{Value: 5}
    result := num.Square()         
    fmt.Println("Square of 5 is:", result) 
} 

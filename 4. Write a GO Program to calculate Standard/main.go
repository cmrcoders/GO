package main  
import (  
    "fmt"  
    "math"  
)  
func main() {  
    var num [10]float64  
    var sum, mean, sd float64  
    fmt.Println("Enter 10 numbers:")  
    for i := 0; i < 10; i++ {  
        fmt.Printf("Value %d: ", i+1)  
        fmt.Scan(&num[i])  
        sum += num[i]  
    }  
    mean = sum / 10  
    for i := 0; i < 10; i++ {  
        sd += math.Pow(num[i]-mean, 2)  
    }  
    sd = math.Sqrt(sd / 10)  
    fmt.Printf("\nStandard Deviation = %.4f\n", sd)  
}  

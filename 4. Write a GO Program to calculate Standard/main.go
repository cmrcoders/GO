package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)
	nums := make([]float64, n)
	var sum, mean, sd float64
	fmt.Println("Enter elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}
	mean = sum / float64(n)
	for i := 0; i < n; i++ {
		sd += math.Pow(nums[i]-mean, 2)
	}
	sd = math.Sqrt(sd / float64(n))
	fmt.Printf("\nStandard Deviation = %.2f\n", sd)
}

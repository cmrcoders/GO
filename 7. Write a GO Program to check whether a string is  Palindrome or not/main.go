// 7. Write a GO Program to check whether a string is Palindrome or not.
package main

import "fmt"

func main() {
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&str)
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	if str == rev {
		fmt.Printf("The string '%s' is a Palindrome.\n", str)
	} else {
		fmt.Printf("The string '%s' is not a Palindrome.\n", str)
	}
}

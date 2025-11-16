// 12. Write a GO program consisting of Contains, Contains Any, Count and Equal Fold String functions.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("HelloWorld", "World")) // true
	fmt.Println(strings.Contains("HelloWorld", "world")) // false
	fmt.Println(strings.ContainsAny("Hello", "aeiou"))   // true
	fmt.Println(strings.ContainsAny("Sky", "aeiou"))     // false
	fmt.Println(strings.Count("cheese", "e"))            // 3 (true means >0)
	fmt.Println(strings.Count("cheese", "z"))            // 0 (false means none found)
	fmt.Println(strings.EqualFold("GoLang", "golang"))   // true
	fmt.Println(strings.EqualFold("Hello", "World"))     // false
}

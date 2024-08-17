package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// original string value remains unchanged
	// fmt.Println("original string value=", greeting)

	ages := []int{20, 65, 14, 89, 63, 55}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 66)
	fmt.Println(index)

	names := []string{"rakib", "jewel", "nizam", "kamrul"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "rakib"))
}

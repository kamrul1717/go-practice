package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	firstName1, SecondName1 := getInitials("kamrul hasan")
	fmt.Println(firstName1, SecondName1)
	firstName2, SecondName2 := getInitials("rakib hasan")
	fmt.Println(firstName2, SecondName2)
	firstName3, SecondName3 := getInitials("nizam")
	fmt.Println(firstName3, SecondName3)
}

package main

import "fmt"

func updateName(x *string) {
	*x = "kamrul"
}

func main() {
	name := "rakib"

	//updateName(name)

	fmt.Println("memory address of name is:", &name)

	m := &name

	fmt.Println("memory address of m is:", m)
	fmt.Println("value at memory address:", *m)
	updateName(m)
	fmt.Println(name)

}

package main

import "fmt"

func main() {
	name := "kamrul"
	age := 30

	// Print Function
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("new line\n")

	// Println Function
	fmt.Println("Hello guys!")
	fmt.Println("Goodbye guys!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.2f points\n", 225.275)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("the saved string is", str)

}

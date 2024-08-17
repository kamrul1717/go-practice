package main

import "fmt"

func main() {
	mybill := newBill("kamrul's bill")
	fmt.Println(mybill.format())
}

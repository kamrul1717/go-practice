package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		9595985: "kamrul",
		959595:  "rakib",
		95955:   "jewel",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[95955])

	phonebook[9595985] = "nizam"
	fmt.Println(phonebook)

	phonebook[959595] = "mahid"
	fmt.Println(phonebook)

}

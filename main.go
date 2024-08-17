package main

import "fmt"

func updateName(x string) string {
	x = "kamrul"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	name := "rakib"

	name = updateName(name)

	fmt.Println(name)

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 7.93,
	}

	updateMenu(menu)
	fmt.Println(menu)

}

package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("kamrul")
	// sayGreeting("rakib")
	// sayBye("nizam")

	cycleNames([]string{"kamrul", "rakib", "mahid"}, sayGreeting)
	cycleNames([]string{"kamrul", "rakib", "mahid"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(55)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %.3f and circle 2 is %.3f\n", a1, a2)
}

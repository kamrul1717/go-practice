package main

import "fmt"

func main() {
	// Arrays
	// var ages []int = []int{30, 28, 29} // auto infer length of the array
	// var ages [3]int = [3]int{30, 28, 29}
	// var ages = [3]int{30, 28, 29}
	// var ages = []int{30, 28, 29}
	ages := []int{30, 28, 29}
	names := [4]string{"kamrul", "Jewel", "Nizam", "Rakib"}
	names[1] = "Mahid"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{100, 89, 90}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "x")
	fmt.Println(rangeOne)

}

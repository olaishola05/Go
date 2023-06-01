package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4}
	numn2 := []int{5, 6, 9}

	combineSlices := append(nums, numn2...)

	sum := sum(combineSlices...)
	fmt.Println(sum)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("addition of two slices")
	num1 := []int{1, 2, 3, 4, 5, 6, 7}
	num2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("the addition of two slices", Add(num1, num2))

}
func Add(num1 []int, num2 []int) []int {
	c := make([]int, len(num1))
	sort.Sort(sort.Reverse(sort.IntSlice(num2)))
	for i, _ := range num1 {
		c[i] = num1[i] + num2[i]

	}
	return c

}

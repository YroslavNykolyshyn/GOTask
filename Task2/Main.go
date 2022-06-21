package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var myInt []int
	min, max := findMinAndMax(myInt)
	myStrSlice := strings.Split(input, " ")
	for _, i := range myStrSlice {
		myInt, _ := strconv.Atoi(i)
		fmt.Println(myInt)
	}
	fmt.Println(myInt)

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func findMinAndMax(myInt []int) (min int, max int) {
	min = myInt[0]
	max = myInt[0]
	for _, value := range myInt {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

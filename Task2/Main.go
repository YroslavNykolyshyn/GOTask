package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var res string
	var min int32
	var max int32
	myStrSlice := strings.Split(input, " ")

	for i := 0; i < len(myStrSlice); i++ {

		IntVal, _ := strconv.Atoi(myStrSlice[i])

		if i == 0 {
			max = int32(IntVal)
			min = int32(IntVal)
		}

		if int32(IntVal) > max {
			max = int32(IntVal)
		}

		if int32(IntVal) < min {
			min = int32(IntVal)
		}

	}

	if max != min {
		res = fmt.Sprintf("Max: %d Min: %d", max, min)
	} else {
		res = fmt.Sprintf("%d", max)
	}

	fmt.Println(res)

}

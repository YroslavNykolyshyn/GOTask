package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3}
	var n = 0
	arr = append(arr[:n], arr[n+1:]...)
	var n1 = 1
	arr = append(arr[:n1], arr[n1+1:]...)
	var n2 = 3
	arr = append(arr[:n2], arr[n2+1:]...)
	fmt.Println(arr)
}

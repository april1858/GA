package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	indexSmallest := 0
	for i, _ := range arr {
		if arr[i] < smallest {
			indexSmallest = i
		}
	}
	return indexSmallest
}

func main() {
	a := []int{4, 3, 7, 5, 6, 9}

	newArray := make([]int, 0)

	for range a {
		smallest := findSmallest(a)

		newArray = append(newArray, a[smallest])

		a[smallest] = a[0]
		a = a[1:]
	}
	fmt.Println("new array - ", newArray, "old array - ", a)
}

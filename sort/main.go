package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	indexSmallest := 0
	for i, _ := range arr {
		fmt.Println("arr - ", arr, "arr[i] - ", arr[i], "i - ", i)
		if arr[i] < smallest {
			smallest = arr[i]
			indexSmallest = i
		}
	}
	return indexSmallest
}

func main() {
	a := []int{4, 3, 1, 2}

	newArray := make([]int, 0)

	for range a {
		smallest := findSmallest(a)

		fmt.Println("smallest = ", smallest)

		newArray = append(newArray, a[smallest])

		copy(a[smallest:], a[smallest+1:])

		//a[len(a)-1] = 0
		a = a[:len(a)-1]
	}
	fmt.Println("new array - ", newArray, "old array - ", a)
}

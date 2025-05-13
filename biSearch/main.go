package main

import "fmt"

func binSearch(list []int, target int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == target {
			return mid
		}
		if list[mid] < target {
			low = mid + 1
		}
		if list[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	l := []int{1, 2, 4, 6, 7, 9}
	t := 3

	fmt.Println(binSearch(l, t))
}

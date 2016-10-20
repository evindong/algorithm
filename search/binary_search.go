package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	log.Println(search(arr, 7))
}

func search(arr []int, target int) int {
	if len(arr) <= 0 {
		return -1
	}

	i, j := 0, len(arr)-1
	for i <= j && i >= 0 && j <= len(arr)-1 {
		mid := (i + j) / 2
		if target < arr[mid] {
			j = mid - 1
		} else if target > arr[mid] {
			i = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

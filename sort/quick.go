package main

import (
	"fmt"
)

func main() {

	arr := []int{6, 4, 5, 2, 1}
	quickSort(arr, 0, len(arr))
	for _, item := range arr {
		fmt.Print(item, " ")
	}
}

//5,4,3,2,1
func quickSort(arr []int, start, end int) {

	mid := getMid(arr, start, end)
	if mid-start > 1 {
		quickSort(arr, start, mid)
	}

	if end-(mid+1) > 1 {
		quickSort(arr, mid+1, end)
	}
}

func getMid(arr []int, start, end int) int {

	sign := arr[start]

	i := start + 1
	j := end - 1

	for i < j {
		for arr[i] < sign && i < end {
			i++
		}
		for arr[j] >= sign && j > 0 {
			j--
		}
		if i != j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	arr[start] = arr[i]
	arr[i] = sign

	return i
}

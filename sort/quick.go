package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 4, 3, 2, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	for _, item := range arr {
		fmt.Print(item, " ")
	}
}

func quickSort(arr []int, start, end int) {
	mid := getMid(arr, start, end)
	if mid > start {
		quickSort(arr, start, mid-1)
	}
	if end > mid {
		quickSort(arr, mid+1, end)
	}
}

func getMid(arr []int, start, end int) int {
	sign := arr[start]
	i := start
	j := end
	for i < j {
		for i < j && arr[j] >= sign {
			j--
		}
		for i < j && arr[i] <= sign {
			i++
		}
		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	arr[start] = arr[i]
	arr[i] = sign

	return i
}

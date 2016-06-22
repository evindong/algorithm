package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	insertSort(arr)
	for _, a := range arr {
		fmt.Print(" ", a)
	}
}

func insertSort(arr []int) {
	//5 6 7 4 3 2 1
	for i := 1; i < len(arr); i++ {
		j, temp := i, arr[i]
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
}

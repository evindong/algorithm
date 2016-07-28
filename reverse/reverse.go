package main

import "fmt"

func reverse(arr []int, start, end int) {
	len := end - start + 1
	if len < 2 {
		return
	}
	halfLen := len / 2
	j := end
	for i := start; i < halfLen+start; i++ {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		j--
	}
}

func reverseBySep(arr []int, separate int) {
	if separate > len(arr) && separate < 1 {
		return
	}
	reverse(arr, 0, separate-1)
	reverse(arr, separate, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverseBySep(arr, 4)
	print(arr)

	arr = []int{1, 2}
	reverseBySep(arr, 1)
	print(arr)

	arr = []int{1, 2, 3, 4, 5}
	reverseBySep(arr, 1)
	print(arr)

	arr = []int{1, 2, 3, 4}
	reverseBySep(arr, 3)
	print(arr)
}

func print(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], ",")
	}
	fmt.Println()
}

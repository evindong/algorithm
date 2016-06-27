package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 11, 13, 45, 3, 2, 8, 1, 7, 7, 7, 7}
	sort(arr, 0, len(arr))
	for _, a := range arr {
		fmt.Print(a, " ")
	}
}

func sort(arr []int, start, end int) {
	if end-start > 1 {
		mid := (end + start) / 2
		sort(arr, start, mid)
		sort(arr, mid, len(arr))
		merge(arr, start, mid, end)
	}
}

func merge(arr []int, start, mid, end int) {
	l := end - start
	temp := make([]int, l)
	i := start
	j := mid
	t := 0

	f := func(temp, arr []int, t, k int) (int, int) {
		temp[t] = arr[k]
		t++
		k++
		return t, k
	}

	for i < mid && j < end {
		if arr[i] < arr[j] {
			t, i = f(temp, arr, t, i)
		} else if arr[i] > arr[j] {
			t, j = f(temp, arr, t, j)
		} else {
			t, i = f(temp, arr, t, i)
			t, j = f(temp, arr, t, j)
		}
	}
	for i < mid {
		t, i = f(temp, arr, t, i)
	}
	for j < end {
		t, j = f(temp, arr, t, j)
	}
	j = 0
	for i := start; i < end; i++ {
		arr[i] = temp[j]
		j++
	}
}

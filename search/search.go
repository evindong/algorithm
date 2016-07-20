package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func isHigh(t, n uint) bool {
	return (t & (1 << n)) != 0
}

func generateRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}

	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func initArray() []int {
	return generateRandomNumber(0, 256, 200)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			arr := initArray()
			sign := search(arr)
			if sign == -1 {
				fmt.Println("not found")
				return
			}
			for _, item := range arr {
				if item == sign {
					fmt.Println("error")
					return
				}
			}
			fmt.Println("test passed sign is", sign)
		}()
	}
	wg.Wait()
}

func search(arr []int) int {
	total := 0
	for level := 7; level >= 0; level-- {
		hItem := []int{}
		lItem := []int{}
		for _, item := range arr {
			if isHigh(uint(item), uint(level)) {
				hItem = append(hItem, item)
			} else {
				lItem = append(lItem, item)
			}
		}
		if len(hItem) == 0 {
			total += (1 << uint(level))
			return total
		}
		if len(lItem) == 0 {
			return total
		}
		if len(hItem) <= len(lItem) {
			total += (1 << uint(level))
			arr = hItem
		} else {
			arr = lItem
		}
	}
	return -1
}

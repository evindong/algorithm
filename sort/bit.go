package main

import "fmt"

type bitSet struct {
	b []byte
}

func NewBitSet(len int) *bitSet {
	b := make([]byte, len/8+1)
	return &bitSet{
		b: b,
	}
}

func (b *bitSet) Set(n int) {
	start, offset := n/8, n%8
	b.b[start] |= 1 << uint(offset)
}

func (b *bitSet) test(n int) bool {
	start, offset := n/8, n%8
	return (b.b[start] & (1 << uint(offset))) != 0
}

func main() {
	arr := []int{23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 6, 5, 4, 3, 2, 1}
	bs := NewBitSet(23)
	for i := 0; i < len(arr); i++ {
		bs.Set(arr[i])
	}

	for i := 0; i < 24; i++ {
		if bs.test(i) {
			fmt.Print(i, " ")
		}
	}
}

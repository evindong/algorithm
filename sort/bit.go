package main

import (
	"fmt"
	"math"
)

type BitField []byte

func NewBitField(len int) BitField {
	n := int(math.Ceil(float64(len) / 8.0))
	return BitField(make([]byte, n))
}

func (b BitField) Len() int {
	return len(b) * 8
}

func (b BitField) Clear(n int) {
	start, offset := n/8, n%8
	b[start] &= ^(1 << uint(offset))

}

func (b BitField) Set(n int) {
	start, offset := n/8, n%8
	b[start] |= 1 << uint(offset)
}

func (b BitField) Test(n int) bool {
	start, offset := n/8, n%8
	return (b[start] & (1 << uint(offset))) != 0
}

func main() {
	arr := []int{23, 22, 21, 20, 17, 18, 19, 16, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 26, 29}
	bs := NewBitField(len(arr))
	for i := 0; i < len(arr); i++ {
		bs.Set(arr[i])
	}

	for i := 0; i < bs.Len(); i++ {
		if bs.Test(i) {
			fmt.Print(i, " ")
		}
	}
}

package main

import "fmt"

type BitField []byte

func NewBitField(len int) BitField {
	return BitField(make([]byte, len/8+1))
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
	arr := []int{23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	bs := NewBitField(23)
	for i := 0; i < len(arr); i++ {
		bs.Set(arr[i])
	}

	for i := 0; i < 24; i++ {
		if bs.Test(i) {
			fmt.Print(i, " ")
		}
	}
}

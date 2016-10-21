package main

import (
	"fmt"
	"log"
)

type Btree struct {
	Left *Btree
	Item int
	Right *Btree
}

func (btree *Btree) iterator() {
	if btree == nil {
		return
	}
	if btree.Left != nil {
		btree.Left.iterator()
	}
	fmt.Print(btree.Item," ")
	if btree.Right != nil {
		btree.Right.iterator()
	}
}

func (btree *Btree) add(i int) {
	if btree != nil {
		if btree.Item > i {
			if btree.Left != nil {
				btree.Left.add(i)
			}else {
				bt := &Btree{
					Item:i,
				}
				btree.Left = bt
			}
		} else {
			if btree.Right != nil {
				btree.Right.add(i)
			}else {
				bt := &Btree{
					Item:i,
				}
				btree.Right = bt
			}
		}
	}
}

func (btree *Btree) search(i int) bool{
	if btree == nil {
		return false
	}
	if btree.Item > i {
		return btree.Left.search(i)
	}
	if btree.Item < i {
		return btree.Right.search(i)
	}
	return true

}

func main() {
	btree := &Btree {
		Item:6,
	}
	arr := []int{6,5,5,1,2,11,4,3,2,1}
	for _,i := range arr {
		btree.add(i)
	}
	log.Println("---------")
	btree.iterator()
	log.Println("------------")
	log.Println(btree.search(11))
	log.Println(btree.search(5))
	log.Println(btree.search(1))
	log.Println(btree.search(6))
	log.Println(btree.search(2))
	log.Println(btree.search(19))
}


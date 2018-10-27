package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello tree")

	t := NewT(50)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		num := rand.Intn(100)
		fmt.Printf("%d ", num)
		t.InsertT(num)
	}
	fmt.Printf("\n")

	t.PrtT()
}

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func NewT(v int) *Tree {
	return &Tree{v, nil, nil}
}

func (t *Tree) InsertT(v int) {

	node := Tree{}
	node.Val = v
	cur := t.FindT(&v)
	if cur.Val == v {
		fmt.Println("Value is equal ", v)
	} else if v < cur.Val { //如果当前值小于节点值，则给左孩子赋值
		cur.Left = &node
	} else { //如果当前值大于节点值，则给右孩子赋值
		cur.Right = &node
	}

}

func (t *Tree) FindT(v *int) *Tree {
	if *v < t.Val {
		if t.Left == nil {
			return t
		} else {
			return t.Left.FindT(v) // 如链表一样，在下一节点继续查找；
		}
	}

	if *v > t.Val {
		if t.Right == nil {
			return t
		} else {
			return t.Right.FindT(v) // 如链表一样，在下一节点继续查找；
		}
	}

	return t
}

func (t *Tree) PrtT() {
	//in-order
	if t.Left != nil {
		t.Left.PrtT()
	}
	fmt.Printf("%d ", t.Val)
	if t.Right != nil {
		t.Right.PrtT()
	}
}

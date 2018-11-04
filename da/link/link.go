package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func IsEmpty(l *link);
func IsLast(e int, l *link)
func Find(e int, l *link) *link
func Delete(e, l *link)
func Insert(e int, l *link)
func Append(e int, l *link)

head->n1->n2->n3

1. link的关注点， 在进入函数的时候需要重新定义一个link指针；
2. 增加一个head进行判断，会省去head节点的判断， 浪费一个节点；
3.

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l := new(ListNode)
	var total int
	CalNode := l1
	newl := l

	for {
		node := new(ListNode)
		if total == 0 {
			l = node
		} else {
			l.Next = node
		}
		total++

		if CalNode.Next == nil {
			break
		}
		CalNode = CalNode.Next

	}

	l = newl
	for i := 0; i < total; i++ {
		val := (newl.Val + l1.Val + l2.Val) % 10
		flg := (newl.Val + l1.Val + l2.Val) / 10

		newl.Val = val

		if l.Next != nil && flg > 0 {
			l.Next.Val = 1
		}
		l1 = l1.Next
		l2 = l2.Next
		l = newl.Next
	}
	return l
}

type link struct {
	val  int
	next *link
}

func main() {
	var flag bool

	fmt.Println("hello link", flag)

	head := new(link)
	l := head

	// for i := 0; i < 10; i++ {
	// 	num := rand.Intn(100)
	// 	l.Append(num)
	// }

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		num := rand.Intn(100)
		l.Insert(num)
	}

	head.Prt()

	head.Delete(15)
	head.Delete(20)
	head.Delete(25)

	l1 := link{0, nil}
	l2 := link{0, nil}
	l1.Append(5)
	l1.Append(3)
	l1.Append(3)
	l2.Append(6)
	l2.Append(6)
	l2.Append(3)
	l2.Append(7)
	l2.Append(9)

	l1.Prt()
	l2.Prt()

	l3 := addTwoNumbers1(&l1, &l2)
	l3.Prt()

}

func addTwoNumbers1(l1 *link, l2 *link) *link {

	l := link{0, nil}
	node1 := l1.next
	node2 := l2.next
	newl := &l
	var extra, n1, n2 int

	for {
		if node1 == nil && node2 == nil {
			break
		}
		if node1 != nil {
			n1 = node1.val
		} else {
			n1 = 0
		}
		if node2 != nil {
			n2 = node2.val
		} else {
			n2 = 0
		}
		sum := (n1 + n2 + extra) % 10
		extra = (n1 + n2 + extra) / 10

		newl.val = sum

		//if extra > 0 || (extra == 0 && (node1.next != nil || node2.next != nil)) {
		//如果extra> 0， 表示需要开一个节点；
		//如果extra == 0， 需要看node1、node2中是否有下一节点；

		fmt.Println("extra", extra)

		if extra > 0 {
			node := link{}
			node.val = extra
			newl.next = &node
			newl = newl.next
		} else if node1 != nil && node1.next != nil {
			node := link{}
			node.val = extra
			newl.next = &node
			newl = newl.next
		} else if node2 != nil && node2.next != nil {
			node := link{}
			node.val = extra
			newl.next = &node
			newl = newl.next
		}
		if node1 != nil {
			node1 = node1.next
		}
		if node2 != nil {
			node2 = node2.next
		}

	}
	return &l
}

func (l *link) Append(v int) {
	n := link{v, nil}

	last := l
	if last.next == nil {
		last.next = &n
	} else {
		for last.next != nil {
			last = last.next
		}
		last.next = &n
	}
}

func (l *link) Insert(v int) {

	n := link{v, nil}

	ins := l
	flag := false // 初始未插入

	if ins.next == nil {
		ins.next = &n
	} else {
		for ins.next != nil {
			if n.val < ins.next.val { // 小于后节点，插入
				n.next = ins.next
				ins.next = &n
				flag = true
				break
			} else {
				ins = ins.next
			}
		}
		if flag == false {
			ins.next = &n
		}
	}
}

func (l *link) Prt() {
	prt := l

	for prt.next != nil {
		fmt.Printf("%d ", prt.val)
		prt = prt.next
	}
	fmt.Println(prt.val)

}

func (l *link) IsEmpty() bool {
	if l.next != nil {
		return false
	}
	return true
}

func (l *link) Delete(v int) {
	cur := l

	if cur.IsEmpty() == true {
		return
	}

	for cur.next != nil {
		if cur.next.val == v {
			fmt.Println("Delete value ", v)
			if cur.next.next != nil {
				cur.next = cur.next.next
			} else {
				cur.next = nil
			}
		}
		cur = cur.next
	}
}

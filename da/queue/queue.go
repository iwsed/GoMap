package main

/*
循环队列，主要是postion + N进行循环；
1-2-3-4-5
|       |
<--------
*/

import "fmt"

type Queue struct {
	Size  int
	Front int
	End   int
	Elem  []int
}

const FullQ = 4

func main() {
	fmt.Println("hello Q")
	q := NewQ(FullQ)

	q.EnQ(10)
	fmt.Println(q.Front, q.End, q.Size)

	q.EnQ(11)
	fmt.Println(q.Front, q.End, q.Size)

	q.EnQ(13)
	fmt.Println(q.Front, q.End, q.Size)

	q.EnQ(14)
	q.EnQ(15)
	//_ = q.DeQ()
	fmt.Println(q.Front, q.End, q.Size)
	_ = q.DeQ()
	fmt.Println(q.Front, q.End, q.Size)

	q.EnQ(16)

	q.PrtQ()
}

func (q *Queue) PrtQ() {

	for i := 0; i < q.Size; i++ {

		if q.Front > q.End { // 进入循环

		}
		pos := (i + q.Front + 1) % FullQ
		fmt.Printf("%d ", q.Elem[pos])
	}
}

func NewQ(num int) *Queue {
	Q := Queue{0, 0, 0, make([]int, num)}

	return &Q
}

func (q *Queue) EnQ(val int) {
	if q.IsFullQ() {
		return
	}

	q.End = (q.End + 1) % FullQ
	q.Elem[q.End%FullQ] = val
	q.Size++
}

func (q *Queue) DeQ() int {

	if q.IsEmptyQ() {
		return 0
	}

	q.Front = (q.Front + 1) % FullQ
	val := q.Elem[q.Front%FullQ]
	q.Size--
	return val

}

func (q *Queue) IsEmptyQ() bool {
	if q.Size == 0 {
		return true
	}
	return false
}

func (q *Queue) IsFullQ() bool {
	if FullQ == q.Size {
		fmt.Println("Q is full")
		return true
	}
	return false
}

package main

import "fmt"

func main() {
	fmt.Println(13%10, 13/10)

}

func twoSum(nums []int, target int) []int {
	t := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				t[0] = i
				t[1] = j
				return t
			}
		}
	}
	return t
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode

	for l1.Next != nil {
		val := l1.Val + l2.Val

		l := new(ListNode)
		l.Val = val % 10

		if l1.Next.Next != nil {
			l3.Next.Val = val / 10
		}
		l3.Next = l

	}
	return l3

}

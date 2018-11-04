# 删除倒数第N个节点

给出一个链表，在倒数第N个节点删除， 两种方法

1. 暴力法， 计算出link的节点数
2. 窗口法， 定义两个pointer， 进行窗口位移

就是计算link的编写能力， 一则看link的头节点， 另外一个数判断节点的最后一个值，同时判断链的边界， 比如：

链表长度为1， 返回 空，  链表长度为2，删除一个节点等等；


代码如下:

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	first := head
	second := head
	last := head

	if head.Next == nil {
		return nil
	}

	for i := 0; i < n; i++ {

		if first.Next == nil { // n euqal num of link
			return head.Next
		}
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	// remove second node
	second.Next = second.Next.Next

	return last
}

```
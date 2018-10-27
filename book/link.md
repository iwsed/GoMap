# golang实现link

使用golang实现single link过程， 许久没有编写数据结构的实现方式， 编写过程需要关注

1. link的关注点， 在进入函数的时候需要重新定义一个link指针；
2. 增加一个head进行判断，会省去head节点的判断， 浪费一个节点；
3. link的next及next的判断很关键；如果判断失误会出现for的死循环；
4. 采用对象的方式编码，先创建了一个head, 实现的功能包括:, 
    - (l *link)IsEmpty  
    - (l *link)Delete()     
    - (l *link)Insert(v int)    link有序存放；
    - (l *link)Append(v int)
5. 在link上实现"addTwoNumbers1"
6. 自己编写的case， 最好首字母是大写，这样就可以直接cp到leetcode上了；
7. 在本子上写出可能需要创建新的下一个节点的条件
    - extra > 0
    - node1 不为 nil 同时存在下一节点不为 nil的情况；
    - node2 不为 nil 同时存在下一节点不为 nil的情况；

![](add_two_number.png)

```go

type link struct {
	val  int
	next *link
}


func main() {
	fmt.Println("hello link")

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

	head.Prt()

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
	prt := l.next

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

```


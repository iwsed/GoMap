# 栈的Golang实现过程

自己写了一个Stack的实现方式， 三种方法

## 自己写的
通过golang的slice方式实现， 偏向结构化的方式实现；

```go

func PushS(s *[]int, v int) {
	fmt.Printf("%p, %p\n", s, *s)
	*s = append(*s, v)
}

func PopS(s *[]int) int {
	if len(*s) == 0 {
		return 0
	}
	len := len(*s)
	fmt.Println("len of stack", len)
	v := (*s)[len-1]
	*s = (*s)[:len-1]

	return v
}

func main(){
    stack := make([]int, 0)
    PushS(&stack, 5)
    PushS(&stack, 10)

    v := PopS(&stack)
    fmt.Println(v, stack)

}
```

## 参考网络一
通过对象的方式实现

```go
type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

func main(){
    s := make(stack, 0)
	s = s.Push(10)
	s = s.Push(11)
	s = s.Push(12)
	fmt.Println(s)
	s, v := s.Pop()
	s, v = s.Pop()
	fmt.Println(s, v)
	s, v = s.Pop()
	fmt.Println(s, v)
}

```

## 参考网络资源二
增加lock模式 这样在多协程的方式下也能正常处理；

```go
type stk struct {
	lock sync.Mutex
	s    []int
}

func NewStk() *stk {
	return &stk{sync.Mutex{}, make([]int, 0)}
}

func (s *stk) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stk) Pop() (int, error) {
	s.lock.Lock()
	s.lock.Unlock()

	len := len(s.s)
	if len == 0 {
		return -1, errors.New("Empty Stack")
	}
	v := s.s[len-1]
	s.s = s.s[:len-1]
	return v, nil
}

func main() {
	fmt.Println("hello go")

	s := NewStk()
	s.Push(10)
	s.Push(11)
	s.Push(12)
	fmt.Println(s.s)
	vv, _ := s.Pop()
	vv, _ = s.Pop()
	fmt.Println(s.s, vv)
}

```

## godoc的参考代码实现

```go
type (
    Stack struct {
        top *node
        length int
    }
    node struct {
        value interface{}
        prev *node
    }
)
// Create a new stack
func New() *Stack {
    return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
    return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
    if this.length == 0 {
        return nil
    }
    return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
    if this.length == 0 {
        return nil
    }

    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
    n := &node{value,this.top}
    this.top = n
    this.length++
}
```
# 原子操作

直接看atomic的AddInt32不太好理解， 增加一个case进行操作就可以熟练了解；

- TOTAL基数较小的时候，并不存在问题；
- TOTAL较大的时候就可以看出，多并发的情况下对共享资源的争抢就会导致增加失败；
- LoadInt32 在等待当前有在对该变量执行的go，如果routine一直对某个值在写，则响应的routine一直在等待呢？
- atomic.CompareAndSwapInt32, 该值可以在AddInt32的同时进行判断，是否进行追加交互处理；

```go
atomic.AddInst32(old int32, delta)
```

示范代码

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
const TOTAL = 10000

var wg sync.WaitGroup

func main() {
	fmt.Println("hello atomic")
	var num,num1 int32
	wg.Add(1)

	for i := 0; i < TOTAL; i++ {
		go AddNum(&num)
		go AddNumD(&num1)
	}

    time.Sleep(time.Second * 2)
    
    

	fmt.Println(num, num1)
	wg.Wait()

}

func AddNum(num *int32) {
	atomic.AddInt32(num, 1)
}

func AddNumD(num *int32) {
	*num++
}

```
package main

import (
	"fmt"
	"sync"
	"time"
)

// 测试string的chan功能
var (
	wg    sync.WaitGroup
	Num   int
	mutex sync.Mutex
)

// NUM total num
const NUM = 10000

func main() {
	fmt.Println("hello world")
	wg.Add(3)

	msg := make(chan string, 100)
	fmt.Printf("%p\n", msg)

	go productor(msg)

	for i := 0; i < 20; i++ {
		go consumer(msg, i)
	}
	go prtChan(msg)

	fmt.Println("main thread end")
	wg.Wait()
}

func prtChan(msg chan string) {
	var i int

	for {
		time.Sleep(time.Millisecond * 10)
		if i%100 == 0 {
			fmt.Println("Num", Num, "buffered cap: ", time.Now(), len(msg))
		}
		i++
	}
}
func productor(msg chan string) {

	for i := 0; i < NUM; i++ {
		msg <- "hello world every where"
		//fmt.Println("msg put: ", time.Now())
		time.Sleep(time.Millisecond * 1)
	}

}

func consumer(msg chan string, node int) {
	var i int

	for ; ; i++ {

		str := <-msg
		if i%1000 == 0 {
			fmt.Println("msg get: ", len(str), time.Now(), str, node)
		}
		mutex.Lock()
		{
			Num++
		}
		mutex.Unlock()
		time.Sleep(time.Millisecond * 10)
	}
}

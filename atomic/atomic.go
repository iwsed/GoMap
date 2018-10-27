package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const TOTAL = 1000000
const DEV = 100000

var wg sync.WaitGroup

func main() {
	fmt.Println("hello atomic")
	var num, num1 int32
	runtime.GOMAXPROCS(4)
	wg.Add(1)

	st := time.Now()

	go check(&num)

	for i := 0; i < TOTAL; i++ {
		//	go AddNum(&num)
		go AddNumD(&num1) //5 ns per add
	}
	et := time.Now().Sub(st)
	fmt.Println("first :", et)

	st = time.Now()
	for i := 0; i < TOTAL; i++ {
		go AddNum(&num)
		//	go AddNumD(&num1)	// 10ns per add
	}
	et = time.Now().Sub(st)
	go ChgNum(&num)
	fmt.Println("second :", et)

	for i := 0; i < TOTAL/DEV; i++ {
		fmt.Println(atomic.LoadInt32(&num))
		fmt.Println(atomic.LoadInt32(&num1))
		atomic.CompareAndSwapInt32(&num, 660000, 3*6600000)

		time.Sleep(time.Nanosecond * 1)
	}

	//atomic.CompareAndSwapInt32(&num1, 1000000, 2*1000000)
	time.Sleep(time.Second * 2)
	wg.Done()
	wg.Wait()
	fmt.Println(num, num1)
}

func AddNum(num *int32) {
	atomic.AddInt32(num, 1)
	// atomic.CompareAndSwapInt32(num, 660000, 3*6600000)

}

func AddNumD(num *int32) {
	*num++
	atomic.CompareAndSwapInt32(num, 660000, 3*6600000)

}

func ChgNum(num *int32) {
	atomic.CompareAndSwapInt32(num, 500000, 3*5000000)

}

func check(num *int32) {
	for {
		if atomic.CompareAndSwapInt32(num, 1000, 0) {
			break
		}
		time.Sleep(time.Nanosecond * 40)
	}
	fmt.Println("Go Out")
}

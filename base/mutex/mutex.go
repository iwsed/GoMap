package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex


func main() {
	fmt.Println("hello world")
	wg.Add(2)

	msgin 

	for i := 0; i < 10; i++ {
		go CntNum(i)
	}

	wg.Wait()

}

// CntNum Count global number
func CntNum(Node int) {
	fmt.Println("Current Node", Node)

	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()

}

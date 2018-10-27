package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Store struct {
	name string
	num  int
	sync.Mutex
}

func main() {
	fmt.Println("hello sync")
	store := Store{}
	store.name = "Tomato"

	fmt.Println(store)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go store.addNum()
	}
	wg.Wait()
	fmt.Println(store)

}

func (s *Store) addNum() {
	s.Lock()
	s.num = s.num + 1
	s.Unlock()
	wg.Done()
}

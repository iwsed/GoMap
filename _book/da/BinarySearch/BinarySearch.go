package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NUM = 100000

func main() {
	fmt.Println("hello binary search")

	arr := NewArr()

	rand.Seed(time.Now().UnixNano())

	t := time.Now()
	for i := 0; i < NUM; i++ {
		num := rand.Intn(2<<16 - 1)
		// _ = FindV(arr, num)
		_ = FindV2(arr, num, 0, 2<<16-1)

	}
	tt := time.Now().Sub(t) / NUM
	fmt.Println(tt)

}

func FindV(arr *[2 << 16]int, v int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return i
		}
	}
	return 0
}

func FindV2(arr *[2 << 16]int, v int, l int, h int) int {

	if l > h {
		return -1
	}

	mid := l + (h-l)/2

	if v == arr[mid] {
		return arr[mid]
	} else if v <= arr[mid] {
		FindV2(arr, v, l, arr[mid])
	} else {
		FindV2(arr, v, arr[mid], h)
	}

	return 0
}

// 如何返回一个数组指针
func NewArr() *[2 << 16]int {
	arr := [2 << 16]int{}
	//	return &arr

	for i := 0; i < 2<<16; i++ {
		arr[i] = i
	}
	return &arr
}

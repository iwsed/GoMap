package main

import (
	"fmt"
	"sort"
)

type By func(p1, p2 *Person) bool

type Person struct {
	name string
	age  int
}

var persons = []Person{
	{"WangWu", 35}, {"ZhangSan", 20}, {"LiSi", 30},
}

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

func main() {
	fmt.Println("hello funcall")

	name := func(p1, p2 *Person) bool {
		return p1.name < p2.name //正序排列
	}
	By(name).Sort(persons)
	fmt.Println("By name", persons)

	sort.Sort(ByAge(persons))

	fmt.Println("By age", persons)

	var g Grams
	g = 10
	fmt.Println("gram :", g)
}

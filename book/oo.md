# GoLang的面向对象

通过golang实现面向对象的过程， 以及golang的继承关系

定义了动物， 定义了鸟类，鱼类， 最后的对象是 燕子， Nimo



```go
package main

import "fmt"

type Animal interface {
	Eat()
	Play()
}
type animal struct {
	name string
}

type Bird interface {
	Animal //匿名嵌入接口
	Fly()
}

type bird struct {
	animal //匿名嵌入继承父亲的name
	wings  int
}

type Fish interface {
	Animal
	Swim()
}

type fish struct {
	animal //集成父亲的name
	name   string
	fin    int
}

func (a *animal) Eat() {
	fmt.Println(a.name, "is eatting now")
}

func (a *animal) Play() {
	fmt.Println(a.name, "is playing now")
}

func (b *bird) Fly() {
	fmt.Println(b.name, "is flying now")
}

func (b *bird) Eat() { // 重写了父亲的吃的方法
	fmt.Println(b.name, "is eatting too much")
}

func (f *fish) Swim() {
	fmt.Println(f.name, "is swimming now")
}

func (f *fish) Eat() {
	fmt.Println(f.name, "is etting now")
}

func main() {
	fmt.Println("hello golang oo")

	bird := bird{animal{"swallow"}, 2}

	bird.Fly() // bird的特有方法, 	"swallow is flying now"
	bird.Eat() //重写了父亲的方法, 	"swallow is eatting too much"

	fish := fish{animal{"Clown Fish"}, "Nimo", 3}

	fish.Play() // 继承父亲的方法， "Clown Fish is playing now"， 父亲的name
	fish.Eat()  // 重写父亲的方法， "Nimo is etting", 用父亲的name
	fish.Swim() // fish 特有的方法 "Nimo is swimming now";
}


```
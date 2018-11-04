package main

//定义一类方法的集合，当这些方法需要在一个大的方法里面顺序执行的时候，
// 可以按照这个数组执行；  这个可以进行界面化编程， 定义一个方法的操作集；
//函数（方法）作为一个参数传人的作用是什么呢？

// 主体是不同的人， 定义一个type 为人；
// 然后会定义一些行为， 刷牙、洗脸、吃饭、穿衣服。。。
// 前台用户可以生产一些动作， 后台通过动作完成相应的工作；；；

// 简化理解思维， 本质上是一个流程管理， 可以定义一组动作， 然后系统独立的完成这些动作过程；
// 可能动作与动作之间有关联，有数据控制， 但是这个是业务上需要关注的目标；
// 如果说这个过程可以完事，那么就可以实现凡人编程的目的；
// 主要是程序实现过程比较动态，不依赖与主进程， 甚至可以参数化配置；

// 再本质一点思考，即这是一个过程  A 与 B 的交互过程
// A可以是前台、用户、客户端等等；  B也可以是相应的部分；
// A告知B 一个数据、一组数据、一个行为及一组行为的过程；；

import "fmt"

// Human 定义了一个人
type Human struct {
	Name  string
	age   int
	birth string
}

// HumanActivity 定义一个人日常的活动
type HumanActivity func(h *Human)

// BrushTooth 定义一个刷牙的动作
func BrushTooth(h *Human) {
	fmt.Println("begin bruth", h.Name)
}

func EatBef(h *Human) {
	fmt.Println("begin eatting")
}

func PutAHuman(ha []HumanActivity) {
	fmt.Println("action: ", ha[0])
}

func main() {
	var human *Human = &Human{"xiao", 1, ""} //可以是系统注册的一个人，或者游戏里面的一个角色
	var ha [3]HumanActivity
	//var op string // 可以考虑定义一串动作，由前端上送

	// op = "BrushTooth" //前端上送的动作, 后台调用才调用这个动作； 如果前端上送了一组动作，这边可以用一个循环来完成这些动作；
	// if op == "BrushTooth" {
	// 	ha = BrushTooth
	// 	ha(human)
	// }
	ha[0] = BrushTooth

	ha[1] = EatBef

	PutAHuman(ha[:])
	//for 循环完成A用户的一系列动作  for i = 0 ; i < all; i++ {op}

	//PutAHuman(&ha, human)

	fmt.Println(human.Name, human.age)

	var p WorkWithFunc // 定义数组函数，这样可以传人函数

	p = ShowLog
	//p("hw")
	PutAFunc(p)

}

// WorkWithFunc test
type WorkWithFunc func(content string) (n int, err error)

// ShowOnStd test
func ShowOnStd(content string) (bytesNm int, err error) {
	return fmt.Println(content)
}

// ShowLog test
func ShowLog(cont string) (n int, err error) {
	fmt.Println("do nothing")
	return
}

// PutAFunc test
func PutAFunc(p WorkWithFunc) {
	p("")
}

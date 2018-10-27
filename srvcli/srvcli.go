package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("hello world")
	wg.Add(2)

	msg := make(chan string, 10)

	go srv(msg)
	go cli(msg)

	wg.Wait()

}

func srv(msg chan string) {
	tcpListen, err := net.Listen("tcp", ":6002")
	if err != nil {
		panic(err)
	}
	for {
		tcpConn, err := tcpListen.Accept()
		if err != nil {
			fmt.Println("accept fail", err)
			return
		}
		defer tcpConn.Close()
		fmt.Println("连接客户端信息：", tcpConn.RemoteAddr().String())

		msg <- "hello from srv" //如果客户端阻塞，那么数据就无法写入， 超过记录将丢包

	}
}

func cli(msg chan string) {

	// 连接服务端直到成功
CONNECT:
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		fmt.Println("Time to sleep 5s, reconnect to server", time.Now())
		time.Sleep(5 * time.Second)
		goto CONNECT
	}

	// 往通道写数据， 发现通道中断， 重新发起连接
	for {
		//str := <-msg

		str := "****0158ALARM     123456,C,U/123456,C,U/123456,C,U/123456,C,U/161002,A,U/161002,S,U/161002,C,U/161002,C,U/90000000207,S,U/3000,A,U/3000,A,U/88518000,C,U/88518000,A,U/"
		fmt.Println(str)

		bs := []byte(str)

		n, err := conn.Write(bs)
		if err != nil { //判断服务器如果无法写入， 开始重连
			fmt.Println(n, err)
			goto CONNECT
		}

		time.Sleep(1 * time.Second)

	}
}

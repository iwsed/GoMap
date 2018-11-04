# 集成服务器，客户端模块功能


- 继承服务器功能， 接受客户端发送的数据；
- 建立外部客户端连接，
- 从接收方的数据，发送到客户端


实现来一个简单的通信程序， 客户端实现连接，发送； 服务端实现接收，保存，反馈的过程；

可以看到程序是串联执行，先对通道 conn进行数据发送， 然后**同步等待**服务器返回；

## 示范代码

### client发起自动重连

```go
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
		str := "测试验证"
		bs := []byte(str)

		n, err := conn.Write(bs)
		if err != nil { //判断服务器如果无法写入， 开始重连
			fmt.Println(n, err)
			goto CONNECT
		}
		time.Sleep(10 * time.Second)
	}
}

```

client.go
```go

func main() {
    //建立连接
	conn, err := net.Dial("tcp", "45.79.84.15:3331")
	if err != nil {
		log.Fatal(err)
	}
	// 别忘了关闭连接
	defer conn.Close()

	SendServer(conn)    
    ReceServer(conn)
}

func ReceServer(conn net.Conn) {
	// 定义一个接收进程， 用于获取服务端的返回，然后记录文件；
	// 通过 ioutil 来读取连接中的内容，返回一个 []byte 类型的对象
	byte, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	// []byte 类型的数据转成字符串型，再将其打印输出
	fmt.Println(string(byte))
}

// SendServer 定义一个发送进程， 用于获取终端的输入， 然后发送到服务端；
func SendServer(conn net.Conn) {

	data := "abc"
	fmt.Printf("Your name is %s.\r\n", data)

	conn.Write([]byte(data))

}
```

server.go

```go

func main() {
	buf := make([]byte, 1024)
	// net 包提供方便的工具用于 network I/O 开发，包括TCP/IP, UDP 协议等。
	// Listen 函数会监听来自 3331 端口的连接，返回一个 net.Listener 对象。
	li, err := net.Listen("tcp", "localhost:3331")
	// 错误处理
	if err != nil {
		log.Panic(err)
	}
	// 释放连接，通过 defer 关键字可以让连接在函数结束前进行释放
	// 这样可以不关心释放资源的语句位置，增加代码可读性
	defer li.Close()

	// 不断循环，不断接收来自客户端的请求
	for {
		// Accept 函数会阻塞程序，直到接收到来自端口的连接
		// 每接收到一个链接，就会返回一个 net.Conn 对象表示这个连接
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("conn.LocalAddr : ", conn.LocalAddr())
		fmt.Println("conn.RemoteAddr : ", conn.RemoteAddr())

		// 读取内容
		len, err := conn.Read(buf)
		fmt.Println("New client connection.", buf, len)

		// 字符串写入到客户端
		fmt.Fprintln(conn, "Hello from TCP server")
		fmt.Println("New client connection.")

		conn.Close()
	}
}
```
# Golang json的实现

验证json格式的数据，转换为string的格式，通过调用json.Marshal结构之后，会转化为[]bytes格式的内容，可以强制转换为string的格式看到转换后的内容；

- json.Marshal(&struct)
- json.Unmarshal([]byte(strData), &stb)


```go

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}

func main() {
	st := &Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	fmt.Println("before JSON encoding :")
	st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		fmt.Println(string(b))
	}
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))

	fmt.Println("--------------------------------")

	strData := <-ch
	stb := &Student{}
	//stb.ShowStu()
	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}
}

```

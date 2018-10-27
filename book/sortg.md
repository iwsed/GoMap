# 排序

使用golang的sort的方法，实现对指定key的排序.

```go
package sort 
```

系统定义了几个function， 包括：

```go
Len() int
Swap(i, j int)
Less(i, j int)bool 
Sort()      
```

其中， Sort系统默认使用quicksort（快排）

```go
type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

var g Grams
g = 10

fmt.Println("gram :", g)    // output is "Gram : 10g"
```

定义Grams一个类型， 通过String的方法重写g的输出内容进行格式化；；对string的String的方法进行重写；；

最后初步看了一下sort的其它方法， 包括：

```go
sort.Ints(s)
sort.Search(n int, f func(int)bool)int
sort.Sort(data interface)
```


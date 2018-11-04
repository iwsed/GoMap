# ZigZag显示字符串

## Golang 收获
1. []string 格式可以通过strings.Join([]string, "")转换成str格式
2. 开始的时候想到的是数组的解决办法， 随后想到可以用map的方式+string的方式进行追加；
3. numRows作为输入，需要转换理解为N Rows， 对于列方面则需要想象成 (numRows - 1)进制
4. 边界条件需要思考，比如传人的值为1的时候
5. LeetCode 上执行的时间是 88ms， 与本地的16.019us不成比例；

## 案例解释如下

通过zigzag的方式显示字符串， 刚看图的时候不太理解， 突然想到小朋友的漫画书上狗狗跑进了公园， a big zig zag;

明白字符的显示是向下，然后45度向右上方输出；


方案如下：
1. 建立一个N维数组；    // 编码过程被替换
2. 建立一个map，key为0,Rownum; 

```go

func convert(s string, numRows int) string {
	var i int // position of s
	var j int // increase position of row,
	var k int // increase postion of col
	type STR []string
	m := make(map[int]STR) // a build with map , the key s j, from 0, 1, 2, 3

	if numRows == 1 {
		return s
	}

	for i < len(s) {
		for j = 0; j < numRows && i < len(s); j++ {
			if k%(numRows-1) == 0 {
				m[j] = append(m[j], string(s[i]))
				i++
			} else if (numRows - 1 - k%(numRows-1)) == j {
				m[j] = append(m[j], string(s[i]))
				i++
			}
		}
		k++ // col+1
	}
	var ss string
	var str string
	for j := 0; j < numRows; j++ {
		str = strings.Join(m[j], "")
		ss = ss + str
	}
	return ss
}

```



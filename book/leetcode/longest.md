# 在一个字符串中查找最长不重复的字符串

## Go代码收获
1. 通过m[key]的方式给j赋值，如果存在则 j = s[i], ok = true, 如果不存在，则 ok 为 false;
2. for 循环通过 ok来识别；
3. 通过delete 来删除map内容；
4. map的key个数可以通过len(m)来获取；

## 内容如下

题目是： 给出一个长的字符串， 在其中找出不重复的最长字符串的长度；

比如："abcabcbb"， 3， "abc";  "bbbbb", 1, "b"

1. 暴力的方法是使用O（n^3)的方法，即对每个位置上往前进行查找，进行排除；
2. 另外一种是使用map的方式，通过hash的方法进行O(1)的查找定位
3. 其他的办法是定义一个窗口，通过偏移窗口来获取最大值

本处代码的思路是：

设置一个窗口（map),  根据字符串，每次输入一个新的字符，需要进行如下判断：

- 这个字符串是第一个字符， 写入map中
- 判断这个字符是否存在，不存在则写入map中，存在需要判断， 
    - 上次出现的位置在哪里？ 需要删除上次位置的前面内容；

```go
j , ok := m[int(s[i])]
```
## 代码如下

```go
func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 { // 解决空格的问题
		return len(s)
	}
	max := 0
	pos := 0
    m := make(map[int]int)
    
	for i := 0; i < len(s); i++ {
		// 记录最大值max
		if len(m) > max {
			max = len(m)
		}
		if j, ok := m[int(s[i])]; ok {  // 判断是否在map中存在，存在则需要在上次出现前的内容删除；
			cur := pos
			for k := j; k >= cur; k-- {
				delete(m, int(s[k]))
				pos++
			}
		}
		m[int(s[i])] = i
	}
	// 最后一条记录不重复判断
	if len(m) > max {
		max = len(m)
	}
	return max
}
```
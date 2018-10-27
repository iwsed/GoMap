# 排序算法实践

在算法排序了解过程， 其中 冒泡、插入及选择排序， 其中冒泡、选择相对与插入来说低效了一点；

这里就主要介绍插入排序；

插入排序的过程是，取出N+1值（N为1），与1，N-1的内容进行比较，前者小于后者，则进行moving，然后插入；

实现两组有序数据的排序， 代码
```go

    i, j := 0, 0
    tmp := make([]int, 0)

	for i+j < len(arra)+len(arrb) {
		if j >= len(arrb) || i < len(arra) && arra[i] < arrb[j] {
			tmp = append(tmp, arra[i])
			i++
		} else {
			tmp = append(tmp, arrb[j])
			j++
		}
	}
	// fmt.Println(tmp)
```

```go
func InsertSort(arr []int) {
    len := cap(arr)
    
	for i := 1; i < len; i++ { // 从位置N开始
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- { //与为准0开始进行比较
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}
```
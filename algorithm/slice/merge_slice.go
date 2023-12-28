package slice

import (
	"fmt"
	"sort"
)

// 合并两个有序切片为一个有序切片

func mergeSortedSlices(a []int,b []int) []int {
	result := make([]int,0,len(a)+len(b))  // 初始化切片类型、长度 和 容量
	i,j := 0,0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func MergeSortSlice() {
	a := []int{2, 5, 8, 9}
	b := []int{1, 6, 7}

	// 归并排序
	// c := mergeSortedSlices(a, b)

	// sort 包
	c := make([]int,0,len(a)+len(b))
	c = append(c,a...)
	c = append(c,b...)
	/*sort.Ints(c)*/

	// sort 闭包
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	// sort 降序
	// sort.Sort(sort.Reverse(sort.IntSlice(c)))

	fmt.Println(c)
}

package main

import "fmt"

type element struct {
	index     int // 横坐标
	curHight  int // 当前高度
	baseHight int // base 高度（base 以下的元素已经算过了）
}

func trap(height []int) int {
	var (
		stack []element
		total int
	)

	for i, h := range height {
		if h == 0 {
			continue
		}
		var baseHight = h
		// 从右往左找自身能够 cover 住的元素
		for j := len(stack) - 1; j >= 0; j-- {
			var (
				element   = stack[j]
				width     = (i - element.index - 1)
				minHeight = element.curHight
			)
			if h < minHeight {
				minHeight = h
			}
			// 计算水位
			total += (minHeight - element.baseHight) * width

			if element.curHight <= h {
				// 比自己低的元素都可以去掉
				stack = stack[:j]
			} else {
				// 比自己高的元素不能去，后面还要算
				break
			}
		}
		for j := len(stack) - 1; j >= 0; j-- {
			// 提高左边所有元素的基础水位线（水位线以下的元素都清算过了）
			if stack[j].baseHight < baseHight {
				stack[j].baseHight = baseHight
			}
		}
		// Push element
		stack = append(stack, element{i, h, 0})
	}
	return total
}

func main() {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

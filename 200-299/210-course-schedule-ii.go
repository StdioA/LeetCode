package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 拓扑排序裸题
	// 常数优化思路：
	// 入度不需要记节点，只需要记数量即可
	// 出度直接用邻接表，不用做 map
	var (
		inDegrees  = make([]int, numCourses)   // 记录节点入度
		outDegrees = make([][]int, numCourses) // 记录出度
		visited    = make([]bool, numCourses)
		result     = make([]int, 0, numCourses)
	)
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		inDegrees[a]++
		outDegrees[b] = append(outDegrees[b], a)
	}
	for {
		// 找到入度为 0 的节点
		var (
			target   = -1
			allEmpty = true
		)
		for node, ins := range inDegrees {
			if visited[node] {
				continue
			}
			if ins == 0 {
				target = node
			} else {
				allEmpty = false
			}
		}
		// 所有节点都被访问过了，则退出
		if allEmpty {
			break
		}
		// 存在环
		if target == -1 {
			return []int{}
		}
		visited[target] = true
		result = append(result, target)
		// 根据出度剔除节点
		for _, node := range outDegrees[target] {
			inDegrees[node]--
		}
	}
	// 把没访问过的节点放到最后
	for node, val := range visited {
		if !val {
			result = append(result, node)
		}
	}
	return result
}

func main() {
	pre := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	fmt.Println(findOrder(20, pre))
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
}

package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 拓扑排序裸题
	var (
		graph   = make([]map[int]struct{}, numCourses) // 记录节点入度
		out     = make([]map[int]struct{}, numCourses) // 记录出度
		visited = make([]bool, numCourses)
	)
	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		if graph[a] == nil {
			graph[a] = make(map[int]struct{})
		}
		graph[a][b] = struct{}{}
		if out[b] == nil {
			out[b] = make(map[int]struct{})
		}
		out[b][a] = struct{}{}
	}
	for {
		// 找到入度为 0 的节点
		var (
			target   = -1
			allEmpty = true
		)
		for node, ins := range graph {
			if visited[node] {
				continue
			}
			if len(ins) == 0 {
				target = node
			} else {
				allEmpty = false
			}
		}
		// 所有节点都被访问过了
		if allEmpty {
			return true
		}
		// 存在环
		if target == -1 {
			return false
		}
		visited[target] = true
		// 根据出度剔除节点
		for node := range out[target] {
			delete(graph[node], target)
		}
	}
}

func main() {
	pre := [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	fmt.Println(canFinish(20, pre))
}

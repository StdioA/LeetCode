package main

type Node struct {
	Val       int
	Neighbors []*Node
}

type status int

const (
	notExist   = iota // 新节点不存在
	processing        // 新节点已生成，但边没有构建完全
	finished          // 节点的所有边都构建完全了
)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var (
		nodeMap   = make(map[*Node]*Node) // 节点映射
		statusMap = make(map[*Node]status)
		queue     []*Node // 待处理的节点队列
		index     int
	)

	// 生成第一个节点
	copyNode := &Node{
		Val: node.Val,
	}
	nodeMap[node] = copyNode
	statusMap[node] = processing
	queue = []*Node{node}

	for index < len(queue) {
		// 本轮处理的节点
		node := queue[index]
		// 新节点应该是已经生成过的，不应该存在 nil
		copyNode := nodeMap[node]
		for _, neighbor := range node.Neighbors {
			switch statusMap[neighbor] {
			case notExist:
				// 创建节点
				copyNeighbor := &Node{
					Val: neighbor.Val,
				}
				nodeMap[neighbor] = copyNeighbor
				statusMap[neighbor] = processing
				queue = append(queue, neighbor)
				fallthrough
			case processing:
				// 构建双向边
				copyNeighbor := nodeMap[neighbor]
				copyNode.Neighbors = append(copyNode.Neighbors, copyNeighbor)
				copyNeighbor.Neighbors = append(copyNeighbor.Neighbors, copyNode)
			case finished:
				continue
			}
		}
		statusMap[node] = finished
		index++
	}
	return copyNode
}

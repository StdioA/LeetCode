package main

func canCompleteCircuit(gas []int, cost []int) int {
	var (
		total  int
		minSum int = 1e8
		start  int
	)
	for i := range gas {
		total += gas[i] - cost[i]
		// 找到一个净费油量的拐点
		// 到这个地方所需要的油最多，那一定是这个拐点触发
		if total < minSum {
			minSum = total
			start = i + 1
		}
	}
	if total < 0 { // 跑不完一圈
		return -1
	}
	return start % len(gas)
}

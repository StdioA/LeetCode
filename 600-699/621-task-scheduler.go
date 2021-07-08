package main

import "sort"

type taskStruct struct {
	task   byte
	amount int
}

// 优先队列，纯模拟，比后面的解慢 185 倍…
func leastIntervalSlow(tasks []byte, n int) int {
	var counter = make(map[byte]int)
	for _, task := range tasks {
		counter[task]++
	}
	var queue []*taskStruct // 粗糙的优先队列
	for task, amount := range counter {
		queue = append(queue, &taskStruct{
			task, amount,
		})
	}
	sort.Slice(queue, func(i, j int) bool {
		return queue[i].amount > queue[j].amount
	})

	var (
		currentTime  int
		lastExecuted = make(map[byte]int)
		allDone      bool
	)

	for !allDone {
		// 从队列中找出剩余最多的任务
		var (
			targetIndex int
			targetTask  *taskStruct
		)
		allDone = true
		currentTime++
		for i, task := range queue {
			if task.amount == 0 {
				continue
			}
			allDone = false
			if lastExecuted[task.task] == 0 || currentTime-lastExecuted[task.task] > n {
				targetIndex, targetTask = i, task
				break
			}
		}
		// 找到任务就执行，找不到就 idle 一轮
		if targetTask != nil {
			// 执行任务
			lastExecuted[targetTask.task] = currentTime
			targetTask.amount--
			for targetIndex < len(queue)-1 && targetTask.amount < queue[targetIndex+1].amount {
				// 将 task 往后挪
				queue[targetIndex], queue[targetIndex+1] = queue[targetIndex+1], queue[targetIndex]
				targetIndex++
			}
		}
	}
	return currentTime - 1 // 最后一轮 queue 遍历空转了，扣掉 1 轮
}

func leastInterval(tasks []byte, n int) int {
	var (
		freq    = make([]int, 26)
		maxFreq int
	)
	for _, task := range tasks {
		index := int(task - 'A')
		freq[index]++
		if freq[index] > maxFreq {
			maxFreq = freq[index]
		}
	}
	// how many task executed with maxFreq
	var count int
	for _, f := range freq {
		if f == maxFreq {
			count++
		}
	}
	result := (n+1)*(maxFreq-1) + count
	if len(tasks) > result {
		result = len(tasks)
	}
	return result
}

// 简单优化，减少一轮遍历
func leastIntervalOptimize(tasks []byte, n int) int {
	var (
		freq    = make([]int, 26) // 用 slice 而不用 map 是因为需要的元素数量有限且确定，用 map 会占用更多内存
		maxFreq int
		count   int
	)
	for _, task := range tasks {
		index := int(task - 'A')
		freq[index]++
		if freq[index] > maxFreq {
			maxFreq = freq[index]
			count = 1
		} else if freq[index] == maxFreq {
			count++
		}
	}
	result := (n+1)*(maxFreq-1) + count
	if len(tasks) > result {
		result = len(tasks)
	}
	return result
}

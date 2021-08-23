package main

import "fmt"

func findCircleNum(M [][]int) int {
	unionSet := make(map[int]int)
	find := func(num int) int {
		for unionSet[num] > 0 {
			num = unionSet[num]
		}
		return num
	}
	for i := 0; i < len(M); i++ {
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 0 {
				continue
			}
			fi, fj := find(i), find(j)
			if fi != fj {
				unionSet[fi] = fj
			}
		}
	}
	count := 0
	for i := 0; i < len(M); i++ {
		if unionSet[i] == 0 {
			count++
		}
	}
	return count
}

func findCircleNum2(isConnected [][]int) int {
	type union struct {
		root *union
	}
	var (
		total     = len(isConnected)
		unions    = make([]*union, total)
		seperated = total
	)
	for i := 0; i < total; i++ {
		unions[i] = &union{}
	}

	for i, row := range isConnected {
		for j, val := range row {
			if val == 0 {
				continue
			}
			// join
			ui, uj := unions[i], unions[j]
			for ui.root != nil {
				ui = ui.root
			}
			for uj.root != nil {
				uj = uj.root
			}
			if ui != uj {
				ui.root = uj
				seperated--
			}
		}
	}
	return seperated
}

func main() {
	relations := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(relations))
}

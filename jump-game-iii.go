package main

func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}
	step := arr[start]
	// Already searched
	if step < 0 {
		return false
	}
	arr[start] = -1

	prev, next := start-step, start+step
	if prev >= 0 {
		prevValid := canReach(arr, prev)
		if prevValid {
			return true
		}
	}
	if next < len(arr) {
		nextValid := canReach(arr, next)
		if nextValid {
			return true
		}
	}
	return false
}

func numTimesAllBlue(light []int) int {
	var max, count int
	for i, l := range light {
		if l > max {
			max = l
		}
		if max == i+1 {
			count++
		}
	}
	return count
}
package main

func hasAllCodes(s string, k int) bool {
	if len(s) < 1<<k-1 { // a magic number _(:з」∠)_
		return false
	}

	var (
		lowMarker  = 0
		highMarker = 1<<k - 1
		mask       = 1<<k - 1
		store      = make(map[int]bool)
		currentNum int
	)
	for i := 0; i < k; i++ {
		currentNum = currentNum<<1 + int(s[i]-'0')
	}
	store[currentNum] = true
	for i := k; i < len(s) && lowMarker <= highMarker; i++ {
		currentNum = (currentNum<<1)&mask + int(s[i]-'0')
		store[currentNum] = true

		// Shrink store
		for store[lowMarker] {
			delete(store, lowMarker)
			lowMarker++
		}
		for store[highMarker] {
			delete(store, highMarker)
			highMarker--
		}
	}
	return lowMarker > highMarker
}

// 这个版本不主动从 store 删除数据，可内存占用居然只多一点点…
func hasAllCodesFaster(s string, k int) bool {
	if len(s) < 1<<k-1 {
		return false
	}

	var (
		mask       = 1<<k - 1
		store      = make(map[int]bool)
		currentNum int
	)
	for i := 0; i < k; i++ {
		currentNum = currentNum<<1 + int(s[i]-'0')
	}
	store[currentNum] = true
	for i := k; i < len(s) && len(store) < 1<<k; i++ {
		currentNum = (currentNum<<1)&mask + int(s[i]-'0')
		store[currentNum] = true
	}
	return len(store) == 1<<k
}

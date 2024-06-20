package main

func checkOnesSegment(s string) bool {
	var zeroMet bool
	for _, n := range s {
		if n == '0' {
			zeroMet = true
		} else if n == '1' && zeroMet {
			return false
		}
	}
	return true
}

package main

func reverseBits(num uint32) uint32 {
	var (
		result uint32
		bits   int
	)
	for num > 0 {
		result = (result << 1) | (num & 1)
		num = num >> 1
		bits++
	}
	return result << (32 - bits) // 注意这里特别标注了，需要加上后缀 0
}

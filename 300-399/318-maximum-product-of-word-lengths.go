package main

import (
	"fmt"
)

func maxProduct(words []string) int {
	var (
		wordsMap = make(map[int]int)
		maxl     int
	)
	for _, word := range words {
		var (
			wordNum int
			length  = len(word)
		)
		for i := 0; i < length; i++ {
			wordNum |= 1 << int(word[i]-'a')
		}
		for num, numl := range wordsMap {
			if wordNum&num == 0 && numl*length > maxl {
				maxl = numl * length
			}
		}
		if length > wordsMap[wordNum] {
			wordsMap[wordNum] = length
		}
	}
	return maxl
}

func main() {
	var words = []string{
		"bdcecbcadca", "caafd", "bcadc", "eaedfcd", "fcdecf", "dee", "bfedd", "ffafd", "eceaffa", "caabe", "fbdb", "acafbccaa", "cdc", "ecfdebaafde", "cddbabf", "adc", "cccce", "cbbe", "beedf", "fafbfdcb", "ceecfabedbd", "aadbedeaf", "cffdcfde", "fbbdfdccce", "ccada", "fb", "fa", "ec", "dddafded", "accdda", "acaad", "ba", "dabe", "cdfcaa", "caadfedd", "dcdcab", "fadbabace", "edfdb", "dbaaffdfa", "efdffceeeb", "aefdf", "fbadcfcc", "dcaeddd", "baeb", "beddeed", "fbfdffa", "eecacbbd", "fcde", "fcdb", "eac", "aceffea", "ebabfffdaab", "eedbd", "fdeed", "aeb", "fbb", "ad", "bcafdabfbdc", "cfcdf", "deadfed", "acdadbdcdb", "fcbdbeeb", "cbeb", "acbcafca", "abbcbcbaef", "aadcafddf", "bd", "edcebadec", "cdcbabbdacc", "adabaea", "dcebf", "ffacdaeaeeb", "afedfcadbb", "aecccdfbaff", "dfcfda", "febb", "bfffcaa", "dffdbcbeacf", "cfa", "eedeadfafd", "fcaa", "addbcad", "eeaaa", "af", "fafc", "bedbbbdfae", "adfecadcabe", "efffdaa", "bafbcbcbe", "fcafabcc", "ec", "dbddd", "edfaeabecee", "fcbedad", "abcddfbc", "afdafb", "afe", "cdad", "abdffbc", "dbdbebdbb",
	}
	fmt.Println(maxProduct(words))
}

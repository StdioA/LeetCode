package main

import "fmt"

// 并查集
type UnionSet struct {
	unions []int
}

func NewUnionSet() *UnionSet {
	return &UnionSet{
		unions: make([]int, 27), // 1-based
	}
}

func (set *UnionSet) Origin(n int) int {
	for set.unions[n] > 0 {
		n = set.unions[n]
	}
	return n
}

func (set *UnionSet) Add(a, b int) {
	originA, originB := set.Origin(a), set.Origin(b)
	if originA != originB {
		set.unions[originB] = originA
	}
}

func (set *UnionSet) Check(a, b int) bool {
	return set.Origin(a) == set.Origin(b)
}

func equationsPossible(equations []string) bool {
	var unionSet = NewUnionSet()
	for _, eq := range equations {
		a, b := int(eq[0]-'a'+1), int(eq[3]-'a'+1)
		if eq[1] == '=' {
			unionSet.Add(a, b)
		}
	}
	for _, eq := range equations {
		a, b := int(eq[0]-'a'+1), int(eq[3]-'a'+1)
		if eq[1] == '!' && unionSet.Check(a, b) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(equationsPossible([]string{"c==c", "b==d", "x!=z"}))
}

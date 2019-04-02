package bitree

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func assertEqual(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func assertNotEqual(tb testing.TB, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func assertNil(tb testing.TB, val interface{}) {
	if !(val == nil || reflect.ValueOf(val).IsNil()) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, val)
		tb.FailNow()
	}
}

func assertNotNil(tb testing.TB, val interface{}) {
	if val == nil || reflect.ValueOf(val).IsNil() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, val)
		tb.FailNow()
	}
}

func TestMakeBiTree(t *testing.T) {
	tree, err := MakeBiTree("{1, nil,2, 3}")
	assertNil(t, err)
	exp := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	assertEqual(t, *tree, exp)

	// Error handling
	_, err = MakeBiTree("34567890")
	assertNotNil(t, err)
	assertEqual(t, err.Error(), "Bad format")
	_, err = MakeBiTree("{1,nil,aaa,nil,nil,3}")
	assertNotNil(t, err)
	assertEqual(t, err.Error(), "strconv.Atoi: parsing \"aaa\": invalid syntax")
}

func TestMakeBiTreeIntf(t *testing.T) {
	r := []interface{}{1, nil, 2, 3}
	/*
			1
		   / \
		 nil  2
			 /
			3
	*/
	tree := MakeBiTreeIntf(r)
	exp := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	assertEqual(t, *tree, exp)
}

func TestBiTreeString(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	str := tree.String()
	assertEqual(t, str, "{1,nil,2,3}")
}

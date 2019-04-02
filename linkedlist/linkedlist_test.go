package linkedlist

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func assertEqual(tb testing.TB, act, exp interface{}) {
	if !reflect.DeepEqual(act, exp) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestMakeList(t *testing.T) {
	list := New([]int{1, 2, 3})
	exp := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	assertEqual(t, list, exp)
}

func TestMakeListWithCycle(t *testing.T) {
	list := NewWithCycle([]int{1, 2, 3}, 1)
	exp := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	exp.Next.Next.Next = exp.Next
	assertEqual(t, list, exp)
}

func TestBiTreeString(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	assertEqual(t, list.String(), "[1,2,3]")
	list.Next.Next.Next = list.Next
	assertEqual(t, list.String(), "[1,2,3] -> 1")
}

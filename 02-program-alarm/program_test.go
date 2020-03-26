package main

import (
	"testing"
)

type testcase struct {
	program   []int
	exppected []int
	retcode   int
}

func TestProgramRun(t *testing.T) {
	cases := []testcase{
		testcase{
			[]int{1, 9, 10, 3,
				2, 3, 11, 0,
				99,
				30, 40, 50},
			[]int{3500, 9, 10, 70,
				2, 3, 11, 0,
				99,
				30, 40, 50},
			0,
		},
		testcase{
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
			0,
		},
		testcase{
			[]int{2, 3, 0, 3, 99},
			[]int{2, 3, 0, 6, 99},
			0,
		},
		testcase{
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
			0,
		},
		testcase{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			0,
		},
	}

	for _, c := range cases {
		if ret := programRun(c.program) != c.retcode; ret {
			t.Errorf("programRun failed. Expected retcode: %v, got: %v",
				c.retcode, ret)
		}
		for i, op := range c.program {
			if op != c.exppected[i] {
				t.Errorf("programRun failed. Expected: %v, got: %v",
					c.exppected, c.program)
			}
		}
	}
}

package twoSum_test

import (
	"reflect"
	"testing"

	. "github.com/wrathofzombies/leetcode/problems/1-two-sum"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type twoSumTest struct {
	test   string
	nums   []int
	target int
	result []int
}

var testCases = []twoSumTest{
	{"Two numbers add to target", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"Two numbers add to target and a number is half of target", []int{3, 2, 4}, 6, []int{1, 2}},
	{"Two numbers add to target and they are the same", []int{3, 3, 2}, 6, []int{0, 1}},
	{"Any two numbers add to target", []int{3, 3, 3}, 6, []int{0, 2}},
}

func TestTwoSum(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		if output := TwoSum(tc.nums, tc.target); !reflect.DeepEqual(tc.result, output) {
			t.Errorf("Test %s: Output %q not equal to expected %q", tc.test, output, tc.result)
		}
	}
}

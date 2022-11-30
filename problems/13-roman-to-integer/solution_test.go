package romanToInteger_test

import (
	"reflect"
	"testing"

	. "github.com/wrathofzombies/leetcode/problems/13-roman-to-integer"
)

type twoSumTest struct {
	roman  string
	result int
}

var testCases = []twoSumTest{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"CMXCIV", 994},
	{"DXCI", 591},
	{"CXVII", 117},
}

func TestTwoSum(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		if output := RomanToInt(tc.roman); !reflect.DeepEqual(tc.result, output) {
			t.Errorf("Test %s: Output %d not equal to expected %d", tc.roman, output, tc.result)
		}
	}
}

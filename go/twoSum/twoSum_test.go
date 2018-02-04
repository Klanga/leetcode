package twoSum

import "testing"

func TestTwoSum(t *testing.T) {

	numsMap := map[int][]int{
		9:  []int{2, 7, 11, 15},
		16: []int{3, 4, 12, 11},
		20: []int{9, 10, 12, 10},
	}

	resultMap := map[int][]int{
		9:  []int{0, 1},
		16: []int{1, 2},
		20: []int{1, 3},
	}

	for target := range numsMap {
		if array := twoSum(numsMap[target], target); array != nil && array[0] != resultMap[target][0] && array[1] != resultMap[target][1] {
			t.Errorf("target : '%d' Test Failed!\n", target)
		} else {
			t.Logf("target : '%d' Test Passed!\n", target)
		}
	}
}

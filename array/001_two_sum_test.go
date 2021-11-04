package array

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	if !arrayEqual(TwoSum(nums, target), []int{0, 1}) {
		t.Errorf("nums: %v, target: %d, found wrong result, result: %v", nums, target, TwoSum(nums, target))
	}
	nums = []int{0, 0, 0, 0}
	target = 2
	if !arrayEqual(TwoSum(nums, target), nil) {
		t.Errorf("nums: %v, target: %d, found wrong result, result: %v", nums, target, TwoSum(nums, target))
	}
}

func arrayEqual(arr1 []int, arr2 []int) bool {
	if arr1 == nil && arr2 == nil {
		return true
	}

	if len(arr1) != len(arr2) {
		return false
	}

	for k, v := range arr1 {
		if v != arr2[k] {
			return false
		}
	}
	return true
}

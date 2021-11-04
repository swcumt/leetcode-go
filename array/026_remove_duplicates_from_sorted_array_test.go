package array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3}
	expectNums := []int{1, 2, 3}
	k := RemoveDuplicates(nums)
	if len(expectNums) != k {
		t.Errorf("nums: %v not match %v", nums, expectNums)
		return
	}
	for i := 0; i < k; i++ {
		if nums[i] != expectNums[i] {
			t.Errorf("nums: %v not match %v", nums, expectNums)
			return
		}
	}
}

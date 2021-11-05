package array

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	t.Log(RemoveElement(nums, 4))
	t.Log(nums)
}

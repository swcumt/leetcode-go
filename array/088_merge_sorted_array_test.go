package array

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1}
	var nums2 []int
	m, n := 1, 0
	Merge(nums1, m, nums2, n)
	t.Log(nums1)
}

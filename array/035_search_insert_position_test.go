package array

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 0
	t.Log(SearchInsert(nums, target))
}

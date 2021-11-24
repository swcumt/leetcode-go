package string

/**
Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].

For example, swapping at indices 0 and 2 in "abcd" results in "cbad".


Example 1:

Input: s = "ab", goal = "ba"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
Example 2:

Input: s = "ab", goal = "ab"
Output: false
Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.
Example 3:

Input: s = "aa", goal = "aa"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.
Example 4:

Input: s = "aaaaaaabc", goal = "aaaaaaacb"
Output: true


Constraints:

1 <= s.length, goal.length <= 2 * 104
s and goal consist of lowercase letters.


https://leetcode-cn.com/problems/buddy-strings
*/

func BuddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var sFirst uint8 = 1
	var gFirst uint8 = 1
	var sSecond uint8 = 2
	var gSecond uint8 = 2
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			continue
		}
		num++
		if num > 2 {
			return false
		}
		if num == 1 {
			sFirst = s[i]
			gFirst = goal[i]
		}
		if num == 2 {
			sSecond = s[i]
			gSecond = goal[i]
			if gSecond != sFirst || sSecond != gFirst {
				return false
			}
		}
	}
	if num == 0 {
		r := 1 << (s[0] - 'a')
		for i := 1; i < len(s); i++ {
			if r&(1<<(s[i]-'a')) != 0 {
				return true
			}
			r |= 1 << (s[i] - 'a')
		}
	}
	if num != 2 {
		return false
	}

	return true
}

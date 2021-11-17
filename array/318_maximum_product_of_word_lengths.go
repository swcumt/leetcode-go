package array

/**
Given a string array words, return the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. If no such two words exist, return 0.

Example 1:

Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16
Explanation: The two words can be "abcw", "xtfn".
Example 2:

Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4
Explanation: The two words can be "ab", "cd".
Example 3:

Input: words = ["a","aa","aaa","aaaa"]
Output: 0
Explanation: No such pair of words.

Constraints:

2 <= words.length <= 1000
1 <= words[i].length <= 1000
words[i] consists only of lowercase English letters.
*/

func MaxProduct(words []string) int {
	nums := make([]int, len(words))

	for k, v := range words {
		for _, c := range v {
			nums[k] |= 1 << (c - 'a')
		}
	}
	result := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]&nums[j] != 0 {
				continue
			}
			tmp := len(words[i]) * len(words[j])
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

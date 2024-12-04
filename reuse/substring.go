// Given a string s, find the length of the longest substring without repeating characters.

package reuse

func FindLongestSubstring(s string) int {

	var visited [256]bool // initialized with 'false' values

	b := []byte(s)

	maxLength, left, right := 0, 0, 0

	for right < len(b) {
		for visited[int(b[right])] {
			visited[int(b[left])] = false
			left++
		}
		visited[int(b[right])] = true
		right++
		if maxLength < (right - left) {
			maxLength = right - left
		}
	}

	return maxLength

}

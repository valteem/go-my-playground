// https://stackoverflow.com/a/67259315

package reuse

import (
	"regexp"
)

func CountMatches(s string, re *regexp.Regexp) int {
	count := 0
	for i := 0; i < len(s); {
		remaining := s[i:]                   // slicing the string
		loc := re.FindStringIndex(remaining) // location of next match
		if loc == nil {
			break
		}
		// loc[0] is the start index of the match,
		// loc[1] is the end index (exclusive)
		i += loc[1]
		count++
	}
	return count
}

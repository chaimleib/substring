package substring

// Given an input string, returns the length of the longest substring where no
// characters are repeated. If multiple substrings are tied for the longest,
// returns the last winner.
func LongestNonrepeating(s string) string {
	last := make(map[rune]int)
	winner := ""
	start := 0
	for i, r := range s {
		startOrZero, ok := last[r] // defaults to 0
		if ok && startOrZero+1 > start {
			start = startOrZero + 1
		}
		current := i - start + 1
		last[r] = i
		if current > len(winner) {
			winner = s[start : start+current]
		}
	}
	return winner
}

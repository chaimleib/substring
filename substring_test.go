package substring

import "testing"

func TestLongestNonrepeating(t *testing.T) {
	cases := []struct {
		In  string
		Out int
	}{
		{"", 0},
		{"a", 1},
		{"abcbcbccbcbabc", 3},
		{"abcdbefg", 6},
	}
	for _, c := range cases {
		actual := LongestNonrepeating(c.In)
		if c.Out != len(actual) {
			t.Errorf("Fail: %s -> %d, instead of %d\n", c.In, len(actual), c.Out)
		}
	}
}
